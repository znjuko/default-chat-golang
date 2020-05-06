package middleware

import (
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	sessions "main/internal/cookie"
	"main/internal/models"
	"net/http"
	"time"
)

type MiddlewareHandler struct {
	logger      *zap.SugaredLogger
	sessChecker sessions.CookieRepository
}

func NewMiddlewareHandler(logger *zap.SugaredLogger, checker sessions.CookieRepository) MiddlewareHandler {
	return MiddlewareHandler{logger: logger, sessChecker: checker}
}

func (mh MiddlewareHandler) SetMiddleware(server *echo.Echo) {
	server.Use(mh.SetCorsMiddleware)

	logFunc := mh.AccessLog()
	server.Use(mh.PanicMiddleWare)
	authFunc := mh.CheckAuthentication()

	server.Use(authFunc)
	server.Use(logFunc)

}

func (mh MiddlewareHandler) SetCorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, POST")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Login, Set-Cookie, Content-Type, Content-Length, Accept-Encoding, X-Csrf-Token, csrf-token, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Vary", "Cookie")

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusOK)
		}

		return next(c)

	}
}

func (mh MiddlewareHandler) PanicMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		defer func() error {
			if err := recover(); err != nil {
				rId := c.Get("REQUEST_ID").(string)
				mh.logger.Info(
					zap.String("ID", rId),
					zap.String("ERROR", err.(error).Error()),
					zap.Int("ANSWER STATUS", http.StatusInternalServerError),
				)

				return c.JSON(http.StatusInternalServerError, models.JsonStruct{Err: "server panic ! "})
			}
			return nil
		}()

		return next(c)
	}
}

func (mh MiddlewareHandler) AccessLog() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(rwContext echo.Context) error {

			uniqueID := uuid.NewV4()
			start := time.Now()
			rwContext.Set("REQUEST_ID", uniqueID.String())

			mh.logger.Info(
				zap.String("ID", uniqueID.String()),
				zap.String("URL", rwContext.Request().URL.Path),
				zap.String("METHOD", rwContext.Request().Method),
			)

			err := next(rwContext)

			respTime := time.Since(start)
			mh.logger.Info(
				zap.String("ID", uniqueID.String()),
				zap.Duration("TIME FOR ANSWER", respTime),
			)

			return err

		}
	}
}

func (mh MiddlewareHandler) CheckAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(rwContext echo.Context) error {

			cookie, err := rwContext.Cookie("session_id")

			userId := 0

			if err == nil {
				userId, err = mh.sessChecker.GetUserIdByCookie(cookie.Value)
			}

			if err != nil {
				userId = -1
				cookie = &http.Cookie{Expires: time.Now().AddDate(0, 0, -1)}
				rwContext.SetCookie(cookie)
			}

			rwContext.Set("user_id", int(userId))
			return next(rwContext)

		}
	}
}
