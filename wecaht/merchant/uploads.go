package merchant

const (
	ImageUrl = "/v3/merchant/media/upload"       // 上传图片URL
	VideoUrl = "/v3/merchant/media/video_upload" // 上传视频URL
)

type Uploads struct {
}

// Image 上传图片 https://pay.weixin.qq.com/doc/v3/partner/4012760490
func (Uploads) Image() {

}

// Video 上传视频 https://pay.weixin.qq.com/doc/v3/partner/4012761084
func (Uploads) Video() {

}

func (Uploads) Request() {

}
