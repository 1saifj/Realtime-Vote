package middleware

import "github.com/labstack/echo"

type UserMiddleware struct {
}

func (m *UserMiddleware) GetUserData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//get ip address
		ctx.Set("ip", ctx.Request().RemoteAddr)
		println(ctx.Request().RemoteAddr)
		return next(ctx)
	}
}
