package server

import (
	"crypto-api/server/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

type RequestHandlerFunction func(c *gin.Context)

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/users/:id", a.handleRequest(services.GetUser))
	a.Post("/users", a.handleRequest(services.CreateUser))
	a.Put("/users/:id", a.handleRequest(services.UpdateUser))

	// Cryptocurrrency related routing
	a.Get("/currency", a.handleRequest(services.GetAllCurrencies))
	a.Get("/currency/:cid", a.handleRequest(services.GetCurrency))
	a.Get("/currency/:cid/timeline", a.handleRequest(services.GetCurrencyTimeline))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(c *gin.Context)) {
	a.Router.GET(path, f)
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(c *gin.Context)) {
	a.Router.POST(path, f)
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(c *gin.Context)) {
	a.Router.PUT(path, f)
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(c *gin.Context)) {
	a.Router.DELETE(path, f)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) handleRequest(handler RequestHandlerFunction) RequestHandlerFunction {
	return func(c *gin.Context) {
		handler(c)
	}
}
