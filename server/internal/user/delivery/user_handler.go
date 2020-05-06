package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/models"
	users "main/internal/user"
	"net/http"
	"time"
)

type UserDeliveryRealisation struct {
	userLogic users.UserUseCase
	logger    *zap.SugaredLogger
}



func (userD UserDeliveryRealisation) Login(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userAuthData := new(models.Session)

	err := rwContext.Bind(&userAuthData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	exprTime := 12 * time.Hour

	cookieValue, err := userD.userLogic.Login(*userAuthData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)


		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(exprTime),
	}
	rwContext.SetCookie(&cookie)

	return rwContext.NoContent(http.StatusOK)
}

func (userD UserDeliveryRealisation) Logout(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return rwContext.NoContent(http.StatusOK)
	}

	err = userD.userLogic.Logout(cookie.Value)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	rwContext.SetCookie(cookie)

	return rwContext.NoContent(http.StatusOK)
}

func (userD UserDeliveryRealisation) Register(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userRegData := new(models.Session)

	err := rwContext.Bind(&userRegData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	exprTime := 12 * time.Hour

	cookieValue, err := userD.userLogic.Register(*userRegData)

	if err != nil {
		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(exprTime),
	}
	rwContext.SetCookie(&cookie)

	return rwContext.NoContent(http.StatusOK)
}


func NewUserDelivery(log *zap.SugaredLogger, userRealisation users.UserUseCase) UserDeliveryRealisation {
	return UserDeliveryRealisation{userLogic: userRealisation, logger: log}
}

func (userD UserDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.POST("/login", userD.Login)
	server.POST("/registration", userD.Register)

	server.DELETE("/login", userD.Logout)
}