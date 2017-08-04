package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
)

// @Title 生成密码
// @Description create AccountAccount
// @Param	body		body 	models.AccountAccount	true		"body for AccountAccount content"
// @Success 201 {int} models.AccountAccount
// @Failure 403 body is empty
func PasswordMD5(passwd, salt string) string {
	h := md5.New()
	// 后面增加一个无意义字符串
	h.Write([]byte(passwd + salt + "@.YnO-"))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}

// StringsJoin 字符串拼接
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}

// GetInt64 类型转换，获得int64
func GetInt64(v interface{}) (re int64, err error) {
	switch v.(type) {
	case string:
		re, err = strconv.ParseInt(v.(string), 10, 64)
	case float64:
		re = int64(v.(float64))
	case float32:
		re = int64(v.(float32))
	default:
		err = errors.New("不能转换")
	}
	return
}
