package SystemService

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
	"www/framework/service/message"
)

func System()  {

	systemData := MessageService.ResponseMessage{}

	go func() {

		for {
			systemData.MessageType = "system_message"

			systemData.SystemMessage.Time = time.Now().Format("15:04:05")

			memory, _ := mem.VirtualMemory()

			systemData.SystemMessage.Memory = memory.UsedPercent

			cpuInfo, _ := cpu.Percent(time.Second, false)

			systemData.SystemMessage.Cpu = cpuInfo

			jsonString, _ := json.Marshal(systemData)

			MessageService.Send("", string(jsonString))

			time.Sleep(2000 * time.Millisecond)
		}
	}()
}