package repositories

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"serverless-functions-go/domain/models/cedar"
	"serverless-functions-go/domain/repositories"
)

type GoogleAccountRepository struct {
	Db *sqlx.DB
}

func NewGoogleAccountRepository(Db *sqlx.DB) repositories.GoogleAccountRepository {
	return &GoogleAccountRepository{Db}
}

func (repo *GoogleAccountRepository) FindByClientCustomerIds(ids []int) ([]cedar.GoogleAccount, error) {
	var googleAccounts = make([]cedar.GoogleAccount, 0)

	query := `
		SELECT
			gac.google_account_refresh_token,
			gac.google_account_user_nme,
			gaw.google_adwords_id,
			gaw.google_adwords_customer_nme,
			gaw.google_adwords_client_customer_id
		FROM
			"GOOGLE_ADWORDS" as gaw
		INNER JOIN "GOOGLE_ACCOUNT" as gac ON gaw.google_account_id = gac.google_account_id
		WHERE 
			gaw.google_adwords_client_customer_id IN (?);
	`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return googleAccounts, errors.Wrap(err, "sqlx.In")
	}
	query = repo.Db.Rebind(query)
	err = repo.Db.Select(&googleAccounts, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("err = %v\n", sql.ErrNoRows)
			return googleAccounts, nil
		}
		return googleAccounts, errors.Wrap(err, "googleAccountRepo.Db.Select error")
	}

	return googleAccounts, nil
}

func (repo *GoogleAccountRepository) FindByCampaignRanByZOWIForZELO() ([]cedar.GoogleAccount, error) {
	var googleAccounts []cedar.GoogleAccount

	query := `
		select
			distinct on (ga.google_account_id)
			ga.* 
		from "CAMPAIGN" as c
		inner join "LOCATION" as loc on loc.location_id = c.location_id
		inner join "ORGANIZATION" as org on loc.organization_id = org.organization_id
		inner join "GOOGLE_ACCOUNT" as ga on ga.organization_id = org.parent_id
		where c.campaign_ran_by_zowi = true;
	`
	err := repo.Db.Select(&googleAccounts, query)
	if err != nil {
		return googleAccounts, errors.Wrap(err, "googleAccountRepo.Db.Select error")
	}
	return googleAccounts, nil
}
