# 后端服务

后端服务，采用Go语言开发，提供用户认证、用户管理等核心功能。

## 项目概述

本项目是一个基于Go和Gin框架的RESTful API服务，提供后端支持。主要功能包括用户认证、用户管理和数据库操作。

## 技术栈

- **编程语言**: Go
- **Web框架**: [Gin Web Framework](https://gin-gonic.com/)
- **数据库**: PostgreSQL 
- **ORM框架**: GORM
- **跨域处理**: gin-contrib/cors
- **日志**: 结构化日志系统

## 项目结构

```
.
├── main.go                 # 应用入口
├── config/                 # 配置管理
│   ├── config.go          # 配置初始化
│   ├── config.yaml        # 配置文件
│   └── db.go              # 数据库连接配置
├── controllers/           # 控制器层 - 业务逻辑处理
│   ├── AuthClient.go      # 客户端认证
│   ├── AuthPolice.go      # 认证
│   ├── CreateUser.go      # 用户创建
│   ├── ManagerAuth.go     # 权限管理
│   └── ...
├── models/                # 数据模型层
│   └── sc_user_police.go  # 用户模型
├── router/                # 路由配置
│   └── router.go          # 路由定义
├── global/                # 全局配置
│   └── global.go          # 全局变量/常量
└── utils/                 # 工具函数
    └── utils.go           # 通用工具
```

## 快速开始

### 环境要求

- Go 1.16 或更高版本
- PostgreSQL 12 或更高版本

### 安装依赖

```bash
go mod download
# 或
go mod tidy
```

### 配置文件

修改 `config/config.yaml` 文件中的数据库连接信息：

```yaml
database:
  host: "your_db_host"
  port: 5432
  user: "your_db_user"
  password: "your_db_password"
  dbname: "your_db_name"

app:
  name: "sc_police_api"
  node: "dev"
  port: ":8088"
```

### 运行应用

```bash
go run main.go
```

应用默认运行在 `http://localhost:8088`

## API 端点

### 认证相关

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/api/auth/checkToken` | 检查/验证Token |

### 用户相关

| 方法 | 端点 | 描述 |
|------|------|------|
| POST | `/api/user/createUser` | 创建新用户 |

## 数据模型

### ScUserPolice

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键 |
| code | varchar(12) | 编码（不为空） |
| police_id | varchar(20) | ID（不为空，唯一） |
| password | varchar(255) | 密码（不为空） |
| name | varchar(50) | 姓名 |
| works | jsonb | 工作信息（JSON格式） |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |
| deleted_at | timestamp | 删除时间（软删除） |

## 文件上传

应用支持静态文件服务，`/uploads` 路径映射到本地 `./uploads` 目录。

```
http://localhost:8088/uploads/file.jpg
```

## 日志配置

日志配置在 `config/config.yaml` 中定义：

```yaml
log:
  level: "info"              # 日志级别
  filename: "log/sc_police_api.log"  # 日志文件路径
  max_size: 100              # 单个文件最大大小（MB）
  max_age: 30                # 日志保留天数
  max_backups: 10            # 最大备份数
```

## 功能特性

- ✅ 用户认证系统（Token验证）
- ✅ 用户管理（用户创建）
- ✅ 权限管理
- ✅ CORS跨域支持
- ✅ 数据库自动迁移
- ✅ 静态文件服务
- ✅ 结构化日志记录

## 开发指南

### 添加新的API端点

1. 在 `controllers/` 目录中创建新的控制器文件
2. 在 `router/router.go` 中注册路由
3. 如需新数据模型，在 `models/` 中定义

### 数据库迁移

应用启动时会自动执行数据库表迁移：

```go
if err := controllers.MigrateTables(); err != nil {
    log.Fatalf("数据库表迁移失败: %v", err)
}
```

## 环境部署

### 开发环境

```bash
go run main.go
```

### 生产环境

```bash
# 编译
go build -o sc_police_api

# 运行
./sc_police_api
```

## 常见问题

### 数据库连接失败

- 检查 `config/config.yaml` 中的数据库配置是否正确
- 确保PostgreSQL服务正在运行
- 验证数据库用户是否有访问权限

### 跨域问题

本应用已启用CORS中间件，支持跨域请求。如需自定义CORS设置，可在 `router/router.go` 中修改：

```go
r.Use(cors.Default())
```

## 1、初始化
- go mod init sc_policing_api
## 2、安装依赖
- go mod tidy
## 3、编译
- GOOS=linux GOARCH=amd64 go build -o libu-admin main.go

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用相关许可证保护。

## 联系方式

如有问题或建议，请联系项目维护者。

---

**最后更新**: 2026年2月13日
