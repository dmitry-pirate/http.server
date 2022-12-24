package app

import (
	"github.com/basketforcode/http.server/internal/middleware"
	"github.com/basketforcode/http.server/internal/user"
	"github.com/basketforcode/http.server/pkg/cache"
	"github.com/basketforcode/http.server/pkg/config"
	"github.com/basketforcode/http.server/pkg/store"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	config *config.Config
	router *gin.Engine
	store  *store.Store
	cache  *cache.Redis

	//handlers ...
	handlerInfo Handler
}

func New() App {
	conf := config.NewConfig()
	return App{
		config: conf,
		router: gin.Default(),
		store:  &store.Store{},
	}
}

func (a *App) Start() error {
	if err := a.configureStore(); err != nil {
		return err
	}

	a.configureCache()

	a.configureHandlers()

	a.configureRouter()

	return a.router.Run(a.config.Server.BindAddr)
}

func (a *App) Shutdown() error {
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

func (a *App) configureHandlers() {
	a.handlerInfo = user.NewHandler(a.store, a.cache, a.config)
}

func (a *App) configureRouter() {
	v1 := a.router.Group("/")
	{
		v1.Use(middleware.Auth(a.store))
		v1.GET("/users", a.handlerInfo.Handle())
	}
}

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

func (a *App) configureCache() {
	r := cache.New(a.config)
	a.cache = &r
}
