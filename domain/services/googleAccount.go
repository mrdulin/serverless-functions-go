package services

import (
	"fmt"
	"serverless-functions-go/domain/models/cedar"
	"serverless-functions-go/domain/repositories"
	"serverless-functions-go/infrastructure/utils"
)

type IGoogleAccountService interface {
	FindGoogleAccountsForReport() ([]cedar.GoogleAccountForReport, error)
}

type GoogleAccountService struct {
	googleAccountRepo repositories.GoogleAccountRepository
	locationRepo      repositories.LocationRepository
}

func NewGoogleAccountService(googleAccountRepo repositories.GoogleAccountRepository, locationRepo repositories.LocationRepository) IGoogleAccountService {
	return &GoogleAccountService{googleAccountRepo, locationRepo}
}

func (svc *GoogleAccountService) FindGoogleAccountsForReport() ([]cedar.GoogleAccountForReport, error) {
	var googleAccountsForReport = make([]cedar.GoogleAccountForReport, 0)
	locations, err := svc.locationRepo.FindLocationsBoundGoogleClientCustomerId()
	if err != nil {
		return nil, err
	}
	if len(locations) == 0 {
		return nil, fmt.Errorf("no location binds google adwords client customer id")
	}

	googleAdWordsClientCustomerIds := make([]int, 0)
	for _, location := range locations {
		if location.GoogleAdwordsClientCustomerId != 0 {
			googleAdWordsClientCustomerIds = append(googleAdWordsClientCustomerIds, location.GoogleAdwordsClientCustomerId)
		}
	}

	if len(googleAdWordsClientCustomerIds) == 0 {
		return nil, fmt.Errorf("no google adwords client customer ids")
	}

	googleAccountsForZOWI, err := svc.googleAccountRepo.FindByClientCustomerIds(googleAdWordsClientCustomerIds)
	if err != nil {
		return nil, err
	}
	for _, googleAccountForZOWI := range googleAccountsForZOWI {
		googleAccountsForReport = append(googleAccountsForReport, cedar.GoogleAccountForReport{
			RefreshToken:     googleAccountForZOWI.GoogleAccountRefreshToken,
			ClientCustomerId: googleAccountForZOWI.GoogleAdwordsClientCustomerId,
		})
	}

	googleAccountsForZELO, err := svc.googleAccountRepo.FindByCampaignRanByZOWIForZELO()
	if err != nil {
		return nil, err
	}
	for _, googleAccountForZELO := range googleAccountsForZELO {
		if googleAccountForZELO.GoogleAccountDefaultCustomerId.Valid {
			googleAccountsForReport = append(googleAccountsForReport, cedar.GoogleAccountForReport{
				RefreshToken:     googleAccountForZELO.GoogleAccountRefreshToken,
				ClientCustomerId: googleAccountForZELO.GoogleAccountDefaultCustomerId.String,
			})
		}
	}

	if len(googleAccountsForReport) == 0 {
		return nil, fmt.Errorf("no google accounts for getting report")
	}

	got := utils.UniqueV2(googleAccountsForReport)
	googleAccountsForReportUnique := make([]cedar.GoogleAccountForReport, len(got))
	for idx, val := range got {
		googleAccountsForReportUnique[idx] = val.(cedar.GoogleAccountForReport)
	}

	return googleAccountsForReportUnique, nil
}
