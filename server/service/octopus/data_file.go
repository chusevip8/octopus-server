package octopus

import (
	"bufio"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"mime/multipart"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type DataFileService struct{}

func (dataFileService *DataFileService) Upload(setupId string, file example.ExaFileUploadAndDownload) error {
	err := dataFileService.saveToDB(setupId, file)
	if err == nil {
		err = global.GVA_DB.Model(&octopus.GenericTaskSetup{}).Where("id = ?", setupId).Updates(map[string]interface{}{
			"data_file":      file.Name,
			"data_file_path": file.Url,
		}).Error
	}
	if err != nil {
		_ = taskBindDataServiceApp.DeleteTaskBindDataBySetupId(setupId, "generic")
		_ = os.Remove(file.Url)
	}
	return err
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
func (dataFileService *DataFileService) saveToDB(setupId string, file example.ExaFileUploadAndDownload) (err error) {
	txt, err := os.Open(file.Url)
	if err != nil {
		return
	}
	defer txt.Close()
	scanner := bufio.NewScanner(txt)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		bindData := octopus.TaskBindData{}
		bindData.MainTaskType = "generic"
		sId, err := strconv.ParseUint(setupId, 10, 0)
		if err != nil {
			return err
		}
		bindData.TaskSetupId = uint(sId)
		v := reflect.ValueOf(&bindData).Elem()
		for i, word := range words {
			if i >= 10 {
				break
			}
			field := v.FieldByName(fmt.Sprintf("Item%d", i+1))
			if field.IsValid() && field.CanSet() {
				field.SetString(word)
			}
		}
		err = taskBindDataServiceApp.CreateTaskBindData(&bindData)
		if err != nil {
			return err
		}
	}
	// 检查扫描过程中是否有错误
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
