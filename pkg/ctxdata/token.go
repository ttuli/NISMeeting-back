/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package ctxdata

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/livekit/protocol/auth"
)

const Identify = "IMchatIM"

func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func GetLiveKitToken(apiKey, apiSecret, uid, room string, metaData map[string]string) (string, error) {
	d, err := json.Marshal(metaData)
	if err != nil {
		return "", err
	}
	at := auth.NewAccessToken(apiKey, apiSecret).
		SetVideoGrant(&auth.VideoGrant{
			RoomJoin: true,
			Room:     room,
		}).SetIdentity(uid).SetValidFor(15 * time.Minute).SetMetadata(string(d))

	token, err := at.ToJWT()
	if err != nil {
		return "", err
	}
	return token, err
}
