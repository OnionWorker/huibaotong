package test

import (
	"encoding/xml"
	"fmt"
	"testing"
)
type ReturnXml struct {
	XMLName    xml.Name `xml:"root"`
	RetCode  int   `xml:"ret_code"`
	RetMsg   string   `xml:"ret_msg"`
	RedirectUrl   string   `xml:"redirectUrl"`
	Sign   string   `xml:"sign"`
}

func TestXmlPare(t *testing.T){
	xmlstr := `<?xml version="1.0" encoding="utf-8"?>
					<root>
						<ret_code>0000</ret_code>
                        <ret_msg></ret_msg>
						<redirectUrl><![CDATA[https://hykjh5.heemoney.com/DirectPay/WxPayment.aspx?stid=H19110813184771G_20900e7010eea9636687365b2be97f69]]></redirectUrl>
						<sign><![CDATA[8bb9400fc87bd22717d398827a94f4e8]]></sign>
					</root>`
	var result ReturnXml
	err2 := xml.Unmarshal([]byte(xmlstr), &result)
	fmt.Println("xml解析后的内容：")
	fmt.Println(result)
	fmt.Println(err2)

}
