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

		octopusRouter.InitCmtTaskParamsRouter(privateGroup, publicGroup)
		octopusRouter.InitCmtTaskSetupRouter(privateGroup, publicGroup)

	}

	holder(publicGroup, privateGroup)
}
