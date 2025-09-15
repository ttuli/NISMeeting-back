package captcha

import (
	"errors"
	"time"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, time.Minute*5)

func GetCaptcha() (id, base64 string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, base64, _, err = cp.Generate()
	return
}

func AuthCaptcha(id, ans string) error {
	if store.Get(id, false) == "" {
		return errors.New("验证码过期")
	}
	if store.Verify(id, ans, false) {
		return nil
	}
	return errors.New("验证码错误")
}
