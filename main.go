package main

import (
	"palworld-panel/core"
)

func main() {

	//for {
	//	conn, err := rcon.Dial("127.0.0.1:25575", "")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	percent, _ := cpu.Percent(time.Second, false)
	//	fmt.Println(percent[0]) // percent[0]
	//	memInfo, _ := mem.VirtualMemory()
	//	//fmt.Println(fmt.Sprintf("broadcast mem占用:%v,cpu占用:%v", memInfo.UsedPercent, percent[0]))
	//	fmt.Sprintf("broadcast mem %v,cpu:%v", memInfo, percent)
	//	response, err := conn.Execute(fmt.Sprintf("broadcast mem:%v,cpu:%v", memInfo.UsedPercent, percent[0]))
	//	if err != nil {
	//	}
	//	fmt.Println(response)
	//	conn.Close()
	//	time.Sleep(time.Minute)
	//}

	p := core.New()
	p.LoadMiddlewareRoutes()
	p.Listen(":8080")
}
