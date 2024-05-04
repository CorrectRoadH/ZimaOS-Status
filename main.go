//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -generate types,server,spec -package codegen api/status/openapi.yaml > codegen/status_api.go"
//go:generate bash -c "mkdir -p codegen/message_bus && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -generate types,client -package message_bus https://raw.githubusercontent.com/IceWhaleTech/CasaOS-MessageBus/main/api/message_bus/openapi.yaml > codegen/message_bus/api.go"

package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/CorrectRoadH/ZimaOS-Status/common"
	"github.com/CorrectRoadH/ZimaOS-Status/config"
	"github.com/CorrectRoadH/ZimaOS-Status/route"
	"github.com/CorrectRoadH/ZimaOS-Status/service"
	"github.com/IceWhaleTech/CasaOS-Common/model"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/coreos/go-systemd/daemon"
	"go.uber.org/zap"

	_ "embed"

	util_http "github.com/IceWhaleTech/CasaOS-Common/utils/http"
)

var (
	commit = "private build"
	date   = "private build"

	//go:embed api/index.html
	_docHTML string

	//go:embed api/status/openapi.yaml
	_docYAML string

	//go:embed build/sysroot/etc/casaos/mod-management.conf.sample
	_confSample string
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse arguments and intialize
	{
		// configFlag := flag.String("c", "", "config file path")
		versionFlag := flag.Bool("v", false, "version")

		flag.Parse()

		if *versionFlag {
			fmt.Printf("version: %s\n", common.Version)
			fmt.Printf("git commit: %s\n", commit)
			fmt.Printf("build date: %s\n", date)

			os.Exit(0)
		}

		fmt.Printf("git commit: %s\n", commit)
		fmt.Printf("build date: %s\n", date)

		// config.InitSetup(*configFlag, _confSample)

		logger.LogInit(config.AppInfo.LogPath, config.AppInfo.LogSaveName, config.AppInfo.LogFileExt)

		service.Initialize(config.CommonInfo.RuntimePath)
	}

	{
		// cron job
		service.MyService.RecordService().StartRecord()
		defer service.MyService.RecordService().StopRecord()
	}

	// register at message bus
	{
		response, err := service.MyService.MessageBus().RegisterEventTypesWithResponse(ctx, common.EventTypes)
		if err != nil {
			logger.Error("error when trying to register one or more event types - some event type will not be discoverable", zap.Error(err))
		}

		if response != nil && response.StatusCode() != http.StatusOK {
			logger.Error("error when trying to register one or more event types - some event type will not be discoverable", zap.String("status", response.Status()), zap.String("body", string(response.Body)))
		}
	}

	// setup listener
	listener, _err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", "0"))
	if _err != nil { // use _err to avoid shadowing the err variables below.
		panic(_err)
	}

	// initialize routers and register at gateway
	{
		apiPaths := []string{
			route.APIPath,
			route.DocPath,
		}

		for _, apiPath := range apiPaths {
			if err := service.MyService.Gateway().CreateRoute(&model.Route{
				Path:   apiPath,
				Target: "http://" + listener.Addr().String(),
			}); err != nil {
				panic(err)
			}
		}
	}

	router := route.GetRouter()
	docRouter := route.GetDocRouter(_docHTML, _docYAML)

	mux := &util_http.HandlerMultiplexer{
		HandlerMap: map[string]http.Handler{
			"v2":  router,
			"doc": docRouter,
		},
	}

	// notify systemd that we are ready
	{
		if supported, err := daemon.SdNotify(false, daemon.SdNotifyReady); err != nil {
			logger.Error("Failed to notify systemd that casaos main service is ready", zap.Any("error", err))
		} else if supported {
			logger.Info("Notified systemd that casaos main service is ready")
		} else {
			logger.Info("This process is not running as a systemd service.")
		}

		logger.Info("Virtualization management service is listening...", zap.Any("address", listener.Addr().String()))
	}

	s := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second, // fix G112: Potential slowloris attack (see https://github.com/securego/gosec)
	}

	_err = s.Serve(listener) // not using http.serve() to fix G114: Use of net/http serve function that has no support for setting timeouts (see https://github.com/securego/gosec)
	if _err != nil {
		panic(_err)
	}
}
