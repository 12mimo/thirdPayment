package merchant

const (
	SubmitRequisitionUrl                    = "/v3/applyment4sub/applyment/"                     // 特约商户进件URL
	QueryRequisitionStatusUrl               = "/v3/applyment4sub/applyment/applyment_id/%v/"     // 申请单号查询申请状态URL
	QueryRequisitionStatusByBusinessCodeUrl = "/v3/applyment4sub/applyment/business_code/%v/"    // 业务申请编号查询申请状态URL
	ModifySettlementAccountUrl              = "/v3/apply4sub/sub_merchants/%v/modify-settlement" // 修改结算账户URL
	QuerySettlementAccountUrl               = "/v3/apply4sub/sub_merchants/%v/settlement"        // 查询结算账户URL
	QuerySettlementAccountModifyStatusUrl   = "/v3/apply4sub/sub_merchants/%v/application/%v"    // 查询结算账户修改状态URL
	CancelRequisitionUrl                    = "/v3/apply4subject/applyment/%v/cancel"            // 取消申请单URL
	QueryRequisitionExamineStatusUrl        = "/v3/apply4subject/applyment"                      // 查询申请单审核结果URL
	QueryMerchantsStateUrl                  = "/v3/apply4subject/applyment/merchants/%v/state"   // 获取商户开户意愿确认状态URL
)

type Service struct {
	Uploads *Uploads
}

// SubmitRequisition  特约商户进件 提交申请单 https://pay.weixin.qq.com/doc/v3/partner/4012719997
func (Service) SubmitRequisition(params map[string]interface{}) (interface{}, error) {
	return nil, nil
}

// QueryRequisitionStatus 申请单号查询申请状态 https://pay.weixin.qq.com/doc/v3/partner/4012697052
func (Service) QueryRequisitionStatus() {

}

// QueryRequisitionStatusByBusinessCode 业务申请编号查询申请状态 https://pay.weixin.qq.com/doc/v3/partner/4012697168
func (Service) QueryRequisitionStatusByBusinessCode() {

}

// ModifySettlementAccount 修改结算账户 https://pay.weixin.qq.com/doc/v3/partner/4012761102
func (Service) ModifySettlementAccount() {

}

// QuerySettlementAccount 修改结算账户 https://pay.weixin.qq.com/doc/v3/partner/4012761113
func (Service) QuerySettlementAccount() {

}

// QuerySettlementAccountModifyStatus 修改结算账户 https://pay.weixin.qq.com/doc/v3/partner/4012761120
func (Service) QuerySettlementAccountModifyStatus() {

}

// CancelRequisition 撤销申请单 https://pay.weixin.qq.com/doc/v3/partner/4012697627
func (Service) CancelRequisition() {

}

// QueryRequisitionExamineStatus 查询申请单审核结果 https://pay.weixin.qq.com/doc/v3/partner/4012697715
func (Service) QueryRequisitionExamineStatus() {

}

// QueryMerchantsState 获取商户开户意愿确认状态 https://pay.weixin.qq.com/doc/v3/partner/4012467549
func (Service) QueryMerchantsState() {

}

func (Service) Request() {

}
