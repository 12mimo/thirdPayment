package merchant

type Service interface {
	SubmitRequisition(params map[string]interface{}) (interface{}, error) // 特约商户进件 提交申请单
	Request()
}

// 特约商户进件 提交申请单 https://pay.weixin.qq.com/doc/v3/partner/4012719997

func Request() {

}
