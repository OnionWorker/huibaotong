package lib

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/thinkoner/openssl"
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
		if isStruct(reflect.TypeOf(x))|| isStructPtr(reflect.TypeOf(x)){
			jsonByte,err := json.Marshal(x)
			fmt.Println(err)
			str = string(jsonByte)
		}else{
			str = fmt.Sprintf("%d", x)
		}
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
	fmt.Println(string(bytesData))
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

//3DES加密处理
func DesEncode(src []byte,key []byte)(string,error){
	Res,Err := openssl.Des3ECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if Err != nil{
		return "",Err
	}

	return strings.ToUpper(hex.EncodeToString(Res)),nil
	//return string(Res),nil
}

//解密
func ThriDESDeCrypt(crypted,key []byte)[]byte{
	//获取block块
	block,_ :=des.NewTripleDESCipher(key)
	//创建切片
	context := make([]byte,len(crypted))
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block,key[:8])
	//解密密文到数组
	blockMode.CryptBlocks(context,crypted)
	//去补码
	context = PKCSUnPadding(context)
	return context
}
//去补码
func PKCSUnPadding(origData []byte)[]byte{
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}
//加密
func ThriDESEnCrypt(origData []byte,key []byte)[]byte{
	//获取block块
	block,_ :=des.NewTripleDESCipher(key)
	//补码
	origData = PKCS7Padding(origData, block.BlockSize())
	//设置加密方式为 3DES  使用3条56位的密钥对数据进行三次加密
	blockMode := cipher.NewCBCEncrypter(block,key[:8])
	//创建明文长度的数组
	crypted := make([]byte,len(origData))
	//加密明文
	blockMode.CryptBlocks(crypted,origData)
	return crypted
}
//补码
func PKCSPadding(origData []byte,blockSize int)[]byte{
	//计算需要补几位数
	padding := blockSize-len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(origData,padtext...)
}
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5Padding(cipherText []byte) []byte {
	return PKCS7Padding(cipherText, 8)
}
