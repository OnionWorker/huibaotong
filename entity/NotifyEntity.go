package entity

type NotityEntity struct {
	Version	    string  `json:"version" form:"version"`
	AppId	    string  `json:"app_id" form:"version"`
	MchUid	    string  `json:"mch_uid" form:"version"`
	IsvMchUid	string  `json:"isv_mch_uid" form:"version"`
	IsvAppId	string  `json:"isv_app_id" form:"version"`
	Subject	    string  `json:"subject" form:"version"`
	OutTradeNo	string  `json:"out_trade_no" form:"version"`
	HyBillNo	string  `json:"hy_bill_no" form:"version"`
	ChannelType	string  `json:"channel_type" form:"version"`
	TotalFee	int     `json:"total_fee" form:"version"`
	RealVol	    int     `json:"real_vol" form:"version"`
	TradeStatus	string  `json:"trade_status" form:"version"`
	TimeEnd	    string  `json:"time_end" form:"version"`
	Attach	    string  `json:"attach" form:"version"`
	MetaOption	string  `json:"meta_option" form:"version"`
	PayOption	string  `json:"pay_option" form:"version"`
	Sign	    string  `json:"sign" form:"version"`
}
