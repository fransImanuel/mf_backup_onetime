package schemas

type SchemaResponses struct {
	Code    int         `json:"code"`
	Method  string      `json:"method,omitempty"`
	Message string      `json:"message"`
	Count   *int64      `json:"count,omitempty"`
	Data    interface{} `json:"data"`
	//Items   interface{} `json:"items,omitempty"`
}
