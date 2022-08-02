DuxGO Example
====================

# 概述
DuxGO 是一款基于 go-echo 框架整合常用的 ORM、日志、队列、缓存等 web 开发常用功能，提供了一个简单、易用、灵活的框架。

本示例集成了 duxgo 框架与 duxgo-ui UI扩展包与 duxgo-admin 后台管理包，本示例用于 duxgo 的基础使用示例。

# 依赖

- Go 1.18+
- Mysql 5.7+
- Redis 5.0+

# 安装

将该仓库代码导出到独立目录使用 go install 安装：

```sh
go install
```

# 使用方法

## 1. 修改数据库配置

```
config/database.toml
```

## 2. 运行框架


```go
go run main.go
```

# 访问地址

```
0.0.0.0:8080
```

# 后台地址

```
0.0.0.0:8080/admin
```
