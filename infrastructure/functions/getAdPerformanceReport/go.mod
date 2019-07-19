module serverless-functions-go/infrastructure/functions/getAdPerformanceReport

require (
	cloud.google.com/go v0.41.0 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	serverless-functions-go/application v0.0.0
	serverless-functions-go/domain v0.0.0
	serverless-functions-go/infrastructure v0.0.0 // indirect
	serverless-functions-go/interfaces v0.0.0 // indirect
)

replace (
	serverless-functions-go/application => ../../../application
	serverless-functions-go/domain => ../../../domain
	serverless-functions-go/infrastructure => ../../
	serverless-functions-go/interfaces => ../../../interfaces
)
