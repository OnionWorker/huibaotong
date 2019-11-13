package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"
	"github.com/OnionWorker/huibaotong"
	"github.com/OnionWorker/huibaotong/conf"
	"github.com/OnionWorker/huibaotong/entity"
)

/**
	测试SDK入口类传参使用方法
 */
func TestHuiBaoSdk(t *testing.T){
	HuiBaoSDK,Err := huibaotong.NewHuiBao(conf.WXH5PAY,"94C49D3629584B918A06BE9F")
	if Err != nil{
		fmt.Print(Err)
	}

	PayDataEntity := &entity.PayEntity{
		AgentId:2122667,
		UserIp:strings.Replace("192.168.0.1", ".", "_", -1),
		PayType:30,
		//IsPhone:1,
		Scene:"h5",
		PaymentMode:"cashier",
		MetaOption:"eyJzIjoiV0FQIiwibiI6Ir6ptqu52c34IiwiaWQiOiJodHRwczovL20uamQuY29tIn0%3d",
		BankCardType:"-1",
		AgentBillId:"25652222565225",
		AgentBillTime:20170705170201,
		NotifyUrl:"http://xxxxx/test/notify.aspx",
		ReturnUrl:"http://xxxxx/test/return.aspx",
		GoodsNum:1,
		GoodsNote:"111",
		Remark:"",
		GoodsName:"111",
		PayAmt:100.50,
	}
	PayDataEntity.Version = 1
	Result := HuiBaoSDK.SetEntity(PayDataEntity).Excute(conf.XMLRETURN)
	t.Log(Result)
}

func TestHuiBaoSdkTwo(t *testing.T){
	HuiBaoSDK,Err := huibaotong.NewHuiBao(conf.WXH5PAY,"94C49D3629584B918A06BE9F")
	if Err != nil{
		fmt.Print(Err)
	}

	PayDataEntity := &entity.PayEntity{
		AgentId:2122667,
		UserIp:strings.Replace("192.168.0.1", ".", "_", -1),
		PayType:22,
		Scene:"h5",
		PaymentMode:"cashier",
		MetaOption:"eyJzIjoiV0FQIiwibiI6Ir6ptqu52c34IiwiaWQiOiJodHRwczovL20uamQuY29tIn0%3d",
		BankCardType:"-1",
		AgentBillId:"256522225fgfg25",
		AgentBillTime:20170705170201,
		NotifyUrl:"http://xxxxx/test/notify.aspx",
		ReturnUrl:"http://xxxxx/test/return.aspx",
		GoodsNum:1,
		GoodsNote:"111",
		Remark:"",
		GoodsName:"111",
		PayAmt:100.50,
	}
	PayDataEntity.Version = 1
	//Err = HuiBaoSDK.SetEntity(PayDataEntity).Excute(huibaotong.Url("http://www.baidu.com"))
	Result := HuiBaoSDK.SetEntity(PayDataEntity).Excute(conf.XMLRETURN)
	t.Log(Result)
}
//测试PC扫码支付
func TestPcQrPay(t *testing.T){
	HuiBaoSDK,Err := huibaotong.NewHuiBao(conf.PCQRPAY,"BF095F5F8C984115BF9AFD7F")
	if Err != nil{
		fmt.Print(Err)
		return
	}
	//公共参数
	PcQrEntity := &entity.NewCommonEntity{
		Version:"1.0",
		Method:"heemoney.pay.applypay",
		AppId:"hyp191112119487000020549C0C330C4",
		MchUid:"1194872122863",
		Charset:"UTF-8",
		Timestamp:"20191113155001",
		SignType:"MD5",
	}
	//业务参数
	PayData := &entity.PcQrPayEntity{
		OutTradeNo:"256522225fgfg25SS",
		Subject:"测试商品",
		TotalFee:100,
		ChannelType:"ALI_QRCODE",
		ClientIp:"127.0.0.1",
		NotifyUrl:"http://xxxxx/test/notify.aspx",
	}
	jsons,err := json.Marshal(PayData)
	if err != nil{
		fmt.Println(err)
		return
	}
	PcQrEntity.BizContent = string(jsons)
	Result := HuiBaoSDK.SetEntity(PcQrEntity).Excute(conf.JSONRETURN)
	t.Log(Result)
}

//签名加密后数据数据：738a935ecaaf587db74c4d2f4bdb3c25
//                    c2b1e7124156416e6f2d762b630be38e


//是  %b7%a2%cb%cd%b5%bd
func TestHanzi(t *testing.T){
	str := "是"
	var data []byte = []byte(str)
	reader := transform.NewReader(bytes.NewReader(data), simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {

	}

	fmt.Println(string(d[:]))
	fmt.Println(url.QueryEscape(string(d[:])))
}

func TestHanziss(t *testing.T){
	str := "是"
	fmt.Println(url.QueryEscape(str))
}