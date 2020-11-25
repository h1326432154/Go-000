# Go-架构实践

## 目录

* [微服务概览](#微服务概览)
* [微服务设计](#微服务设计)
* [gRPC & 服务发现](#GgRPC&服务发现)
* [多集群 & 多租户](#多集群&多租户)

## 微服务概览

> 微服务 是 SOA(面向服务)的一种实践

围绕业务功能构建的，服务关注单一业务，服务 间采用轻量级的通信机制，可以全自动独立部署， 可以使用不同的编程语言和数据存储技术。
微服务的架构与传统架构的不同之处在于，微服务的每个服务与其数据库都是独立的，可以无依赖地进行部署。

* 优点
  * 原子服务
  * 独立进程
  * 隔离部署
  * 去中心化
* 劣势
  * 基础设施建设 复杂度高(监控报警系统)
  * 请求放大 单次请求 下游请求放大百倍
  * 分区数据架构
  * 测试基于微服务架构的应用复杂
  * 连锁故障 服务模块间的依赖

### 组件化服务

* kit: 一个微服务的基础框架
* service： 业务代码+kit依赖+第三方依赖
* rpc+message queue(kafka):轻量级通讯

本质上等同于 多个微服务组合(compose)完成了一个完整的用户场景(usercase)

### 按业务组织服务

> 大前端(移动/Web) => 网关接入 =>业 务服务 =>平台服务 =>基础设施(PaaS/Saas) 开发团队对软件在生产环境的运行负全部责任！

### 去中心化

* 数据去中心化（每个服务独享数据存储设施，利于服务独立性）
* 治理去中心化
* 技术去中心化

### 基础设施自动化

* CICD：Gitlab + Gitlab Hooks + k8s
* Testing：测试环境、单元测试、API自动化测试
* 在线运行时：k8s，以及一系列 Prometheus、ELK、Conrtol Panle

### 可用性 & 兼容性设计

* 隔离
* 超时控制
* 负载保护
* 限流
* 降级
* 重试
* 负载均

`Design For Failure`思想

`伯斯塔尔法则`Be conservative in what you send, be liberal in what you accept.  
发送时要保守，接收时要开放。按照伯斯塔尔法则的思想来设计和实现服务时，发送的数据要更保守， 意味着最小化的传送必要的信息，接收时更开放意味着要最大限度的容忍冗余数据，保证兼容性。

## 微服务设计

### API-Gateway

统一的协议出口，在服务内进行大量的 dataset join，按照业务场景来设计粗粒度的 API，给后续服务的演进带来的很多优势

> 移动端 -> API Gateway -> BFF -> Mircoservice
在 FE Web业务中，BFF 可以是 nodejs 来做服务端渲染 (SSR，Server-Side Rendering)，注意这里忽略了上 游的 CDN、4/7层负载均衡(ELB)

### 微服务划分

在实际项目中通常会采用两种不同的方式划分服务边界：

* 通过业务职能(Business Capability)
* DDD 的限界上下文(Bounded Context)

CQRS(Command Query Responsibility Segregation) 将应用程序分为两部分：命令端和查询端。命令端处理程序创建、更新和删除请求，并在数据更改时发出事件。查询端通过针对一个或多个物化视图执行查询来处理查询，这些物化视图通过订阅数据更改时发出的事件流而保持最新。

### 微服务安全

> 完整请求流程： API Gateway --> BFF --> Service Biz Auth --> JWT --> Req args

在服务内部，要区分身份认证和授权，一般有三种：

Full Trust
Half Trust
Zero Trust

## gRPC

> A high-performance, open-source universal RPC framework

优势:

* 多语言：语言中立，支持多种语言。
* 轻量级、高性能：序列化支持 PB(Protocol Buffer)和 JSON，PB 是一种语言无关的高性能序列化框架。
* 可插拔
* IDL：基于文件定义服务，通过proto3工具生成指定 语言的数据结构、服务端接口以及客户端 Stub。
* 设计理念
* 移动端：基于标准的HTTP2设计，支持双向流、消 息头压缩、单TCP的多路复用、服务端推送等特性， 这些特性使得 gRPC在移动端设备上更加省电和节 省网络流量。

* 服务而非对象、消息而非引用：促进微服务的系统 间粗粒度消息交互设计理念。
* 负载无关的：不同的服务需要使用不同的消息类型 和编码，例如protocol buffers、JSON、XML和Thrift。
* 流：Streaming API。
* 阻塞式和非阻塞式：支持异步和同步处理在客户端 和服务端间交互的消息序列。
* 元数据交换：常见的横切关注点，如认证或跟踪， 依赖数据交换。
* 标准化状态码：客户端通常以有限的方式响应API调 用返回的错误

gRPC - HealthCheck

## 服务发现

* 客户端发现
  * 一个服务实例被启动时，它的网络地址会被写到注册表上；
  * 当服务实例终止时，再从注册表中删除；
  * 这个服务实例的注册表通过心跳机制动态刷新；
  * 客户端使用一个负载均衡算法，去选择一个可用的服务实例，来响应这个请求。
  * 直连，比服务端发现少一次网络跳转，Consumer 需要内置特定的服务发现客户端和发现逻辑
* 服务端发现
  * 客户端通过负载均衡器向一个服务发送请求
  * 这个负载均衡器会查询服务注册表，并将请求路由到可用的服务实例上。
  * 服务实例在服务注册表上被注册和注销(Consul Template+Nginx， kubernetes+etcd)。
  * Consumer 无需关注服务发现的细节，只需知道服务的 DNS 域名即可，支持异构语言开发，需要基础设施支撑，多一次网络跳转，可能有性能损耗

## 多集群 & 多租户

### 多集群

* 多集群的必要性
* 从单一集群考虑，多节点保证可用性，N+2 的节点来冗余节点
* 单一集群故障带来的影响考虑冗余多套集群
* 单机房内的机房故障导致的问题

### 多租户

> 在一个微服务架构中允许多系统共存是利用微服务稳定性以及模块化最有效的方式之一，这种方 式一般被称为多租户(multi-tenancy)。租户可以是测试，金丝雀发布，影子系统(shadow systems)，甚至服务层或者产品线，使用租户能 够保证代码的隔离性并且能够基于流量租户做路由决策。

多租户架构本质上描述为：跨服务传递请求携带上下文(context)，数据隔离的流 量路由方案。

## 相关文档

* 课堂提问整理

  * [第一颗](https://shimo.im/docs/x8dxHkQRcdCHX8j3)
  * [第二颗](https://shimo.im/docs/WxJp66WCtjVwKDK3)

* 参考文献
  * 《SRE：Google运维解密》
  * 《UNIX环境高级编程第3版》
  * [康威定律](https://zh.wikipedia.org/wiki/%E5%BA%B7%E5%A8%81%E5%AE%9A%E5%BE%8B)
  * [伯斯塔尔法则-鲁棒性原理 (Robustness Principle)](https://en.wikipedia.org/wiki/Robustness_principle)
  * [CAP定理](https://zh.wikipedia.org/wiki/CAP%E5%AE%9A%E7%90%86)

* 相关链接
  * [左耳微服务](https://time.geekbang.org/column/article/11116)  `其中介绍大量参考文献`
