package conf

import (
	"encoding/json"
	"os"
)

var Config C

type C struct{}

func (c *C) Get(key string) any {
	var tmp = make(map[string]any)
	if _, err := os.Stat("conf.json"); err == nil {
		// 文件存在,读取并反序列化
		file, _ := os.Open("conf.json")
		defer file.Close()
		json.NewDecoder(file).Decode(&tmp)
	}
	return tmp[key]
}

func (c *C) Set(key string, value any) {
	var tmp = make(map[string]any)
	if _, err := os.Stat("conf.json"); err == nil {
		// 文件存在,读取并反序列化
		file, _ := os.Open("conf.json")
		defer file.Close()
		json.NewDecoder(file).Decode(&tmp)
	}
	tmp[key] = value
	data, _ := json.MarshalIndent(tmp, "", "  ")
	os.WriteFile("conf.json", data, 0644)
}

// Clone 返回配置文件的副本
func (c *C) Clone() map[string]any {
	var tmp = make(map[string]any)
	if _, err := os.Stat("conf.json"); err == nil {
		// 文件存在,读取并反序列化
		file, _ := os.Open("conf.json")
		defer file.Close()
		json.NewDecoder(file).Decode(&tmp)
	}
	return tmp
}
