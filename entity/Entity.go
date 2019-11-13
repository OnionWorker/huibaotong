package entity


type Entity interface {
	GetSign()[]string     //自动签名的数组(顺序与官方签名顺序一致)
}

type ConmmonEntity struct {
	Version    int       `json:"version"` //当前接口版本号1
	Signs      string      `json:"sign"` //MD5签名结果
}


