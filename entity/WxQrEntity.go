package entity

//微信H5支付实体
type WxQrEntity struct {
	ConmmonEntity
	PayType           int              `json:"pay_type"`
	AgentId           int              `json:"agent_id"`
	AgentBillId       string           `json:"agent_bill_id"`
	PayAmt            float64          `json:"pay_amt"`
	NotifyUrl         string           `json:"notify_url"`
	ReturnUrl         string           `json:"return_url"`
	UserIp            string           `json:"user_ip"`
	AgentBillTime     int64            `json:"agent_bill_time"`
	GoodsName         string           `json:"goods_name"`
	GoodsNum          int              `json:"goods_num"`
	Remark            string           `json:"remark"`
	GoodsNote         string           `json:"goods_note"`
}

func (entity *WxQrEntity)GetSign() []string  {
	return []string {
		"version",
		"agent_id",
		"agent_bill_id",
		"agent_bill_time",
		"pay_type",
		"pay_amt",
		"notify_url",
		"return_url",
		"user_ip",
		"key",
	}
}
