# 1、项目介绍

使用 go-zero 开发的微服务电商系统。（开发阶段）

# 2、技术栈

### 1、前端

开发框架：vue3

状态管理：pinia 

UI 组件库：Element-Plus

路由管理：vue-router

### 2、后端

微服务框架：go-zero

网关：kong 

服务注册 / 发现：etcd 

鉴权：jwt 

ORM框架：gorm

项目部署：docker-compose

日志系统：ELK

链路追踪：jaeger

服务监控：prometheus 、grafana

### 3、数据库

关系型数据库：mysql

非关系型数据库：redis 

# 3、目录结构

```sh
|- server 		# 后端服务主目录
|- web
	|- system 	# 商户 or 管理员管理系统
	|- mall   	# 用户显示 & 主体界面
```

