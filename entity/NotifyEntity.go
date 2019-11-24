package entity

type NotityEntity struct {
	Version	    string  `json:"version" form:"version"`
	AppId	    string  `json:"app_id" form:"app_id"`
	MchUid	    string  `json:"mch_uid" form:"mch_uid"`
	IsvMchUid	string  `json:"isv_mch_uid" form:"isv_mch_uid"`
	IsvAppId	string  `json:"isv_app_id" form:"isv_app_id"`
	Subject	    string  `json:"subject" form:"subject"`
	OutTradeNo	string  `json:"out_trade_no" form:"out_trade_no"`
	HyBillNo	string  `json:"hy_bill_no" form:"hy_bill_no"`
	ChannelType	string  `json:"channel_type" form:"channel_type"`
	TotalFee	string  `json:"total_fee" form:"total_fee"`
	RealFee	    string   `json:"real_fee" form:"real_fee"`
	TradeStatus	string  `json:"trade_status" form:"trade_status"`
	TimeEnd	    string  `json:"time_end" form:"time_end"`
	Attach	    string  `json:"attach" form:"attach"`
	MetaOption	string  `json:"meta_option" form:"meta_option"`
	PayOption	string  `json:"pay_option" form:"pay_option"`
	Sign	    string  `json:"sign" form:"sign"`
}
