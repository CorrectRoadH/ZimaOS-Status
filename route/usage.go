package route

import (
	"net/http"

	"github.com/CorrectRoadH/ZimaOS-Status/codegen"
	"github.com/CorrectRoadH/ZimaOS-Status/service"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (r *StatusRoute) GetCPUUsage(ctx echo.Context, param codegen.GetCPUUsageParams) error {
	history, err := service.MyService.RecordService().GetCPUUsage(param.Start, param.End)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.BaseResponse{
			Message: lo.ToPtr(err.Error()),
		})
	}
	return ctx.JSON(http.StatusOK, codegen.BaseResponse{
		Data: history,
	})
}
