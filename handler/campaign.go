package handler

import (
	"net/http"
	"strconv"

	"github.com/codingsluv/crowdfounding/campaign"
	"github.com/codingsluv/crowdfounding/helper"
	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// * /api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

	// userIDStr := c.Query("users_id")

	// userID, err := strconv.Atoi(userIDStr)
	// if err != nil {
	// 	response := helper.ApiResponse("Invalid user ID", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// campaigns, err := h.service.GetCampaigns(userID)
	// if err != nil {
	// 	response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// response := helper.ApiResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	// c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)

}
