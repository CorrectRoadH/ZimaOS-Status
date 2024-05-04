package service

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUService struct {
}

func NewCPUService() *CPUService {
	return &CPUService{}
}

func (s *CPUService) Record() {
	value, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	MyService.DBService().InsertCPUData(value[0])
}
