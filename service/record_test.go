package service_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/CorrectRoadH/ZimaOS-Status/service"
	"github.com/stretchr/testify/assert"
)

func TestGet15MinsRecord(t *testing.T) {
	service.Initialize("/tmp")
	rs := service.NewRecordService()

	startTimeStamp := time.Now().Unix()
	rs.StartRecord()
	time.Sleep(15 * time.Second)
	rs.StopRecord()
	endTimeStamp := time.Now().Unix()

	cpuHistory, err := rs.GetCPUUsageHistory(fmt.Sprint(startTimeStamp), fmt.Sprint(endTimeStamp))
	assert.Nil(t, err)
	assert.Equal(t, 5, len(cpuHistory))
}
