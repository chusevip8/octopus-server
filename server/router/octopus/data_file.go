package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type DataFileRouter struct{}

func (d *DataFileRouter) InitDataFileRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	dataFileRouter := Router.Group("dataFile")
	dataFileApi := v1.ApiGroupApp.OctopusApiGroup.DataFileApi
	{
		dataFileRouter.POST("upload", dataFileApi.UploadFile)
	}
}
