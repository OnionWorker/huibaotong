package conf

import "errors"

type HuiBaoType int

const (
	CUSTOM HuiBaoType = iota
	WXQRPAY
	WXH5PAY
	ALIQRPAY
	ALIH5PAY
	PCQRPAY
)

var HuiBaoApiMap map[HuiBaoType]string = map[HuiBaoType]string{
	CUSTOM:  "",  //自定义URL
	WXQRPAY:  "https://pay.heepay.com/Payment/Index.aspx",  //微信扫码支付接口
	WXH5PAY:  "https://Pay.Heepay.com/DirectPay/applypay.aspx",  //微信H5支付接口
	//WXH5PAY:  "http://211.103.157.45/PayHeepay/DirectPay/applypay.aspx",  //微信H5支付接口
	ALIQRPAY:  "https://pay.heepay.com/Payment/Index.aspx",  //支付宝扫码支付接口
	ALIH5PAY:  "https://pay.heepay.com/Payment/Index.aspx",  //支付宝H5支付接口
	PCQRPAY:  "https://pay.heemoney.com/v1/ApplyPay",  //pc扫码支付接口
}

const (
	XMLRETURN HuiBaoType = iota
	URLRETURN
	JSONRETURN
)
type Config struct {
	AgentId    int   //商户ID
	SecretKey string //支付秘钥
	UserIp    string //用户IP
	SdkType  HuiBaoType
}

func (this *Config) CheckConfig()error{
	if this.SecretKey == ""{
		return errors.New("SDK: HuiBaoTong SecretKey is required")
	}
	return nil
}

func (this *Config) GetAccess()string{
	return HuiBaoApiMap[this.SdkType]
}