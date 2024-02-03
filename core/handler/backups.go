package handler

import (
	"github.com/labstack/echo/v4"
	"os"
	"palworld-panel/core/backups"
	"palworld-panel/core/conf"
	"path/filepath"
	"strconv"
)

func Backups(c echo.Context) error {
	configData := conf.Config.Clone()
	if configData["serverPath"] == nil {
		return c.String(200, "未找到服务器路径,请运行服务器后在打开面板")
	}
	if c.Param("frequency") != "" {
		frequency, _ := strconv.ParseInt(c.Param("frequency"), 10, 64)
		conf.Config.Set("backupsFrequency", frequency)
		return c.String(200, "设置成功")
	}
	if c.Param("backupNow") == "1" {
		backups.Compress()
		return c.Redirect(302, "/backups")
	}
	var zipFiles []string
	// 遍历文件夹
	if err := filepath.Walk(configData["serverPath"].(string)+"\\backups", func(path string, info os.FileInfo, err error) error {
		// 检查err
		if err != nil {
			return err
		}
		// 过滤zip文件
		if !info.IsDir() && filepath.Ext(path) == ".zip" {
			zipFiles = append(zipFiles, filepath.Base(path))
		}
		return nil
	}); err != nil {
		return err
	}
	return c.Render(200, "backups.template", map[string]any{
		"data":    configData,
		"backups": zipFiles,
	})
}
