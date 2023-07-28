// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	private "server/app/user/cmd/api/internal/handler/private"
	public "server/app/user/cmd/api/internal/handler/public"
	"server/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: public.HealthHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: public.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: public.CreateHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/find",
					Handler: private.FindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: private.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: private.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/user"),
	)
}
