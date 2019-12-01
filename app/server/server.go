package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//APIServer type
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
}

//New create new APIServer instance
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
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
	return s.router.Run(s.config.server.BindAddr)
}

//GetRouter - configure router and return gin engine pointer
func (s *APIServer) GetRouter() *gin.Engine {
	s.configureRouter()
	return s.router
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.server.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	v1 := s.router.Group("/")
	{
		v1.POST("/pac", s.handlePac())
		v1.POST("/ping", s.handlePing())
	}
}

func (s *APIServer) handlePac() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mode": "vpn"})
	}
}

func (s *APIServer) handlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mode": "vpn"})
	}
}
