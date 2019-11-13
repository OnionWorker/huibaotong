package conf

import "errors"

type HuiBaoType int

const (
	CUSTOM HuiBaoType = iota
	H5PAY
	PCQRPAY
)

var HuiBaoApiMap map[HuiBaoType]string = map[HuiBaoType]string{
	CUSTOM:  "",  //自定义URL
	H5PAY:  "https://Pay.Heepay.com/DirectPay/applypay.aspx",  //微信H5支付接口
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