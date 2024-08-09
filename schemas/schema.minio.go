package schemas

import "errors"

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
