package datastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type EnvVarEntity struct {
	AD_CHANNEL_API_BASE_URL      string
	APPLICATION_URL              string
	DAILY_REPORT_FROM            string
	DAILY_REPORT_TO              string
	DEVELOPMENT_BUILD            string
	FB_X_API_KEY                 string
	INSTANCE_CONNECTION_NAME     string
	SENDGRID_API_KEY             string
	SENDGRID_CC                  string
	SENDGRID_FROM                string
	SENDGRID_TO                  string
	SQL_DATABASE                 string
	SQL_PASSWORD                 string
	SQL_USER                     string
	STORAGE_BUCKET_NAME          string
	SQL_INSTANCE_CONNECTION_NAME string
	STORAGE_PRIVATE_KEY          string
	STORAGE_CLIENT_EMAIL         string
}

type IService interface {
	GetEnvVars() (*EnvVarEntity, error)
}

type Service struct {
	options *Options
	client  *datastore.Client
	ctx     context.Context
}

type Options struct {
	ProjectID       string
	CredentialsFile string
}

func New(options *Options) (*Service, error) {
	ctx := context.Background()
	clientOption := option.ClientOption(
		option.WithCredentialsFile(options.CredentialsFile),
	)
	client, err := datastore.NewClient(ctx, options.ProjectID, clientOption)
	if err != nil {
		return nil, errors.Wrap(err, "datastore.NewClient")
	}
	return &Service{options, client, ctx}, nil
}

func (svc *Service) GetEnvVars() (*EnvVarEntity, error) {
	var envVarEntities []EnvVarEntity
	kind := "envVars"
	q := datastore.NewQuery(kind)
	if _, err := svc.client.GetAll(svc.ctx, q, &envVarEntities); err != nil {
		return nil, errors.Wrap(err, "svc.client.GetAll(svc.ctx, q, &entities)")
	}
	envVarEntity := envVarEntities[0]
	return &envVarEntity, nil
}
