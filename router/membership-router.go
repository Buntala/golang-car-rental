package router

import (
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	membershipRepository repository.MembershipRepository = repository.NewMembershipRepository()
	membershipService    service.MembershipService       = service.NewMembershipService(membershipRepository)
	membershipController controller.MembershipController = controller.NewMembership(membershipService)
)

func Membership(router *gin.Engine) {
	member_r := router.Group("/memberships")
	{
		member_r.GET("/", membershipController.FindAll)
		member_r.GET("/:id", membershipController.FindOne)
		member_r.POST("/",  membershipController.Save)
		member_r.PATCH("/:id",  membershipController.Update)
		member_r.DELETE("/:id",  membershipController.Delete)
	}
}
