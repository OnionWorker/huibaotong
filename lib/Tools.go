package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func toString(elemnt interface{})string{
	var str string
	switch x := elemnt.(type) {
	case int:
		str = strconv.Itoa(x)
	case int32:
		str = strconv.FormatInt(int64(x),10)
	case int64:
		str = strconv.FormatInt(x,10)
	case float64:
		str = strconv.FormatFloat(x, 'f', 2, 64)
	case float32:
		str = strconv.FormatFloat(float64(x), 'f', 2, 64)
	case string:
		str = x
	default:
		str = fmt.Sprintf("%d", x)
	}
	return str
}

func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func  post(url string,postData string) (string,error) {
	resp, err := http.Post(url,"application/x-www-form-urlencoded",strings.NewReader(postData))
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(body),nil
}

func jsonPost(url string,postData map[string]interface{})map[string]interface{} {
	bytesData, err := json.Marshal(postData)
	if err != nil {
		return map[string]interface{}{"return_code":"FAIL","return_msg":err.Error()}
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return map[string]interface{}{"return_code":"FAIL","return_msg":err.Error()}
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return map[string]interface{}{"return_code":"FAIL","return_msg":err.Error()}
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"return_code":"FAIL","return_msg":err.Error()}
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return map[string]interface{}{"return_code":"FAIL","return_msg":err.Error()}
	}
	return result
}

func getSignStr(signDataAtr []string,signMap map[string]interface{})string{
	var signData string
	for i := 0; i < len(signDataAtr); i++{
		elemnt := signMap[signDataAtr[i]]
		var str string = toString(elemnt)
		signData = signData+"&"+signDataAtr[i]+"="+str
	}
	signData = strings.Trim(signData,"&")
	return signData
}

func getPostStr(postData map[string]interface{})string{
	var postStr string
	for k,v := range postData{
		var str string = toString(v)
		postStr = postStr+"&"+k+"="+str
	}
	postStr = strings.Trim(postStr,"&")
	return postStr
}

