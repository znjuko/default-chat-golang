package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/emojies"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
)

type ChatsDelivery struct {
	emojiLogic emojies.EmojiUse
	logger     *zap.SugaredLogger
}

func NewChatsDelivery(log *zap.SugaredLogger, emoji emojies.EmojiUse) ChatsDelivery {
	return ChatsDelivery{logger: log, emojiLogic: emoji}
}

func (CD ChatsDelivery) CreateEmoji(rwContext echo.Context) error {
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

	newEmoji := new(models.Emoji)

	err := rwContext.Bind(&newEmoji)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = CD.emojiLogic.CreateEmoji(*newEmoji)

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
		zap.Int("ANSWER STATUS", http.StatusCreated),
	)

	return rwContext.NoContent(http.StatusCreated)
}

func (CD ChatsDelivery) InitHandlers(server *echo.Echo) {
	server.POST("/emoji", CD.CreateEmoji)
}
