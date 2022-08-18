package membership

import "github.com/gin-gonic/gin"

func Membership(router *gin.Engine) {
	membershipRouter := router.Group("/memberships")
	membershipRouter.GET("/", getMembership)
	membershipRouter.GET("/:id", getMembershipById)
	membershipRouter.POST("/", postMembership)
	membershipRouter.PATCH("/:id", patchMembership)
	membershipRouter.DELETE("/:id", deleteMembership)
}