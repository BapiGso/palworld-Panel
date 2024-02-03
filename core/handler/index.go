package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/process"
	"palworld-panel/core/conf"
	"time"
)

func Index(c echo.Context) error {
	Status := struct {
		IsRunning    bool
		CreatTime    int64         //运行时间
		DurationHour time.Duration //运行小时数
		CPUPercent   float64       //cpu占用
		MemPercent   float32       //内存占用
	}{}
	processes, _ := process.Processes()
	for _, p := range processes {
		if name, _ := p.Name(); name == "PalServer-Win64-Test-Cmd.exe" {
			Status.IsRunning = true
			Status.CreatTime, _ = p.CreateTime()
			Status.CPUPercent, _ = p.CPUPercent()
			Status.MemPercent, _ = p.MemoryPercent()
			pastTime := time.Unix(Status.CreatTime/1000, 0)
			Status.DurationHour = time.Now().Sub(pastTime) / 24
			conf.Config.Set("status", Status)
		}
	}
	return c.Render(200, "admin.template", Status)
}
