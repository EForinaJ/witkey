package router

import (
	"server/internal/controller/account"
	"server/internal/controller/auth"
	"server/internal/controller/dashboard"
	"server/internal/controller/distribute"
	"server/internal/controller/site"
	"server/internal/controller/upload"
	"server/internal/controller/withdraw"
	"server/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LoadRouter(s *ghttp.Server) {
	s.Group("/api/v1/witkey", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			auth.NewV1(),
			site.NewV1(),
		).Middleware(middleware.Response)
		group.Bind(
			account.NewV1(),
			distribute.NewV1(),
			withdraw.NewV1(),
			dashboard.NewV1(),
			upload.NewV1(),
		).Middleware(middleware.Auth).Middleware(middleware.Response)
	})
}
