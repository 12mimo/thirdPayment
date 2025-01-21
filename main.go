package thirdopen

import (
	"thirdPayment/wecaht"
	"thirdPayment/zhifubao"
)

type Client struct {
	Wechat *wecaht.Client
	ZFB    *zhifubao.Client
}
