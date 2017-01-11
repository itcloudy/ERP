package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// @Title 生成密码
// @Description create AccountAccount
// @Param	body		body 	models.AccountAccount	true		"body for AccountAccount content"
// @Success 201 {int} models.AccountAccount
// @Failure 403 body is empty
func PasswordMD5(passwd, salt string) string {
	h := md5.New()
	h.Write([]byte(passwd + salt))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)

	return result
}
