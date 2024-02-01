package request

import "mime/multipart"

type FileUpload struct {
	File *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}

func (fileUpload FileUpload) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"file.required": "请选择文件",
	}
}
