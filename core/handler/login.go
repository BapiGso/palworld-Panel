package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"palworld-panel/core/conf"
	"palworld-panel/core/mymiddleware"
	"time"
)

func LoginGet(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.Render(http.StatusOK, "login.template", nil)
	}
	if user.Valid { // 校验token
		return c.Redirect(302, "/")
	}
	return echo.ErrUnauthorized
}

func LoginPost(c echo.Context) error {
	password := conf.Config.Get("password")
	//计算提交表单的密码与盐 scrypt和数据库中密码是否一致
	if c.FormValue("password") == password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), //过期日期设置7天
		})
		t, err := token.SignedString(mymiddleware.JWTKey)
		if err != nil {
			return err
		}
		c.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    t,
			HttpOnly: true,
		})
		return c.Redirect(302, "/")
	}
	return echo.ErrUnauthorized
}
