package services

import (
	"fmt"
	adChannelModels "serverless-functions-go/domain/models/adChannel"
	"sync"

	googleChannelAd "serverless-functions-go/domain/models/adChannel/ad"
	googleChannelAdGroup "serverless-functions-go/domain/models/adChannel/adGroup"
	googleChannelCampaign "serverless-functions-go/domain/models/adChannel/campaign"
	cedarCampaign "serverless-functions-go/domain/models/cedar/campaign"

	"serverless-functions-go/domain/repositories"
)

type ICampaignResultService interface {
	UpdateStatusTransaction(rows []adChannelModels.AdPerformanceReportRow) error
}

type CampaignResultService struct {
	campaignResultRepo repositories.CampaignResultRepository
}

func NewCampaignResultService(campaignResultRepo repositories.CampaignResultRepository) ICampaignResultService {
	return &CampaignResultService{campaignResultRepo}
}

func (svc *CampaignResultService) UpdateStatusTransaction(rows []adChannelModels.AdPerformanceReportRow) error {
	wg := sync.WaitGroup{}
	wg.Add(len(rows))
	// TODO: batch update
	for _, row := range rows {
		var campaignChannelStatus cedarCampaign.CampaignChannelStatus = ""

		if row.ApprovalStatus == googleChannelAd.APPROVED {

			if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.SUCCESS
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.PAUSED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.FAILED
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.PAUSED {
				campaignChannelStatus = cedarCampaign.FAILED
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.REMOVED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.FAILED
				// TODO: ad doesn't have "removed" status
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.DISABLED {
				campaignChannelStatus = cedarCampaign.FAILED
			} else if row.CampaignState == googleChannelCampaign.PAUSED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.SUCCESS
			} else if row.CampaignState == googleChannelCampaign.REMOVED {
				campaignChannelStatus = cedarCampaign.FAILED
			}
			// TODO: campaign ended

		} else if row.ApprovalStatus == googleChannelAd.DISAPPROVED {
			if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.FAILED
			}
		}

		if campaignChannelStatus == "" {
			return fmt.Errorf("invalid campaign channel status = %s", campaignChannelStatus)
		}

		go func(row adChannelModels.AdPerformanceReportRow) {
			defer wg.Done()
			err := svc.campaignResultRepo.UpdateStatusTransaction(row, campaignChannelStatus)
			if err != nil {
				fmt.Printf("update status transaction error for row = %+v. error = %v\n", row, err)
			}
		}(row)
	}
	wg.Wait()
	return nil
}
