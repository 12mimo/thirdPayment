package wecaht

import (
	"thirdPayment/wecaht/callback"
	"thirdPayment/wecaht/extend"
	"thirdPayment/wecaht/merchant"
	"thirdPayment/wecaht/models"
	"thirdPayment/wecaht/pay"
	"thirdPayment/wecaht/secure"
	"thirdPayment/wecaht/service"
)

type Client struct {
	Pay      *pay.Service
	Extend   *extend.Handler
	Merchant *merchant.Service
	Secure   *secure.Service
	Service  *service.Manager
	Config   *models.Config
	Callback *callback.Callback
}
