package server

import (
	"vpn_api/app/config"
	"vpn_api/app/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//APIServer type
type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	router *gin.Engine
}

//New create new APIServer instance
func New() *APIServer {
	return &APIServer{
		config: config.NewConfig(),
		logger: logrus.New(),
		router: gin.Default(),
	}
}

//Start APIServer with config
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	return s.router.Run(s.config.Server.BindAddr)
}

//GetRouter - configure router and return gin engine pointer
func (s *APIServer) GetRouter() *gin.Engine {
	s.configureRouter()
	return s.router
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.Server.LogLevel)
	if err != nil {
		s.logger.Fatal(err)
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	v1 := s.router.Group("/")
	{
		v1.POST("/pac", controllers.HandlePac())
		v1.POST("/ping", controllers.HandlePing())
	}
}
