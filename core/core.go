package core

import (
	"embed"
	"github.com/labstack/echo/v4"
	"palworld-panel/assets"
)

type PalWorld struct {
	assetsFS *embed.FS  //主题所在文件夹
	e        *echo.Echo //后台框架
}

func New() (p *PalWorld) {
	p = &PalWorld{}
	p.assetsFS = &assets.Assets
	p.e = echo.New()
	return p
}
