package repositories

import (
	"serverless-functions-go/domain/models/cedar"
)

type GoogleAccountRepository interface {
	FindByClientCustomerIds(ids []int) ([]cedar.GoogleAccount, error)
	FindByCampaignRanByZOWIForZELO() ([]cedar.GoogleAccount, error)
}
