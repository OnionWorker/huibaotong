# 汇宝通API封装包

由于我司这些日子需要支付功能，但又不想直接接入WX、AL。所以使用了汇宝通的API来完成公司项目的业务。<br>
总体来说对接中的过程还是让人觉得很混乱的，还是希望官方对文档、以及接口传参、返回值都优化处理下，尽量多一点共性。


包说名（目前包含：微信H5、支付宝H5、PC扫码支付）

调用示例请参考 test/two_test.go 文件中的前三个方法

常量说明
  API请求地址常量
  
  * CUSTOM   自定义API URL
  * H5PAY    H5手机支付 API URL （微信支付宝共用，具体文档：http://dev.heepay.com/index.php?s=/31&page_id=251） 
  * PCQRPAY  pc扫码支付 文档 http://dev.heemoney.com/#/API?id=%e7%bb%9f%e4%b8%80%e4%b8%8b%e5%8d%95
  
  返回类型常量
  
  * XMLRETURN  用于解析XML的返回值（POST content-type application/x-www-form-urlencoded）
  * URLRETURN  当接口直接返回URL地址时必须使用此类型 (直接组装参数返回URL，无需要发起请求）
  * JSONRETURN  当接口返回JSON时 （POST content-type application/json）
  
  调用说明
  ```Go
  HuiBaoSDK,Err := huibaotong.NewHuiBao //(API请求地址常量,参与签名的KEY)
  ......
  Result := HuiBaoSDK.SetEntity(实体struck).Excute //(返回类型常量,可选参数URL)  
  ```
  
  可选参数URL说明
     参数传递需要经过 huibaotong.Url(网址字符串) 包装
     huibaotong.NewHuiBao(此处必须为 CUSTOM,参与签名的KEY)
     
 扩展说明
    
    如果想调用汇付宝其他API，请在构造entity时 实现 entity.Entity 中的 GetSign()[]string 方法
    如需定义常量，请前往 conf/Config.go
    
（如有问题请ISSUES看到就会回复，如想吐槽---- thank you 现在还不用）



  


