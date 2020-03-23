package validate

import (
	"time"
)

type Validator interface {

	// 根据所给的key生成验证码
	GenerateCode(key string, expire time.Duration) error
}
