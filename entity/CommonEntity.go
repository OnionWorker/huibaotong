package entity

//汇宝通最新API的公共实体参数
type NewCommonEntity struct {
	Version       string       `json:"version"` //当前接口版本号1.0
	Method        string       `json:"method"` //具体业务接口名称
	AppId         string       `json:"app_id"` //应用ID，商户的应用id
	MchUid        string       `json:"mch_uid"` //商户统一编号
	Charset       string       `json:"charset"` //编码格式默认为UTF-8
	SignType      string       `json:"sign_type"` //商户生成签名字符串所使用的签名算法类型
	Timestamp     string       `json:"timestamp"` //发送请求的时间
	BizContent    interface{}       `json:"biz_content"` //鉴权参数集合,Json格式,长度不限,具体参数见如下业务参数
	Sign          string       `json:"sign"` //鉴权参数集合,Json格式,长度不限,具体参数见如下业务参数
}

func (entity *NewCommonEntity)GetSign() []string  {
	return []string {
		"app_id",
		"biz_content",
		"charset",
		"mch_uid",
		"method",
		"sign_type",
		"timestamp",
		"version",
		"key",
	}
}