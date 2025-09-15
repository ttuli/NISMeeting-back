package file

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"time"

	"Meeting/apps/file/api/internal/svc"
	"Meeting/apps/file/api/internal/types"

	"github.com/aliyun/credentials-go/credentials"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignatureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getSignature
func NewGetSignatureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignatureLogic {
	return &GetSignatureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSignatureLogic) GetSignature(req *types.GetSignatureReq) (resp *types.GetSignatureResp, err error) {
	data, err := l.getPolicyToken()
	if err != nil {
		return nil, err
	}
	return &types.GetSignatureResp{
		Data: data,
	}, nil
}

func (l *GetSignatureLogic) getPolicyToken() (types.PolicyToken, error) {
	bucketName := l.svcCtx.Config.OssConf.BucketName
	region := l.svcCtx.Config.OssConf.Region
	dir := l.svcCtx.Config.OssConf.Dir
	product := l.svcCtx.Config.OssConf.Product

	host := fmt.Sprintf("https://%s.oss-%s.aliyuncs.com", bucketName, region)

	config := new(credentials.Config).
		SetType("ram_role_arn").
		SetAccessKeyId(os.Getenv("ossAccessKeyId")).
		SetAccessKeySecret(os.Getenv("ossAccessKeySecret")).
		SetRoleArn(os.Getenv("ossRoleArn")).
		SetRoleSessionName("mxshop").
		SetPolicy("").
		SetRoleSessionExpiration(3600)

	provider, err := credentials.NewCredential(config)
	if err != nil {
		return types.PolicyToken{}, err
	}

	cred, err := provider.GetCredential()
	if err != nil {
		return types.PolicyToken{}, err
	}

	utcTime := time.Now().UTC()
	date := utcTime.Format("20060102")
	expiration := utcTime.Add(1 * time.Hour)
	policyMap := map[string]any{
		"expiration": expiration.Format("2006-01-02T15:04:05.000Z"),
		"conditions": []any{
			map[string]string{"bucket": bucketName},
			map[string]string{"x-oss-signature-version": "OSS4-HMAC-SHA256"},
			map[string]string{"x-oss-credential": fmt.Sprintf("%v/%v/%v/%v/aliyun_v4_request", *cred.AccessKeyId, date, region, product)},
			map[string]string{"x-oss-date": utcTime.Format("20060102T150405Z")},
			map[string]string{"x-oss-security-token": *cred.SecurityToken},
		},
	}

	policy, err := json.Marshal(policyMap)
	if err != nil {
		log.Fatalf("json.Marshal fail, err:%v", err)
	}

	stringToSign := base64.StdEncoding.EncodeToString([]byte(policy))

	hmacHash := func() hash.Hash { return sha256.New() }
	signingKey := "aliyun_v4" + *cred.AccessKeySecret
	h1 := hmac.New(hmacHash, []byte(signingKey))
	io.WriteString(h1, date)
	h1Key := h1.Sum(nil)

	h2 := hmac.New(hmacHash, h1Key)
	io.WriteString(h2, region)
	h2Key := h2.Sum(nil)

	h3 := hmac.New(hmacHash, h2Key)
	io.WriteString(h3, product)
	h3Key := h3.Sum(nil)

	h4 := hmac.New(hmacHash, h3Key)
	io.WriteString(h4, "aliyun_v4_request")
	h4Key := h4.Sum(nil)

	h := hmac.New(hmacHash, h4Key)
	io.WriteString(h, stringToSign)
	signature := hex.EncodeToString(h.Sum(nil))

	return types.PolicyToken{
		Policy:           stringToSign,
		SecurityToken:    *cred.SecurityToken,
		SignatureVersion: "OSS4-HMAC-SHA256",
		Credential:       fmt.Sprintf("%v/%v/%v/%v/aliyun_v4_request", *cred.AccessKeyId, date, region, product),
		Date:             utcTime.UTC().Format("20060102T150405Z"),
		Signature:        signature,
		Host:             host, // 返回 OSS 上传地址
		Dir:              dir,  // 返回上传目录
	}, nil
}
