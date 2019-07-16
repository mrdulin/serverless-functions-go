package repositories

import (
	adChannelModels "serverless-functions-go/domain/models/adChannel"
	"serverless-functions-go/domain/models/cedar/campaign"
)

type CampaignResultRepository interface {
	UpdateStatusTransaction(row adChannelModels.AdPerformanceReportRow, status campaign.CampaignChannelStatus) error
}
