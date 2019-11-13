package entity

type OldAuthEntity struct {
	AgentId         string       `json:"agent_id"` //商家ID，由汇付宝提供
	BillId          string       `json:"bill_id"` //商户订单号
	BillTime        string       `json:"bill_time"` //商户订单时间，(格式为yyyyMMddHHmmss 4位年+2位月+2位日+2位时+2位分+2位秒)
	BankCardType    string       `json:"bank_card_type"` //卡类型（1=储蓄卡，2=信用卡） 默认为1
	BankCardInfo    string       `json:"bank_card_info"` //银行卡信息，格式为：银行卡号|身份证号|身份证姓名|手机号，双方协商的对称加密，使用3DES加密，3DES密钥由汇付宝提供，请求数据前请UrlEncode编码
	ClientIp        string       `json:"client_ip"` //用户来源IP，用户提交信息所在的IP
	TimeStamp       string       `json:"time_stamp"` //提交时间戳(格式为yyyyMMddHHmmss 4位年+2位月+2位日+2位时+2位分+2位秒)，1970/1/1 0点到现在的毫秒值
	VersionId       string       `json:"version_id"` //接口版本号，默认1
	IsTest          string       `json:"is_test"` //是否为测试环境，1=测试环境version_id=2时有效，非测试请不要传本参数，测试环境status默认返回1
	Sign            string       `json:"sign"` //签名结果
}

func (entity *OldAuthEntity)GetSign() []string  {
	return []string {
		"agent_id",
		"bill_id",
		"bill_time",
		"bank_card_type",
		"bank_card_info",
		"time_stamp",
		"key",
	}
}
