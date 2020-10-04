package app

import (
	"github.com/basketforcode/http.server/app/handlers"
	"github.com/basketforcode/http.server/app/middleware"
	"github.com/basketforcode/http.server/app/services/cache"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/basketforcode/http.server/config"
	"github.com/gin-gonic/gin"
)

//App main structure
type App struct {
	config *config.Config
	router *gin.Engine
	store  *store.Store
	cache  *cache.Redis

	//handlers ...
	handlerInfo handlers.Handler
}

//New clear app
func New() App {
	conf := config.NewConfig()
	return App{
		config: conf,
		router: gin.Default(),
		store:  &store.Store{},
	}
}

//Start server
func (a *App) Start() error {
	if err := a.configureStore(); err != nil {
		return err
	}

	a.configureCache()

	a.configureHandlers()

	a.configureRouter()

	return a.router.Run(a.config.Server.BindAddr)
}

//Close all connections
func (a *App) Shutdown() error {
	err := a.store.Close()
	if err != nil {
		return err
	}

	err = a.cache.Close()
	if err != nil {
		return err
	}

	return nil
}

//set handler functions
func (a *App) configureHandlers() {
	a.handlerInfo = handlers.NewUserHandler(a.store, a.cache, a.config)
}

//bind router endpoints
func (a *App) configureRouter() {
	v1 := a.router.Group("/")
	{
		v1.Use(middleware.AuthMiddleware(a.store))
		v1.GET("/user/info", a.handlerInfo.Handle())
	}
}

//configure db store
func (a *App) configureStore() error {
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
func (a *App) configureCache() {
	r := cache.New(a.config)
	a.cache = &r
}
