package test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/OnionWorker/huibaotong"
	"github.com/OnionWorker/huibaotong/conf"
	"github.com/OnionWorker/huibaotong/entity"
	"github.com/OnionWorker/huibaotong/lib"
	"github.com/mahonia"
	"github.com/thinkoner/openssl"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"
)

/**
	测试SDK入口类传参使用方法
 */
func TestHuiBaoSdk(t *testing.T){
	HuiBaoSDK,Err := huibaotong.NewHuiBao(conf.H5PAY,"94C49D3629584B918A06BE9F")
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
	HuiBaoSDK,Err := huibaotong.NewHuiBao(conf.H5PAY,"94C49D3629584B918A06BE9F")
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
		Timestamp:"20191115122001",
		SignType:"MD5",
	}
	//业务参数
	PayData := &entity.PcQrPayEntity{
		OutTradeNo:"256522225fgfSCCS",
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

//json 解析测试
func TestJson(t *testing.T){
	var str string = `{\"out_trade_no\":\"256522225fgfg25SSaa\",\"auth_type\":\"2\",\"auth_detail_type\":\"3\",\"auth_data_info\":\"3ad5db688a251024697bbe93fc25d363272ad7ca6a04fbbfa15614a290653f8f9312e6225d33d76a9681eccaa5a70b2353d1950285d738f5a035ee6b881a01f9a02b8a9c4541dc42a5003fc7113b10770a8b4db65076653c8e67d58e6290cc5b7ff1920558d309491c91d038a50bbe23eef7dbc9aea56ac97facc4672171451421fbc30785eb14e8ed4232c40d50d3bcca29a8e8599de94e\"}`
	var maps map[string]interface{}
	err:=json.Unmarshal([]byte(str),&maps)
	fmt.Println(err)
	fmt.Println(maps)
}
//鉴权测试
func TestAuth(t *testing.T){
	HuiFuBaoSDK,Err := huibaotong.NewHuiBao(conf.AUTHSUB,"2C5BA54B58654699853D7B9B")
	if Err != nil{
		fmt.Println(Err)
		return
	}
	//公共参数
	AuthCommon := &entity.NewCommonEntity{
		Version:"1.0",
		Method:"heemoney.user.auth.submit",
		AppId:"hyp191115119487000020781A1355F84",
		MchUid:"1194872122863",
		Charset:"UTF-8",
		Timestamp:"20191118110001",
		SignType:"MD5",
	}
	//鉴权的私有业务参数
	AuthEntity := &entity.AuthEntity{
		OutTradeNo:"256522225fgfg25SSaa",
		AuthType:"2",
		AuthDetailType:"3",
		//AuthDataInfo:"",
	}
	//组合银行卡四要素验证数据[{"bank_card_type":"1","auth_bank_card":"6217000130000751966","auth_name":"张三"}]
	var CardData map[string]string = make(map[string]string)
	CardData["bank_card_type"] = "1"  //类型 1银行卡 2信用卡
	CardData["auth_bank_card"] = "6217000130000751966" //卡号
	CardData["auth_id_card"] = "320926195511175276" //身份证号
	CardData["auth_name"] = "张三"  //卡主姓名
	CardData["auth_mobile"] = "13811111111" //卡手机号码
	var CardArr []map[string]string
	CardArr = append(CardArr,CardData)
	jsonByte,err := json.Marshal(CardArr)
	fmt.Println(err)
	fmt.Println(string(jsonByte))
	enc:=mahonia.NewEncoder("gbk")
	var str string = enc.ConvertString(string(jsonByte))
	key :=enc.ConvertString("031EE29ABD2F4C609CD2A5CF")
	desString,err :=lib.DesEncode([]byte(str),[]byte(key))
	fmt.Println(err)
	fmt.Println(desString)
	AuthEntity.AuthDataInfo = desString
	bizContent,err := json.Marshal(AuthEntity)
	fmt.Println("=================================")
	fmt.Println(string(bizContent))
	fmt.Println(err)
	AuthCommon.BizContent = string(bizContent)
	Result := HuiFuBaoSDK.SetEntity(AuthCommon).Excute(conf.JSONRETURN)
	t.Log(Result)
}
func Test3des(t *testing.T){
	var str string = `[{"auth_bank_card":"6217000130000751966","auth_id_card":"320926195511175276","auth_mobile":"13811111111","auth_name":"张三","bank_card_type":"1"}]`
	desString := lib.ThriDESEnCrypt([]byte(str),[]byte("7E1682A2CDDD456D97F9EED0"))
	fmt.Println(hex.EncodeToString(desString))
}
func UTF82GB2312(s []byte)([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func DesEncodes(src []byte,key []byte)([]byte,error){
	Res,Err := openssl.Des3ECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if Err != nil{
		return nil,Err
	}
	return Res,nil
}
//hexstring
/*
	 public static String byteArr2HexStr(byte[] arrB) {
        int iLen = arrB.length;
        // 每个byte用两个字符才能表示，所以字符串的长度是数组长度的两倍
        StringBuffer sb = new StringBuffer(iLen * 2);
        for (int i = 0; i < iLen; i++) {
            int intTmp = arrB[i];
            // 把负数转换为正数
            while (intTmp < 0) {
                intTmp = intTmp + 256;
            }
            // 小于0F的数需要在前面补0
            if (intTmp < 16) {
                sb.append("0");
            }
            sb.append(Integer.toString(intTmp, 16));
        }
        // 最大128位
        String result = sb.toString();
        return result;
    }
*/
func getHexString()string{
	return ""
}

func Test3des3(t *testing.T){
	var str string = `[{"auth_bank_card":"6217000130000751966","auth_id_card":"320926195511175276","auth_mobile":"13811111111","auth_name":"张三","bank_card_type":"1"}]`
	strByt,_ :=UTF82GB2312([]byte(str))
	key,_ :=UTF82GB2312([]byte("7E1682A2CDDD456D97F9EED0"))
	desByte,err := DesEncodes(strByt,key)
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(desByte))
}

func Test3des4(t *testing.T){
	enc:=mahonia.NewEncoder("gbk")
	var str string = enc.ConvertString(`[{"auth_bank_card":"6217000130000751966","auth_id_card":"320926195511175276","auth_mobile":"13811111111","auth_name":"张三","bank_card_type":"1"}]`)
	key :=enc.ConvertString("7E1682A2CDDD456D97F9EED0")
	desByte,err := DesEncodes([]byte(str),[]byte(key))
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(desByte))
}

func Test3des2(t *testing.T){
	var str string = `[{"auth_bank_card":"6217000130000751966","auth_id_card":"320926195511175276","auth_mobile":"13811111111","auth_name":"张三","bank_card_type":"1"}]`
	desString,err :=lib.DesEncode([]byte(str),[]byte("7E1682A2CDDD456D97F9EED0"))
	fmt.Println(err)
	fmt.Println(desString)
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