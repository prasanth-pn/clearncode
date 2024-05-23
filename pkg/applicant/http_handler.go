package applicant

import "github.com/gin-gonic/gin"

func (h *Handler) MountRoutes(engine *gin.Engine) {
	applicantApi:=engine.Group(basePath)


	applicantApi.POST("/basic-info",h.ApplicantBasicInfo)

}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler)ApplicantBasicInfo(c *gin.Context){
	
}