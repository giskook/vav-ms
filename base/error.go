package base

const (
	HTTP_OK                                         string = "0"
	HTTP_ACCEPTED_AV_REQUIRE                        string = "20200"
	HTTP_ACCEPTED_AV_WAIT                           string = "20201"
	HTTP_ACCEPTED_AV_FORMAT_NOT_SET                 string = "20202"
	HTTP_ACCEPTED_AV_ACCESS_ADDR_NOT_SET            string = "20203"
	HTTP_ACCEPTED_AV_STREAM_MEDIA_NOT_SET           string = "20204"
	HTTP_BAD_REQUEST_DECODE                         string = "40000"
	HTTP_BAD_REQUEST_MISSING                        string = "40001"
	HTTP_INTERNAL_SERVER_ERROR                      string = "50000"
	HTTP_INTERNAL_SERVER_ERROR_SET_STREAM_MEDIA     string = "50001"
	HTTP_INTERNAL_SERVER_ERROR_GET_STREAM_MEDIA     string = "50002"
	HTTP_INTERNAL_SERVER_ERROR_DEL_STREAM_MEDIA     string = "50003"
	HTTP_INTERNAL_SERVER_ERROR_UPD_STREAM_MEDIA     string = "50004"
	HTTP_INTERNAL_SERVER_ERROR_SET_ACCESS_ADDR      string = "50005"
	HTTP_INTERNAL_SERVER_ERROR_GET_ACCESS_ADDR      string = "50006"
	HTTP_INTERNAL_SERVER_ERROR_SET_VEHICLE_PROPERTY string = "50007"
	HTTP_INTERNAL_TIMEOUT                           string = "50400"
)

var ErrorMap map[string]string = map[string]string{
	HTTP_OK: "成功",

	HTTP_ACCEPTED_AV_REQUIRE:                        "请下发1078音视频指令",
	HTTP_ACCEPTED_AV_WAIT:                           "1078音视频指令已下发,请等待",
	HTTP_ACCEPTED_AV_FORMAT_NOT_SET:                 "车机音视频格式未设置",
	HTTP_ACCEPTED_AV_ACCESS_ADDR_NOT_SET:            "车机接入地址未设置",
	HTTP_ACCEPTED_AV_STREAM_MEDIA_NOT_SET:           "流媒体地址未设置",
	HTTP_BAD_REQUEST_DECODE:                         "参数解析出错",
	HTTP_BAD_REQUEST_MISSING:                        "缺少参数",
	HTTP_INTERNAL_SERVER_ERROR:                      "服务器内部错误",
	HTTP_INTERNAL_SERVER_ERROR_SET_STREAM_MEDIA:     "添加流媒体服务器地址失败",
	HTTP_INTERNAL_SERVER_ERROR_GET_STREAM_MEDIA:     "读取流媒体服务器地址失败",
	HTTP_INTERNAL_SERVER_ERROR_DEL_STREAM_MEDIA:     "删除流媒体服务器地址失败",
	HTTP_INTERNAL_SERVER_ERROR_UPD_STREAM_MEDIA:     "更新流媒体服务器地址失败",
	HTTP_INTERNAL_SERVER_ERROR_SET_ACCESS_ADDR:      "设置车机连接地址失败",
	HTTP_INTERNAL_SERVER_ERROR_GET_ACCESS_ADDR:      "得到车机连接地址失败",
	HTTP_INTERNAL_SERVER_ERROR_SET_VEHICLE_PROPERTY: "设置车机音视频属性失败",
	HTTP_INTERNAL_TIMEOUT:                           "超时",
}

var (
	ERROR_BAD_REQUEST_MISSING = NewErr(nil, HTTP_BAD_REQUEST_MISSING)
)

type VavmsError struct {
	Err      error
	Code     string
	Describe string
}

func (e *VavmsError) Error() string {
	if e.Err != nil {
		return e.Describe + "(" + e.Code + ")" + e.Err.Error()
	} else {
		return e.Describe + "(" + e.Code + ")"
	}
}

func NewErr(err error, code string) *VavmsError {
	return &VavmsError{
		Err:      err,
		Code:     code,
		Describe: ErrorMap[code],
	}
}
