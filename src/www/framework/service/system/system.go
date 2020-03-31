package SystemService

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
	"www/framework/service/socket"
)

func System() {

	systemData := SocketService.SocketMessage{}

	go func() {

		for {

			systemData.MessageType = "system"

			systemData.SystemMessage.Time = time.Now().Format("15:04:05")

			memory, _ := mem.VirtualMemory()

			systemData.SystemMessage.Memory = memory.UsedPercent

			cpuInfo, _ := cpu.Percent(time.Second, false)

			systemData.SystemMessage.Cpu = cpuInfo

			SocketService.Channel.Channel = systemData

			time.Sleep(2000 * time.Millisecond)
		}

	}()
}