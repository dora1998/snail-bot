package http

import (
	"github.com/dora1998/snail-bot/commands"
	"github.com/dora1998/snail-bot/db"
	"github.com/dora1998/snail-bot/repository"
	"github.com/dora1998/snail-bot/twitter"
	"github.com/dora1998/snail-bot/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

type Server struct {
	dbInstance     *sqlx.DB
	repository     *repository.DBRepository
	twitterClient  *twitter.TwitterClient
	commandHandler *commands.CommandHandler
}

func NewServer() *Server {
	dbConfig, err := utils.ReadDBConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbInstance, err := db.NewDBInstance(dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.RunMigration(dbInstance)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := repository.NewDBRepository(dbInstance)
	twitterClient := twitter.NewTwitterClient()
	handler := commands.NewCommandHandler(repo, twitterClient)

	return &Server{dbInstance: dbInstance, repository: repo, twitterClient: twitterClient, commandHandler: handler}
}

func (s *Server) Routes() *gin.Engine {
	router := gin.Default()
	router.POST("/webhook/ifttt", s.PostIFTTTWebHook)
	router.GET("/webhook/twitter", s.GetCRCToken)
	router.POST("/webhook/twitter", s.PostWebHook)
	return router
}

func (s *Server) Start() error {
	router := s.Routes()
	defer s.dbInstance.Close()
	return router.Run(":8080")
}
