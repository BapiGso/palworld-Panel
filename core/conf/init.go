package conf

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
	"path/filepath"
)

func init() {
	if _, err := os.Stat("conf.json"); err == nil {
		fmt.Println("检测到配置文件，读取中")
		configData := Config.Clone()
		if _, ok := configData["password"]; !ok {
			fmt.Println("配置文件中缺少password字段")
			os.Exit(1)
		}
	} else {
		fmt.Println("检测到您是第一次使用面板,查找帕鲁服务器进程中")
		processes, _ := process.Processes()
		var name string
		for _, p := range processes {
			if name, _ = p.Name(); name == "PalServer.exe" {
				exe, _ := p.Exe()
				dir, _ := filepath.Split(exe)
				if _, err := os.Stat(filepath.Join(dir, "backups")); os.IsNotExist(err) {
					os.Mkdir(filepath.Join(dir, "backups"), 0755)
				}
				if _, err := os.Stat(dir + "\\Pal\\Binaries\\Win64\\steam_appid.txt"); err == nil {
					Config.Set("installedSteam", true)
				}
				Config.Set("serverPath", dir)
				fmt.Println("默认管理密码是123456，您可以在配置文件或面板中修改")
				Config.Set("password", "123456")
				break
			}
		}
		if name != "PalServer.exe" {
			fmt.Println("未找到帕鲁服务器进程，请运行帕鲁服务器后在打开本面板")
			os.Exit(1)
		}
	}
}
