package service

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

type MEMService struct {
}

func NewMEMService() *MEMService {
	return &MEMService{}

}

func (s *MEMService) Record() {
	value, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	MyService.DBService().InsertMemData(value.UsedPercent)
}
