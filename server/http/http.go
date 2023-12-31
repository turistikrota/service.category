package http

import (
	"strings"
	"time"

	httpServer "github.com/turistikrota/service.shared/server/http"

	"github.com/cilloparch/cillop/helpers/http"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/turistikrota/service.category/app"
	"github.com/turistikrota/service.category/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"
)

type srv struct {
	config      config.App
	app         app.Application
	i18n        i18np.I18n
	validator   validation.Validator
	tknSrv      token.Service
	sessionSrv  session.Service
	httpHeaders config.HttpHeaders
}

type Config struct {
	Env         config.App
	App         app.Application
	I18n        i18np.I18n
	Validator   validation.Validator
	HttpHeaders config.HttpHeaders
	TokenSrv    token.Service
	SessionSrv  session.Service
}

func New(config Config) server.Server {
	return srv{
		config:      config.Env,
		app:         config.App,
		i18n:        config.I18n,
		validator:   config.Validator,
		tknSrv:      config.TokenSrv,
		sessionSrv:  config.SessionSrv,
		httpHeaders: config.HttpHeaders,
	}
}

func (h srv) Listen() error {
	http.RunServer(http.Config{
		Host:        h.config.Http.Host,
		Port:        h.config.Http.Port,
		I18n:        &h.i18n,
		AcceptLangs: []string{},
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Use(h.cors(), h.deviceUUID(), h.rateLimit())
			admin := router.Group("/admin", h.currentUserAccess(), h.requiredAccess())
			admin.Post("/", h.adminRoute(config.Roles.Category.Create), h.wrapWithTimeout(h.CategoryCreate))
			admin.Get("/", h.adminRoute(config.Roles.Category.List), h.wrapWithTimeout(h.CategoryAdminList))
			admin.Get("/parents", h.adminRoute(config.Roles.Category.List), h.wrapWithTimeout(h.CategoryAdminListParents))
			admin.Get("/:uuid/child", h.adminRoute(config.Roles.Category.ListChildren), h.wrapWithTimeout(h.CategoryAdminListChild))
			admin.Put("/:uuid/enable", h.adminRoute(config.Roles.Category.Enable), h.wrapWithTimeout(h.CategoryEnable))
			admin.Put("/:uuid/re-order", h.adminRoute(config.Roles.Category.ReOrder), h.wrapWithTimeout(h.CategoryUpdateOrder))
			admin.Put("/:uuid/disable", h.adminRoute(config.Roles.Category.Disable), h.wrapWithTimeout(h.CategoryDisable))
			admin.Put("/:uuid", h.adminRoute(config.Roles.Category.Update), h.wrapWithTimeout(h.CategoryUpdate))
			admin.Get("/:uuid", h.adminRoute(config.Roles.Category.ViewAdmin), h.wrapWithTimeout(h.CategoryAdminView))
			admin.Delete("/:uuid", h.adminRoute(config.Roles.Category.Delete), h.wrapWithTimeout(h.CategoryDelete))
			router.Get("/", h.wrapWithTimeout(h.CategoryList))
			router.Get("/fields", h.wrapWithTimeout(h.CategoryFindFields))
			router.Get("/:slug", h.wrapWithTimeout(h.CategoryView))
			router.Get("/:uuid/child", h.wrapWithTimeout(h.CategoryListChild))
			return router
		},
	})
	return nil
}

func (h srv) parseBody(c *fiber.Ctx, d interface{}) {
	http.ParseBody(c, h.validator, h.i18n, d)
}

func (h srv) parseParams(c *fiber.Ctx, d interface{}) {
	http.ParseParams(c, h.validator, h.i18n, d)
}

func (h srv) parseQuery(c *fiber.Ctx, d interface{}) {
	http.ParseQuery(c, h.validator, h.i18n, d)
}

func (h srv) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       &h.i18n,
		MsgKey:     Messages.Error.CurrentUserAccess,
		HeaderKey:  httpServer.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h srv) rateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
	})
}

func (h srv) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.httpHeaders.Domain,
	})
}

func (h srv) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   h.i18n,
		MsgKey: Messages.Error.RequiredAuth,
	})
}

func (h srv) adminRoute(extra ...string) fiber.Handler {
	claims := []string{config.Roles.Admin}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return claim_guard.New(claim_guard.Config{
		Claims: claims,
		I18n:   h.i18n,
		MsgKey: Messages.Error.AdminRoute,
	})
}

func (h srv) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     h.httpHeaders.AllowedOrigins,
		AllowMethods:     h.httpHeaders.AllowedMethods,
		AllowHeaders:     h.httpHeaders.AllowedHeaders,
		AllowCredentials: h.httpHeaders.AllowCredentials,
		AllowOriginsFunc: func(origin string) bool {
			origins := strings.Split(h.httpHeaders.AllowedOrigins, ",")
			for _, o := range origins {
				if strings.Contains(origin, o) {
					return true
				}
			}
			return false
		},
	})
}

func (h srv) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}
