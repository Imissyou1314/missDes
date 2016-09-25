package utils

/**
 * 工具模块
 */
import (
	"crypto/md5"
	"encoding/hex"
)

//  进行MD5加密
//  @params text 待加密的字段个
// @return 加密后的字段
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
