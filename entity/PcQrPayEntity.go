package entity

type PcQrPayEntity struct {
	OutTradeNo        string       `json:"out_trade_no"` //商户系统内部订单号,要求64个字符内、且在同一个商户号下唯一
	Subject           string       `json:"subject"` //订单标题
	TotalFee          int          `json:"total_fee"` //订单总金额,单位为分
	ChannelType       string       `json:"channel_type"` //	通道类型，扫码支付，WX_NATIVE、ALI_QRCODE、微信小程序：WX_APPLET，微信代扣：WX_WITHHOLD，等具体见接口规则-参数规定-通道类型
	ClientIp          string       `json:"client_ip"` //	APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	NotifyUrl         string       `json:"notify_url"` //	异步通知的地址
}
