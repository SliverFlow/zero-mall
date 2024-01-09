// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "server/app/system/cmd/api/internal/handler/base"
	business "server/app/system/cmd/api/internal/handler/business"
	category "server/app/system/cmd/api/internal/handler/category"
	file "server/app/system/cmd/api/internal/handler/file"
	log "server/app/system/cmd/api/internal/handler/log"
	menu "server/app/system/cmd/api/internal/handler/menu"
	order "server/app/system/cmd/api/internal/handler/order"
	product "server/app/system/cmd/api/internal/handler/product"
	role "server/app/system/cmd/api/internal/handler/role"
	user "server/app/system/cmd/api/internal/handler/user"
	"server/app/system/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/captcha",
				Handler: base.CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: base.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/system/base"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/findByUUID",
					Handler: user.UserFindByUUIDHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/changeRole",
					Handler: user.UserAdminChangeRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.UserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/changeStatus",
					Handler: user.UserChangeStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/delete",
					Handler: user.UserDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/updateByUUID",
					Handler: user.UserUpdateByUUIDHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role/create",
					Handler: role.RoleCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/update",
					Handler: role.RoleUpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/menu/listByUUID",
					Handler: menu.MenuListByUUIDHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/listAllByRole",
					Handler: menu.MenuListAllByRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/changeStatus",
					Handler: menu.MenuChangeStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/create",
					Handler: menu.MenuCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/update",
					Handler: menu.MenuUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/find",
					Handler: menu.MenuFindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/delete",
					Handler: menu.MenuDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/business/list",
					Handler: business.BusinessListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/business/changeStatus",
					Handler: business.BusinessChangeStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/business/find",
					Handler: business.BusinessFindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/business/update",
					Handler: business.BusinessUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/business/dict",
					Handler: business.BusinessDictHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/category/listAll",
					Handler: category.CategoryListAllHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category/create",
					Handler: category.CategoryCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category/changeStatus",
					Handler: category.CategoryChangeStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category/find",
					Handler: category.CategoryFindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category/update",
					Handler: category.CategoryUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category/batchDelete",
					Handler: category.CategoryBatchDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category/delete",
					Handler: category.CategoryDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/category/treeList",
					Handler: category.CategoryTreeListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/product/create",
					Handler: product.ProductCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/product/list",
					Handler: product.ProductListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/product/delete",
					Handler: product.ProductDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/product/changeStatus",
					Handler: product.ProductChangeStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/product/find",
					Handler: product.ProductFindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/product/update",
					Handler: product.ProductUpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/order/list",
					Handler: order.OrderListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/log/computer",
					Handler: log.ComputerHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/token/:path",
					Handler: file.FileTokenHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system/v1"),
	)
}
