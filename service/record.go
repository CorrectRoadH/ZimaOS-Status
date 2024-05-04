package service

import (
	"fmt"

	"github.com/CorrectRoadH/ZimaOS-Status/codegen"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type RecordImpl interface {
	Record()
}

type RecordService struct {
	imps           []RecordImpl
	recordInterval int
	crontab        *cron.Cron
}

func NewRecordService() *RecordService {
	return &RecordService{
		recordInterval: 3,
		imps: []RecordImpl{
			NewCPUService(),
			NewMEMService(),
		},
		crontab: cron.New(cron.WithSeconds()),
	}
}

func (s *RecordService) Record() {
	for _, imp := range s.imps {
		imp.Record()
	}
}

func (s *RecordService) StartRecord() {
	if _, err := s.crontab.AddFunc(fmt.Sprintf("@every %ds", s.recordInterval), func() {
		s.Record()
	}); err != nil {
		logger.Error("error when trying to add cron job", zap.Error(err))
	}
	s.crontab.Start()
}

func (s *RecordService) StopRecord() {
	s.crontab.Stop()
}

func (s *RecordService) GetCPUUsage(start string, end string) ([]codegen.CpuInfo, error) {
	return MyService.DBService().QueryCPUUsageHistory(start, end)
}
