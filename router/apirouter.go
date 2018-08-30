package router

import (
	"fmt"
	"io"
	"os"

	gin "github.com/gin-gonic/gin"
	"github.com/onesafe/license_manager/log"
)

var _API_ROUTER *APIRouter

type APIRouter struct {
	mainRouter  *gin.Engine
	initialized bool
}

func (router *APIRouter) Run() {
	router.mainRouter.Run()
}

func GetAPIRouter() *APIRouter {
	if _API_ROUTER.initialized {
		return _API_ROUTER
	}
	_init()
	return _API_ROUTER
}

func init() {
	_init()
}

func _init() {
	f, err := os.Create("log/license_manager.log")
	if err != nil && os.IsNotExist(err) {
		os.Mkdir("log", os.FileMode(0777))
		f, _ = os.Create("log/license_manager.log")
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	_API_ROUTER = &APIRouter{
		mainRouter:  gin.New(),
		initialized: false,
	}

	_API_ROUTER.mainRouter.Use(gin.Recovery())
	_API_ROUTER.mainRouter.Use(log.GinLogger(gin.DefaultWriter))
	_API_ROUTER.Register("GET", "/ping", _ping)
	_API_ROUTER.initialized = true
}

func _ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
		"data":    "test",
	})
}

func (r *APIRouter) Register(method, path string, f gin.HandlerFunc) error {
	var err error
	err = nil

	handlers := gin.HandlersChain{}
	handlers = append(handlers, f)

	switch method {
	case "GET":
		r.mainRouter.GET(path, handlers...)
	case "POST":
		r.mainRouter.POST(path, handlers...)
	case "PUT":
		r.mainRouter.PUT(path, handlers...)
	case "DELETE":
		r.mainRouter.DELETE(path, handlers...)
	default:
		err = fmt.Errorf("Invalid Method to register")
	}
	return err
}
