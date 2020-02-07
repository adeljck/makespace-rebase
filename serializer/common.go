package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

//PURE ERROR Resopnse
type PureErrorResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

//validate error
type TagError struct {
	Tag   string `json:"tag"`
	Error string `json:"error"`
}

// BuildListResponse 列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
