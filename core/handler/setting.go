package handler

import "github.com/labstack/echo/v4"

func Setting(c echo.Context) error {
	return c.JSON(200, "开发中")
}

func SettingPost(c echo.Context) error {
	//如果没有配置文件，把配置文件复制一个过去

	//然后根据form的值更新配置文件
	return nil
}
