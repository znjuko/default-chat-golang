package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	cookie "main/internal/cookie/repository"
	mware "main/internal/middleware"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	ch "main/internal/chats/delivery"
	eh "main/internal/emojies/delivery"
	sh "main/internal/socket/delivery"
	th "main/internal/socket_token/delivery"
	uh "main/internal/user/delivery"

	cr "main/internal/chats/repository"
	er "main/internal/emojies/repository"
	mr "main/internal/message/repository"
	sr "main/internal/socket/repository"
	tr "main/internal/socket_token/repository"
	ur "main/internal/user/repository"

	cu "main/internal/chats/usecase"
	eu "main/internal/emojies/usecase"
	su "main/internal/socket/usecase"
	tu "main/internal/socket_token/usecase"
	uu "main/internal/user/usecase"
)

type RequestHandlers struct {
	userHandler   uh.UserDeliveryRealisation
	chatHandler   ch.ChatsDelivery
	tokenHandler  th.TokenDelivery
	socketHandler sh.SocketDelivery
	emojiHandler  eh.ChatsDelivery
}

func InitializeHandlers(db *sql.DB, auth cookie.CookieRepositoryRealisation, logger *zap.SugaredLogger) RequestHandlers {

	redisPas := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")

	userR := ur.NewUserRepositoryRealisation(db)
	emojiR := er.NewEmojiUseRealisation(db)
	chatR := cr.NewChatRepoRealistaion(db)
	onlineR := sr.NewOnlineRepoRealosation(db)
	msgR := mr.NewMessageRepositoryRealisation("", "", db)
	tokenR := tr.NewTokenRepositoryRealisation(redisPort, redisPas)

	userU := uu.NewUserUseCaseRealisation(userR, auth)
	emojiU := eu.NewEmojiUseRealisation(emojiR)
	chatU := cu.NewChatUseCaseRealisation(chatR, emojiR)
	tokenU := tu.NewTokenUseCaseRealisation(tokenR)
	socketU := su.NewSocketUseCaseRealisation(msgR, tokenR, onlineR)

	userH := uh.NewUserDelivery(logger, userU)
	chatH := ch.NewChatsDelivery(logger, chatU)
	tokenH := th.NewTokenDelivery(logger, tokenU)
	socketH := sh.NewSocketDelivery(logger, socketU)
	emojiH := eh.NewChatsDelivery(logger, emojiU)

	return RequestHandlers{
		userHandler:   userH,
		chatHandler:   chatH,
		tokenHandler:  tokenH,
		socketHandler: socketH,
		emojiHandler:  emojiH,
	}
}

func InitializeDataBases(server *echo.Echo) (*sql.DB, cookie.CookieRepositoryRealisation) {
	err := godotenv.Load("project.env")
	if err != nil {
		server.Logger.Fatal("can't load .env file :", err.Error())
	}
	usernameDB := os.Getenv("POSTGRES_USERNAME")
	passwordDB := os.Getenv("POSTGRES_PASSWORD")
	nameDB := os.Getenv("POSTGRES_NAME")

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		server.Logger.Fatal("NO CONNECTION TO BD", err.Error())
	}

	redisPas := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")

	sessionDB := cookie.NewCookieRepositoryRealisation(redisPort, redisPas)

	return db, sessionDB
}

func main() {

	server := echo.New()
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	db, auth := InitializeDataBases(server)

	reqHandlers := InitializeHandlers(db, auth, logger)

	reqHandlers.chatHandler.InitHandlers(server)
	reqHandlers.socketHandler.InitHandlers(server)
	reqHandlers.tokenHandler.InitHandlers(server)
	reqHandlers.userHandler.InitHandlers(server)
	reqHandlers.emojiHandler.InitHandlers(server)

	midHandler := mware.NewMiddlewareHandler(logger, auth)
	midHandler.SetMiddleware(server)

	port := os.Getenv("PORT")

	server.Logger.Fatal(server.Start(port))

}
