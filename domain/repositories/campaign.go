package repositories

import (
	"serverless-functions-go/domain/models/cedar/campaign"
)

type CampaignRepository interface {
	FindById(id string) (*campaign.Campaign, error)
	FindValidGoogleCampaign() ([]*campaign.Campaign, error)
}
