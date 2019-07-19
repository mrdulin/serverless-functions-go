module serverless-functions-go/infrastructure/functions/getAdPerformanceReport

require (
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
