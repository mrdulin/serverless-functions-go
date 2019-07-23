package application

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"serverless-functions-go/domain/services"
	"serverless-functions-go/infrastructure/config"
	"serverless-functions-go/infrastructure/database"
	"serverless-functions-go/infrastructure/gcloud/datastore"
	"serverless-functions-go/interfaces/repositories"
)

type CompositionRoot struct {
	AppConfig *config.ApplicationConfig
	Db        *sqlx.DB

	CampaignService       services.ICampaignService
	CampaignResultService services.ICampaignResultService
	GoogleAccountService  services.IGoogleAccountService

	DataStoreService datastore.IService
}

func NewCompositionRoot() *CompositionRoot {
	projectId := os.Getenv("GCP_PROJECT")
	dataStoreCredentials := os.Getenv("DATASTORE_CREDENTIALS")
	dataStoreOptions := datastore.Options{ProjectID: projectId, CredentialsFile: dataStoreCredentials}
	dataStoreService, err := datastore.New(&dataStoreOptions)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	appConfig, err := config.New("dataStore", dataStoreService)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Printf("env: %#v\n", os.Getenv("ENV"))
	fmt.Printf("application config: %#v\n", appConfig)

	dbConf := database.PGDatabaseConfig{
		Host:     fmt.Sprintf("/cloudsql/%s", appConfig.SqlInstanceConnectionName),
		User:     appConfig.SqlUser,
		Password: appConfig.SqlPassword,
		Dbname:   appConfig.SqlDb,
	}
	db, err := database.ConnectPGDatabase(&dbConf)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// repositories
	campaignResultRepo := repositories.NewCampaignResultRepository(db)
	campaignRepo := repositories.NewCampaignRepository(db)
	googleAccountRepo := repositories.NewGoogleAccountRepository(db)
	locationRepo := repositories.NewLocationRepository(db)

	// services
	campaignService := services.NewCampaignService(campaignRepo)
	campaignResultService := services.NewCampaignResultService(campaignResultRepo)
	googleAccountService := services.NewGoogleAccountService(googleAccountRepo, locationRepo)

	return &CompositionRoot{
		appConfig,
		db,
		campaignService,
		campaignResultService,
		googleAccountService,
		dataStoreService,
	}
}
