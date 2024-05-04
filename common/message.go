package common

import (
	"github.com/CorrectRoadH/ZimaOS-Status/codegen/message_bus"
	"github.com/samber/lo"
)

// common properties
var (
	PropertyTypeMessage = message_bus.PropertyType{
		Name:        "message",
		Description: lo.ToPtr("message at different levels, typically for error"),
	}
)

// module properties
var (
	PropertyTypeModuleName = message_bus.PropertyType{
		Name:        "module:name",
		Description: lo.ToPtr("name of the module"),
		Example:     lo.ToPtr("zima-chat"),
	}

	PropertyTypeModuleTitle = message_bus.PropertyType{
		Name:        "module:title",
		Description: lo.ToPtr("title of the module"),
		Example:     lo.ToPtr("Zima Chat"),
	}
)

// event types for module
var (
	EventTypeModuleInstallBegin = message_bus.EventType{
		SourceID: ServiceName,
		Name:     "module:install-begin",
		PropertyTypeList: []message_bus.PropertyType{
			PropertyTypeModuleName,
			PropertyTypeModuleTitle,
		},
	}

	EventTypeModuleInstallEnd = message_bus.EventType{
		SourceID: ServiceName,
		Name:     "module:install-end",
		PropertyTypeList: []message_bus.PropertyType{
			PropertyTypeModuleName,
			PropertyTypeModuleTitle,
		},
	}

	EventTypeModuleInstallError = message_bus.EventType{
		SourceID: ServiceName,
		Name:     "module:install-error",
		PropertyTypeList: []message_bus.PropertyType{
			PropertyTypeModuleName,
			PropertyTypeModuleTitle,
			PropertyTypeMessage,
		},
	}
)

var EventTypes = []message_bus.EventType{
	EventTypeModuleInstallBegin, EventTypeModuleInstallEnd, EventTypeModuleInstallError,
}
