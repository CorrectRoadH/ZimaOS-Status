package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CorrectRoadH/ZimaOS-Status/codegen/message_bus"
	"github.com/CorrectRoadH/ZimaOS-Status/common"
	"github.com/CorrectRoadH/ZimaOS-Status/config"
	"github.com/IceWhaleTech/CasaOS-Common/external"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"

	"go.uber.org/zap"
)

var MyService *Services

type Services struct {
	gateway     external.ManagementService
	record      *RecordService
	runtimePath string
	db          *DBService
}

func Initialize(runtimePath string) {
	MyService = &Services{
		runtimePath: runtimePath,
	}
}

func (s *Services) DBService() *DBService {
	if s.db == nil {
		s.db = NewDBService()
	}
	return s.db
}

func (s *Services) RecordService() *RecordService {
	if s.record == nil {
		s.record = NewRecordService()
	}
	return s.record
}

func (s *Services) Gateway() external.ManagementService {
	if s.gateway == nil {
		gateway, err := external.NewManagementService(s.runtimePath)
		if err != nil && len(s.runtimePath) > 0 {
			panic(err)
		}

		s.gateway = gateway
	}

	return s.gateway
}

func (s *Services) MessageBus() *message_bus.ClientWithResponses {
	client, _ := message_bus.NewClientWithResponses("", func(c *message_bus.Client) error {
		// error will never be returned, as we always want to return a client, even with wrong address,
		// in order to avoid panic.
		//
		// If we don't avoid panic, message bus becomes a hard dependency, which is not what we want.

		messageBusAddress, err := external.GetMessageBusAddress(config.CommonInfo.RuntimePath)
		if err != nil {
			c.Server = "message bus address not found"
			return nil
		}

		c.Server = messageBusAddress
		return nil
	})

	return client
}

func PublishEventWrapper(ctx context.Context, eventType message_bus.EventType, properties map[string]string) {
	if MyService == nil {
		fmt.Println("failed to publish event - messsage bus service not initialized")
		return
	}

	if properties == nil {
		properties = map[string]string{}
	}

	// merge with properties from context
	for k, v := range common.PropertiesFromContext(ctx) {
		properties[k] = v
	}

	response, err := MyService.MessageBus().PublishEventWithResponse(ctx, common.ServiceName, eventType.Name, properties)
	if err != nil {
		logger.Error("failed to publish event", zap.Error(err))
		return
	}
	defer response.HTTPResponse.Body.Close()

	if response.StatusCode() != http.StatusOK {
		logger.Error("failed to publish event", zap.String("status code", response.Status()))
	}
}
