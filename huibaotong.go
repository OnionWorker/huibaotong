package huibaotong

import (
	"errors"
	"github.com/OnionWorker/huibaotong/conf"
	"github.com/OnionWorker/huibaotong/entity"
	"github.com/OnionWorker/huibaotong/lib"
)

/*****************************************************************************************************
 *  @Title: 汇包通SDK                                                                                *
 *  @Author: Onion                                                                                   *
 *  @Email: 133433354@qq.com                                                                         *
 *****************************************************************************************************
 *  说明:                                                                                            *
 *         本SDK非官方SDK，纯粹公司用个人封装。其中包含：支付宝扫码支付、支付宝H5支付、微信扫码支付、*
 *         微信H5支付、批量付款。SDK实习配置后自动签名，可继续扩展汇付宝其他的接口。具体扩展规则方式 *
 *         请阅读 Readme 文档  （代码好坏不接受质疑，你可以不用，用希望保留署名） Thank You !        *
 *****************************************************************************************************
*/

type HuiBaoOption interface {
	apply(*HuiBao)
}

type huibaoOption struct {
	f func(*HuiBao)
}

func Url(u string)HuiBaoOption{
	return  newHuiBaoOption(func(o *HuiBao){
		o.Url = u
	})
}

func(fdo *huibaoOption) apply(do *HuiBao){
	fdo.f(do)
}

func newHuiBaoOption(f func(*HuiBao))*huibaoOption{
	return &huibaoOption{
		f:f,
	}
}

type HuiBao struct {
	Config *conf.Config
	errMsg error
	Url string
	entityData entity.Entity
	signData []string
}
/**
	returnType: XMLRETURN = xml URLRETURN = url 例子: XMLRETURN
	opts: 例子 huibaotong.Url("http://www.baidu.com")
 */
func (this *HuiBao) Excute(returnType conf.HuiBaoType,opts ...HuiBaoOption)map[string]interface{}{
	if this.errMsg != nil{
		return map[string]interface{}{"Code":"1001","Msg":this.errMsg.Error(),"Url":""}
	}
	for _,opt := range opts {
		opt.apply(this)
	}
	if this.Config.SdkType == conf.CUSTOM{
		return map[string]interface{}{"Code":"1002","Msg":"sdk type eq CUSTOM,Excute func must have second arg","Url":""}
	}
	Body,Err,Rmap := lib.NewRequest().Run(this.Config,this.signData,this.entityData,this.Url,returnType)
	if Err != nil{
		return map[string]interface{}{"Code":"1003","Msg":Err.Error(),"Url":""}
	}
	switch returnType {
		case conf.XMLRETURN:
			 ReturnData,Err := lib.ParseXml(Body)
			 if Err != nil{
				 return map[string]interface{}{"Code":"1004","Msg":Err.Error(),"Url":""}
			 }
			 var Msg string = "success"
			 if len(ReturnData.RetMsg) > 0{
				 Msg = ReturnData.RetMsg
			 }
			return map[string]interface{}{"Code":ReturnData.RetCode,"Msg":Msg,"Url":ReturnData.RedirectUrl}
		case conf.URLRETURN:
			return map[string]interface{}{"Code":"0","Msg":"success","Url":Body}
		case conf.JSONRETURN:
			return Rmap

	}
	return map[string]interface{}{"Code":"1005","Msg":"result type out now type (XMLRETURN,URLRETURN,JSONRETURN)"}
}

func (this *HuiBao) SetEntity(Entitys entity.Entity)*HuiBao{
	SignData := Entitys.GetSign()
	if len(SignData) == 0{
		this.errMsg = errors.New("SDK: entity.GetSign() result len can not 0")
	}
	this.entityData = Entitys
	this.signData = SignData
	return this
}

//实例化SDK入口
func NewHuiBao(SdkType conf.HuiBaoType ,SecretKey string)(*HuiBao,error){
	Conf := &conf.Config{
		SecretKey:SecretKey,
		SdkType:SdkType,
	}
	Err := Conf.CheckConfig()
	if Err != nil{
		return nil,Err
	}
	HB := &HuiBao{Config:Conf}
	HB.Url = Conf.GetAccess()
	return HB,nil
}





