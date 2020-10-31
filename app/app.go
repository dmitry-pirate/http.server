package app

import (
	"github.com/basketforcode/http.server/app/handlers"
	"github.com/basketforcode/http.server/app/middleware"
	"github.com/basketforcode/http.server/app/services/cache"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/basketforcode/http.server/config"
	"github.com/gin-gonic/gin"
	"log"
)

//App main structure
type app struct {
	config *config.Config
	router *gin.Engine
	store  *store.Store
	cache  *cache.Redis

	//handlers ...
	handlerInfo handlers.Handler
}

//New clear app
func New() *app {
	conf := config.NewConfig()
	return &app{
		config: conf,
		router: gin.Default(),
		store:  &store.Store{},
	}
}

//Start server
func (a *app) Start() error {
	if err := a.configureStore(); err != nil {
		return err
	}

	a.configureCache()

	a.configureRouter()

	return a.router.Run(a.config.Server.BindAddr)
}

//Close all connections
func (a *app) Shutdown() error {
	err := a.store.Close()
	if err != nil {
		return err
	}

	log.Println("Store connection closed...")

	err = a.cache.Close()
	if err != nil {
		return err
	}

	log.Println("Cache connection closed...")

	return nil
}

//bind router endpoints
func (a *app) configureRouter() {
	v1 := a.router.Group("/")
	{
		v1.Use(middleware.InjectMiddleware(a.store, a.cache))
		v1.Use(middleware.AuthMiddleware())
		v1.GET("/user/info", handlers.NewUserHandler().Handle())
	}
}

//configure db store
func (a *app) configureStore() error {
	st, err := store.New(a.config)
	if err != nil {
		return err
	}

	if err := st.MasterConnection().Ping(); err != nil {
		return err
	}

	a.store = st
	return nil
}

//connect to cache driver
func (a *app) configureCache() {
	r := cache.New(a.config)
	a.cache = &r
}
