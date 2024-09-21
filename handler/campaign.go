package handler

import (
	"net/http"
	"strconv"

	"github.com/codingsluv/crowdfounding/campaign"
	"github.com/codingsluv/crowdfounding/helper"
	"github.com/gin-gonic/gin"
)

// * tangkap parameter di handler
// * handler ke service
// * service menentukan repository mana yg akan di panggil
// * repository : FindAll FindByUserID
// * db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// * /api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("users_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
