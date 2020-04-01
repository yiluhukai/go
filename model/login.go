package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/xlstudio/wxbizdatacrypt"
	"yiluhuakai/logger"
)

type Login struct {
	Code      string `json:"code"`
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type ResTenXun struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type LoginHeader struct {
	Code           string `json:"X-WX-Code" header:"X-WX-Code"`
	ENCRYPTED_DATA string `json:"X-WX-Encrypted-Data" header:"X-WX-Encrypted-Data"`
	IV             string `json:"X-WX-IV" header:"X-WX-IV"`
	Skey           string
	SessionKey     string
	AppId          string
}

// 解密ENCRYPTED_DATA中的敏感数据
func (lh *LoginHeader) AesDecrypt() (decryptedData string, err error) {

	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(lh.SessionKey))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
	lh.Skey = fmt.Sprintf("%x", bs)
	//如果需要对另一个字符串加密，要么重新生成一个新的散列，要么一定要调用h.Res

	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: lh.AppId, SessionKey: lh.SessionKey}
	decryptedDataTemp, err := pc.Decrypt(lh.ENCRYPTED_DATA, lh.IV, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false
	if err != nil {
		logger.LogError("decode failed %v", err)
		return
	}
	logger.LogDebug("decode successfully")

	decryptedData, ok := decryptedDataTemp.(string)
	if !ok {
		logger.LogError("parse interface to string fail")
		err = errors.New("failed to parse interface to string")
		return
	}

	logger.LogDebug("parse interface to string successfully")
	return
}
