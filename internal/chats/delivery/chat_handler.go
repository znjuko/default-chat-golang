package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/chats"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
)

type ChatsDelivery struct {
	chatsLogic chats.ChatUseCase
	logger     *zap.SugaredLogger
}

func NewChatsDelivery(log *zap.SugaredLogger, chatsRealisation chats.ChatUseCase) ChatsDelivery {
	return ChatsDelivery{logger: log, chatsLogic: chatsRealisation}
}

func (CD ChatsDelivery) CreateChat(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	newChat := new(models.NewChatUsers)

	err := rwContext.Bind(&newChat)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = CD.chatsLogic.CreateChat(*newChat, userId)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	CD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}


func (CD ChatsDelivery) GetChatMessages(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	chatId , _  := strconv.Atoi(rwContext.Param("id"))

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	chatInfo, err := CD.chatsLogic.GetChat(chatId)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	CD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, chatInfo)

}

func (CD ChatsDelivery) GetAllChats(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	allData, err := CD.chatsLogic.GetChatsAndOnlineUsers(userId)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	CD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, allData)
}

func (CD ChatsDelivery) InitHandlers(server *echo.Echo) {
	server.POST("/chats", CD.CreateChat)
	server.GET("/chats/:id", CD.GetChatMessages)
	server.GET("/chats", CD.GetAllChats)
}