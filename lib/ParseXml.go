package lib

import "encoding/xml"

type ReturnXml struct {
	XMLName    xml.Name `xml:"root"`
	RetCode  string   `xml:"ret_code"`
	RetMsg   string   `xml:"ret_msg"`
	RedirectUrl   string   `xml:"redirectUrl"`
	Sign   string   `xml:"sign"`
}

func ParseXml(xmlstr string)(*ReturnXml,error){
	var result ReturnXml
	err := xml.Unmarshal([]byte(xmlstr), &result)
	if err != nil{
		return nil,err
	}
	return &result,nil
}
