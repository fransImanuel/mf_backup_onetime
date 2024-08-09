package dto

import "errors"

type FilterBaseDto struct {
	FilterID    int    `json:"filter_id" form:"filter_id"`
	IsLookup    bool   `json:"is_lookup" form:"is_lookup"`
	SearchText  string `json:"search_text" form:"search_text" example:"Search example name,email or code"`
	OrderField  string `json:"order_field" form:"order_field" example:"id|desc"`
	FilterPage  int    `json:"filter_page" form:"filter_page" example:"1"`
	FilterLimit int    `json:"filter_limit" form:"filter_limit" example:"10"`
}

type LookupBaseDto struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ResponseBaseDto struct {
	Code         int    `json:"code"`
	Message      string `json:"message,omitempty"`
	MessageError string `json:"message_error,omitempty"`
}

type ResponseAPIDto struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Count   int64       `json:"count"`
	Data    interface{} `json:"data"`
}

type UploadBase64Dto struct {
	Filename    string
	Extension   string
	ContentType string
	Base64      string
	ResultPath  string
}

func (d *UploadBase64Dto) Validate(operation string) error {
	switch operation {
	case "upload":
		if d.Filename == "" {
			return errors.New("file name required")
		}
		if d.Extension == "" {
			return errors.New("extension required")
		}
		if d.ContentType == "" {
			return errors.New("content type required")
		}
		if d.Base64 == "" {
			return errors.New("base64 required")
		}

	case "delete":
		if d.Filename == "" {
			return errors.New("file name required")
		}
		if d.Extension == "" {
			return errors.New("extension required")
		}
	}

	return nil
}

type CallAPIDto struct {
	Method       string
	Url          string
	ContentType  string
	Headers      map[string]interface{}
	BodyRequest  string
	BodyResponse string
	HttpCode     int
}

func (d *CallAPIDto) Validate() error {
	if d.Method == "" {
		return errors.New("method required")
	}
	if d.Url == "" {
		return errors.New("url required")
	}

	return nil
}

type ExportBase64ResponseDto struct {
	ResponseBaseDto
	Extension string `json:"extension"`
	Base64    string `json:"base_64"`
}
