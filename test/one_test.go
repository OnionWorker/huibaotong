package test

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/OnionWorker/huibaotong/entity"
	"time"
)

func TestEntiyOne(t *testing.T){
	var EntityOne *entity.PayEntity = &entity.PayEntity{
		PayType:1,
	}
	EntityOne.Version = 1
	var TestData entity.Entity = EntityOne
	fmt.Print(TestData.GetSign())
	t.Log(EntityOne)
}

func TestGetEntityTag(t *testing.T){
	var EntityOne entity.PayEntity = entity.PayEntity{
		PayType:1,
	}
	EntityOne.Version = 1
	var TestInterface interface{} = EntityOne
	ts := reflect.TypeOf(TestInterface)
	vs := reflect.ValueOf(TestInterface)
	for i := 0; i < ts.NumField(); i++ {
		field := ts.Field(i)
		fmt.Printf("是否机构体:%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name(), field.Type.Kind() == reflect.Struct)
		fmt.Printf("值打印:%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name(), vs.Field(i).Interface())
		tag := field.Tag.Get("json")
		fmt.Printf("Tag打印:%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}

	t.Log(EntityOne)
}

func TestTimes(t *testing.T){
	fmt.Println(time.Now().Format("20060102010500"))  //2019-07-31
}

func TestMoney(t *testing.T){
	var money float64 = 100.20
	var fen float64 = 100
	ff := money*fen
	var  aa int = int(ff)
	fmt.Println(aa)
}