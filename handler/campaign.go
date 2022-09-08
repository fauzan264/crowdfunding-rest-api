package handler

import (
	"net/http"

	"github.com/fauzan264/crowdfunding-rest-api/campaign"
	"github.com/fauzan264/crowdfunding-rest-api/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId := c.Query("user_id")

	if len(userId) != 0 {
		_, err := uuid.Parse(userId)
		if err != nil {
			response := helper.APIResponse("Error to get Campaigns, User ID not valid", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	campaigns, err := h.service.GetCampaigns(userId)

	if err != nil {
		response := helper.APIResponse("Error to Get Campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}
