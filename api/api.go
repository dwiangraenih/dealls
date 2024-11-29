package api

import (
	"fmt"
	"github.com/dwiangraeni/dealls/handler"
	"github.com/dwiangraeni/dealls/infra"
	"github.com/dwiangraeni/dealls/manager"
	"github.com/dwiangraeni/dealls/middleware"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

// Server API server interface
type Server interface {
	Run()
}

type server struct {
	router         chi.Router
	infra          infra.Infra
	serviceManager manager.ServiceManager
}

// NewServer construct new API server
func NewServer(infra infra.Infra) Server {
	return &server{
		router:         chi.NewRouter(),
		infra:          infra,
		serviceManager: manager.NewServiceManager(infra),
	}
}

func (c *server) Run() {
	c.endpoint()
	c.run()
}

func (c *server) endpoint() {
	authHandler := handler.NewAuthHandler(c.serviceManager.AuthService())
	accountHandler := handler.NewAccountHandler(c.serviceManager.AccountService())
	token := middleware.NewTokenValidator(c.serviceManager.AccountManager())
	userSwipeLogHandler := handler.NewUserSwipeLogHandler(c.serviceManager.UserSwipeLogService())

	c.router.Route("/dealls", func(r chi.Router) {
		// auth
		r.Route("/auth", func(an chi.Router) {
			an.With().Post("/login", authHandler.HandlerLogin)
			an.With().Post("/register", authHandler.HandlerRegister)
		})

		// account
		r.Route("/account", func(an chi.Router) {
			an.With(token.RequireAccountToken()).Group(func(an chi.Router) {
				an.With(token.RequireAccountType("FREE")).Post("/upgrade", accountHandler.UpgradeAccount)
			})
			an.With(token.RequireAccountToken()).Get("/list", accountHandler.GetListAccountNewMatchPagination)

		})

		// swipe
		r.Route("/swipe", func(an chi.Router) {
			an.With(token.RequireAccountToken()).Post("/interaction", userSwipeLogHandler.ProcessUserSwipe)
		})
	})

}

func (c *server) run() {
	apiConfig := c.infra.Config().Sub("api")
	addr := fmt.Sprintf("%s:%d", apiConfig.GetString("host"), apiConfig.GetInt("port"))

	svr := &http.Server{
		Addr:         addr,
		Handler:      c.router,
		ReadTimeout:  time.Duration(apiConfig.GetInt("read_timeout")) * time.Second,
		WriteTimeout: time.Duration(apiConfig.GetInt("write_timeout")) * time.Second,
		IdleTimeout:  time.Duration(apiConfig.GetInt("idle_timeout")) * time.Second,
	}

	if err := svr.ListenAndServe(); err != nil {
		fmt.Errorf("error Server start error with err : %v", err.Error())
	}

	log.Printf("Server started at %s", addr)
}
