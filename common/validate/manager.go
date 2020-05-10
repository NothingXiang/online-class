/*
 package:validate
	提供生成验证码，向相应客户端发送验证码和校验等服务
	预计提供邮箱验证，手机号验证，微信openid验证等3种验证器
	目前，三种验证器生成的验证码都存储在redis
*/
package validate

import (
	"fmt"
	"time"

	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/common/utils"
)

const (
	VerifyKey = "Verify:%v"
)

var globalManager *Manager

// 验证器的管理者
type Manager struct {
	validators map[string]Validator
}

func newManager() *Manager {
	return &Manager{
		validators: map[string]Validator{
			"email": &Email,
		},
	}
}

func init() {
	globalManager = newManager()
}

/*--------------------------------------------*/
// 外部包请通过下面几种方式来调用验证器

// 注册一个验证器
func Register(name string, v Validator) {
	globalManager.validators[name] = v
}

// 检查某种验证器是否存在
func CheckValidatorExist(key string) bool {
	_, ok := globalManager.validators[key]

	return ok
}

// 校验验证码
func Validate(key string, value string) bool {

	code, err := dbutil.Redis().Get(fmt.Sprintf(VerifyKey, key)).Result()
	if err == nil && code == value {
		dbutil.Redis().Del(fmt.Sprintf(VerifyKey, key))
		return true
	}

	return false

}

/*--------------------------------------------*/

// 生成并存储验证码,返回生成的验证码
func GenerateCode(key string, expire time.Duration) string {

	code := utils.RandomCode(6)

	dbutil.Redis().SetNX(fmt.Sprintf(VerifyKey, key), code, expire)

	return code
}
