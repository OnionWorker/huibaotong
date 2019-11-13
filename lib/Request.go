package lib

import (
	"fmt"
	"github.com/OnionWorker/huibaotong/conf"
	"github.com/OnionWorker/huibaotong/entity"
	"reflect"
	"strings"
)

type Request struct {
	postData map[string]interface{}
	signMap map[string]interface{}
	signDataAtr []string
	signData string
	postUrl string
}

func NewRequest()*Request{
	return &Request{}
}

func (this *Request) Run(config *conf.Config,signDatas []string,entitys entity.Entity,url string,rtype conf.HuiBaoType)(string,error,map[string]interface{}){
	this.postUrl = url
	this.signMap = make(map[string]interface{})
	this.postData = make(map[string]interface{})
	this.signDataAtr = signDatas
	Err := this.getData(entitys)
	if Err != nil{
		return "",Err,nil
	}
	for _,v := range signDatas{
		switch v {
			case "key":
				this.signMap[v] = strings.ToUpper(config.SecretKey)
				break
		}
	}
	this.signData = getSignStr(this.signDataAtr,this.signMap)
	this.postData["sign"] = encodeMD5(this.signData)
	if rtype == conf.XMLRETURN{
		resultString,Err := post(this.postUrl,getPostStr(this.postData))
		return resultString,Err,nil
	}else if rtype == conf.JSONRETURN{
		this.postData["sign"] = strings.ToUpper(this.postData["sign"].(string))
		resultMap := jsonPost(this.postUrl,this.postData)
		return "",nil,resultMap
	}
	return this.postUrl+"?"+getPostStr(this.postData),nil,nil
}

func (this *Request)getData(entitys interface{})error{
	var err error
	ts := reflect.TypeOf(entitys)
	vs := reflect.ValueOf(entitys)
	switch {
	case isStruct(ts):
	case isStructPtr(ts):
		ts = ts.Elem()
		vs = vs.Elem()
	default:
		err = fmt.Errorf("%v must be a struct or a struct pointer", entitys)
		return err
	}
	for i := 0; i < ts.NumField(); i++ {
		field := ts.Field(i)
		tag := field.Tag.Get("json")
		value := vs.Field(i).Interface()
		if vs.Field(i).Kind() == reflect.Ptr {
			if vs.Field(i).IsNil() {
				value = ""
			} else {
				value = vs.Field(i).Elem().Interface()
			}
		}

		if tag == "" && (isStruct(field.Type) || isStructPtr(field.Type)) && value != ""{
			err = this.getData(value)
			continue
		}
		//构造POST参数
		this.postData[tag] = value
		this.getSigns(tag,value)
	}
	return err
}

func (this *Request) getSigns(field string,value interface{}){
	for _,v := range this.signDataAtr{
		if v == field{
			this.signMap[v] = value
		}
	}
}





