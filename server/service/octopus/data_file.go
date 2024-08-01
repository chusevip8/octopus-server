package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"mime/multipart"
	"strings"
)

type DataFileService struct{}

func (dataFileService *DataFileService) Upload(setupId string, file example.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Model(&octopus.GenericTaskSetup{}).Where("id = ?", setupId).Updates(map[string]interface{}{
		"data_file":      file.Name,
		"data_file_path": file.Url,
	}).Error
}

func (dataFileService *DataFileService) UploadFile(setupId string, header *multipart.FileHeader) (file example.ExaFileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := example.ExaFileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	return f, dataFileService.Upload(setupId, f)
}
