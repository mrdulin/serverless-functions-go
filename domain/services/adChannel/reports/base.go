package services

import models "serverless-functions-go/domain/models/adChannel"

type IBaseReportService interface {
	GetRows(report models.BaseReport) []interface{}
}

type BaseReportService struct {
}

func NewBaseReportService() IBaseReportService {
	return &BaseReportService{}
}

func (svc *BaseReportService) GetRows(report models.BaseReport) []interface{} {
	return report.Table.Row
}
