package repositories

import (
	"serverless-functions-go/domain/models/cedar"
)

type LocationRepository interface {
	FindLocationsBoundGoogleClientCustomerId() ([]cedar.Location, error)
}
