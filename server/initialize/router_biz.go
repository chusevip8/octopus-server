package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	{
		octopusRouter := router.RouterGroupApp.Octopus
		octopusRouter.InitScriptRouter(privateGroup, publicGroup)
		octopusRouter.InitDeviceRouter(privateGroup, publicGroup)

		octopusRouter.InitCmtThreadRouter(privateGroup, publicGroup)
		octopusRouter.InitCmtConversationRouter(privateGroup, publicGroup)
		octopusRouter.InitCommentRouter(privateGroup, publicGroup)
		octopusRouter.InitTaskRouter(privateGroup, publicGroup)

		octopusRouter.InitCmtTaskSetupRouter(privateGroup, publicGroup)

		octopusRouter.InitCmtTaskRouter(privateGroup, publicGroup)
		octopusRouter.InitTaskParamsRouter(privateGroup, publicGroup)
		octopusRouter.InitIntervalTaskSetupRouter(privateGroup, publicGroup)
		octopusRouter.InitIntervalTaskRouter(privateGroup, publicGroup)
		octopusRouter.InitGenericTaskSetupRouter(privateGroup, publicGroup)
		octopusRouter.InitDataFileRouter(privateGroup, publicGroup)
		octopusRouter.InitGenericTaskRouter(privateGroup, publicGroup)
		octopusRouter.InitTaskBindDataRouter(privateGroup, publicGroup)
		octopusRouter.InitMsgTaskSetupRouter(privateGroup, publicGroup)

	}

	holder(publicGroup, privateGroup)
}
