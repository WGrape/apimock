# apimock
A lightweight solution in golang for API mock and automatic generation of API documentation / 一种适用于Go语言的接口Mock和接口文档自动生成的轻量级方案

## 一、背景介绍
在[《关于接口文档高效治理方案的研究和思考》](https://github.com/WGrape/Blog/issues/229) 这篇博客中，已经深入探讨过关于接口文档的现状、问题和解决方案。由于不同团队面临的问题不同，对接口文档的管理方式也不尽相同

本项目也是基于团队规模不同的观点考虑，提出了一种适用于Go语言的接口Mock和接口文档自动生成的轻量级方案，关于此方案的全面了解请参考文章[《基于数据Mock的接口治理方案设计与实现》](https://github.com/WGrape/Blog/issues/233)

## 二、快速安装

### 1、go get

```
go get github.com/WGrape/apimock@latest
```

### 2、go mod

```
module XXX

go 1.16

require (
    github.com/WGrape/apimock latest
)
```

## 三、如何使用

### 1、

## 五、完整示例
关于详细的使用示例请参考 [example](./example) ，其中的 [apidoc.md](./example/apidoc.md) 文件即为 apimock 工具自动生成的接口文档。


### 1、多种语言

如在```/mock/mock_zh.go```和```/mock/mock_en.go```文件中分别定义了中文和英文两个版本的```apiMock```结构。这样当需要不同语言文档时，都可以快速支持。
