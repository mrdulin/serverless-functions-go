package application

import (
	"fmt"
	"serverless-functions-go/domain/models"
	"serverless-functions-go/domain/services"
	"serverless-functions-go/domain/services/adChannel/reports/adPerformance"
	"serverless-functions-go/infrastructure/config"
)

type IAdPerformanceReportUseCase interface {
	Get() error
}

type AdPerformanceReportUseCase struct {
	campaignService       services.ICampaignService
	campaignResultService services.ICampaignResultService
	googleAccountService  services.IGoogleAccountService
	appConfig             *config.ApplicationConfig
}

func NewAdPerformanceReportUseCase(
	campaignService services.ICampaignService,
	campaignResultService services.ICampaignResultService,
	googleAccountService services.IGoogleAccountService,
	appConfig *config.ApplicationConfig,
) IAdPerformanceReportUseCase {
	return &AdPerformanceReportUseCase{
		campaignService,
		campaignResultService,
		googleAccountService,
		appConfig,
	}
}

func (uc *AdPerformanceReportUseCase) Get() error {
	googleAccounts, err := uc.googleAccountService.FindGoogleAccountsForReport()
	if err != nil {
		return err
	}

	googleCampaignIds, err := uc.campaignService.FindValidGoogleCampaignIds()
	if err != nil {
		if err, ok := err.(*models.AppError); ok {
			fmt.Printf("%v\n", err)
			return nil
		}
		return err
	}

	for _, googleAccount := range googleAccounts {
		options := adPerformance.AdPerformanceReportServiceOptions{
			ClientCustomerId:      googleAccount.GoogleAdwordsClientCustomerId,
			RefreshToken:          googleAccount.GoogleAccountRefreshToken,
			BaseURL:               uc.appConfig.AdChannelApi,
			CampaignResultService: uc.campaignResultService,
		}
		adPerformanceService := adPerformance.NewAdPerformanceReportService(options)
		reportDefinition := adPerformanceService.FormReportDefinition(googleCampaignIds)
		report, err := adPerformanceService.Get(reportDefinition)
		if err != nil {
			fmt.Printf("update status for google account customer id = %d error", googleAccount.GoogleAdwordsClientCustomerId)
			return err
		}
		reportRows := report.GetRows()
		err = adPerformanceService.UpdateStatusTransaction(reportRows)
		if err != nil {
			return err
		}
		fmt.Printf("update status for google account customer id = %d transaction done\n", googleAccount.GoogleAdwordsClientCustomerId)
	}

	return nil
}
