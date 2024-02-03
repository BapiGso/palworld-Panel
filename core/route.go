package core

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"palworld-panel/core/conf"
	"palworld-panel/core/handler"
	"palworld-panel/core/mymiddleware"
	"text/template"
)

func (p *PalWorld) LoadMiddlewareRoutes() {
	p.e.Validator = &mymiddleware.Validator{}
	p.e.Renderer = &mymiddleware.TemplateRender{
		Template: template.Must(
			template.ParseFS(
				p.assetsFS,
				"*.template",
			),
		).Funcs(template.FuncMap{}),
	}
	p.e.Use(middleware.Recover())
	p.e.Use(middleware.Gzip())
	p.e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/login"
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.Redirect(302, "/login")
		},
		SigningKey:  mymiddleware.JWTKey,
		TokenLookup: "cookie:token",
	}))

	p.e.GET("/login", handler.LoginGet)
	p.e.POST("/login", handler.LoginPost)
	p.e.GET("/", handler.Index)
	p.e.GET("/backups", handler.Backups)
	p.e.GET("/setting", handler.Setting)
	p.e.StaticFS("/assets", p.assetsFS)
	p.e.Static("/backups", conf.Config.Get("serverPath").(string))
}
