package entity

type AuthQueryEntity struct {
	OutTradeNo           string       `json:"out_trade_no"` //单号,要求64个字符内、且在同一个商户号下唯一
	AuthType             string       `json:"auth_type"` //鉴权类型,0=未知,1=身份证鉴权,2=银行卡鉴权
	AuthDetailType       string       `json:"auth_detail_type"` //鉴权详细类型,0=未知,1=二要素,2=三要素,3=四要素
}
