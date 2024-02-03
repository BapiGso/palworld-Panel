package rcon

import (
	"fmt"
	"github.com/gorcon/rcon"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func init() {
	go func() {
		ticker := time.NewTicker(time.Duration(5) * time.Minute)
		for {
			<-ticker.C
			conn, err := rcon.Dial("127.0.0.1:25575", "")
			if err != nil {
				fmt.Println(err)
			}
			percent, _ := cpu.Percent(time.Second, false)
			fmt.Println(percent[0]) // percent[0]
			memInfo, _ := mem.VirtualMemory()
			_, err = conn.Execute(fmt.Sprintf("broadcast mem:%.2f%%,cpu:%.2f%%", memInfo.UsedPercent, percent[0]))

			_, err = conn.Execute("ShowPlayers")
			if err != nil {
				fmt.Println(err)
			}
			conn.Close()
		}
	}()
}
