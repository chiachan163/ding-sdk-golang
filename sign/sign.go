package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strconv"
)

// 钉钉签名计算方法：
// 1.把timestamp+"\n"+suiteTicket当做签名字符串，suiteSecret/customSecret做为签名密钥，
// 2.使用HmacSHA256算法计算签名，然后进行Base64 encode获取最后结果。
// 3.然后把签名参数再进行urlEncode，加到请求url后面。
// 签名参数：
// timestamp 当前时间戳，单位是毫秒
// suiteTicket 钉钉给应用推送的ticket，测试应用随意填写如：TestSuiteTicket，正式应用需要从推送回调获取suiteTicket
// suiteSecret/customSecret 三方应用或者定制应用的密钥
func DingSign(timestamp int64, ticket string, secret string) string {
	var stringToSign string = strconv.FormatInt(timestamp, 10) + "\n" + ticket
	return ComputeHmacSha256(stringToSign, secret)
}
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}
