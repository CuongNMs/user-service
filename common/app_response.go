package common

type AppResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewAppResponse(data, paging, filter interface{}) *AppResponse {
	return &AppResponse{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func NewAppResponseSimple(data interface{}) *AppResponse {
	return &AppResponse{
		Data:   data,
		Paging: nil,
		Filter: nil,
	}
}
