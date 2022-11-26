# EchoSkeleton

<div align="center"> 项目正在快速开发</div><br>
<div align="center">

<a href="https://github.com/00ty/EchoSkeleton/blob/master/LICENSE">
<img src="https://img.shields.io/github/license/00ty/EchoSkeleton"></img></a>
<a href="https://github.com/wejectchen/Ginblog">
<img src="https://img.shields.io/github/stars/00ty/EchoSkeleton?label=GitHub">
</a>
</div><br>

![HmPEmVkR.jpg](https://img1.imgtp.com/2022/11/26/HmPEmVkR.png)

## 介绍

这是一个基于Golang语言Echo框架的web mvc项目骨架，专注于前后端分离的单体应用业务场景（非微服务方向）,其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。

- 欢迎你的维护！
- 日志（待完善日志文件记录）
- gorm (待完善日志文件记录)
- 待完善jwt封装

## 目录结构
```shell
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  config.yml // 配置文件
│  LICENSE
│  README.md
│          
├─ api // 项目主路径
│  ├─ common // 公共文件夹
│  ├─ controller // 控制器
│  ├─ dao
│  ├─ entity
│  ├─ routes // 路由
│  └─ service
│  
├─ utils // 项目公用工具库 
│  └─config // 配置 
│ 
├─ cmd // 项目入口   
│  └─main.go
├─ models // 数据模型层    
├─ database // 数据库文件夹
└─────────────────────────────
```

## 开始

```shell
运行
go run ./cmd/main.go
打包部署
go build ./cmd/main.go
```
