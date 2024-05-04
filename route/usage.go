package route

import (
	"net/http"

	"github.com/CorrectRoadH/ZimaOS-Status/codegen"
	"github.com/CorrectRoadH/ZimaOS-Status/service"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (r *StatusRoute) GetCPUUsageHistory(ctx echo.Context, param codegen.GetCPUUsageHistoryParams) error {
	history, err := service.MyService.RecordService().GetCPUUsage(param.Start, param.End)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.BaseResponse{
			Message: lo.ToPtr(err.Error()),
		})
	}
	return ctx.JSON(http.StatusOK, codegen.ResponseGetCpuInfoOk{
		Data: lo.ToPtr(history),
	})
}

func (r *StatusRoute) GetUsage(ctx echo.Context) error {
	cpu, err := service.MyService.DBService().LatestCPUUsage()
	mem, err2 := service.MyService.DBService().LatestMemUsage()
	if err != nil || err2 != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.BaseResponse{
			Message: lo.ToPtr(err.Error()),
		})
	}

	currenctPerformance := codegen.PerformanceUsage{
		Cpu:    &cpu,
		Memory: &mem,
	}
	return ctx.JSON(http.StatusOK, codegen.ResponseGetPerformanceUsageOk{
		Data: &currenctPerformance,
	})
}
