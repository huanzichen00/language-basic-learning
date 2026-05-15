# Go 后端基础学习计划

## 学习定位

你已经掌握 Java 和 C，所以学习 Go 时不需要把精力平均花在所有语法点上。重点应该放在三件事上：

1. 理解 Go 和 Java/C 的设计差异。
2. 养成 Go 项目的代码组织、错误处理和测试习惯。
3. 为以后做 Go 后端开发打基础，包括 HTTP、并发、数据库和工程化。

本计划采用聊天协作方式推进：

- 我在聊天窗口给你讲解、代码和练习。
- 你把代码手动写入本地文件。
- 你运行代码后，把结果或报错贴回来。
- 我根据运行结果继续讲解、修改或加练习。

## 总体目标

完成本阶段后，你应该能够：

- 使用 `go run`、`go build`、`go test` 和 `go mod`。
- 熟练掌握 Go 基础语法、切片、map、struct、interface。
- 理解 Go 的错误处理方式，并能写出清晰的错误返回逻辑。
- 使用 goroutine、channel、WaitGroup 和 context 编写基础并发程序。
- 使用标准库 `net/http` 编写简单 REST API。
- 理解 Go 后端项目的基本分层方式。

## 学习节奏

建议每天学习 60 到 90 分钟：

- 15 到 20 分钟：理解概念。
- 30 到 40 分钟：手写代码。
- 15 到 20 分钟：运行、调试、记录问题。
- 10 分钟：对比 Java/C，总结差异。

每周至少完成 3 个小练习。不要只看代码，一定要手写和运行。

## 第 1 周：环境、基础语法和函数

### 目标

建立 Go 的基本语感，理解 Go 程序如何组织和运行。

### 内容

- 安装和检查 Go 环境。
- `go version`
- `go run`
- `go build`
- `go mod init`
- `package main`
- `import`
- 变量和常量。
- 基础类型。
- `if`
- `for`
- `switch`
- 函数。
- 多返回值。
- 指针基础。

### 和 Java/C 的对比重点

- Go 没有 class。
- Go 没有 `while`，只有 `for`。
- Go 可以返回多个值。
- Go 的指针比 C 简化，没有指针运算。
- Go 的入口函数是 `main` 包里的 `main` 函数。

### 练习

1. 写一个计算两个整数和、差、积、商的程序。
2. 写一个判断闰年的函数。
3. 写一个函数，输入两个字符串，返回交换后的结果。
4. 写一个简单菜单程序，用 `switch` 处理用户选择。

## 第 2 周：数组、切片、map 和字符串

### 目标

掌握 Go 中最常用的数据容器。

### 内容

- array 和 slice 的区别。
- slice 的声明、截取和追加。
- `append`
- `len`
- `cap`
- `range`
- map 的增删改查。
- 字符串、byte 和 rune。
- 中文字符串遍历。

### 和 Java/C 的对比重点

- slice 类似动态数组，但底层有数组、长度和容量。
- map 类似 Java 的 `HashMap`，但语法更轻。
- Go 字符串是只读字节序列，处理中文时要注意 rune。

### 练习

1. 找出整数切片中的最大值和最小值。
2. 反转一个整数切片。
3. 统计字符串中每个字符出现的次数。
4. 实现一个简单通讯录，支持新增、查询和删除联系人。

## 第 3 周：struct、方法和 interface

### 目标

学会用 Go 的方式组织业务对象，而不是照搬 Java 的类继承模型。

### 内容

- struct 定义。
- struct 初始化。
- 匿名字段。
- 方法。
- 值接收者和指针接收者。
- interface。
- 空接口 `any`。
- 组合代替继承。

### 和 Java/C 的对比重点

- Go 没有类和继承。
- 方法是绑定到类型上的函数。
- interface 是隐式实现的，不需要 `implements`。
- Go 更鼓励小接口。

### 练习

1. 定义 `User`、`Book`、`Order` 三个结构体。
2. 给 `User` 添加修改昵称和打印信息的方法。
3. 定义一个 `Notifier` 接口，实现 Email 和 SMS 两种通知方式。
4. 用组合实现一个简单的管理员用户模型。

## 第 4 周：错误处理、包管理、文件操作和测试

### 目标

进入 Go 工程化的基本写法。

### 内容

- `error`
- `errors.New`
- `fmt.Errorf`
- 错误包装。
- `defer`
- 文件读取。
- 文件写入。
- 包拆分。
- `go mod`
- `go test`
- 表驱动测试。

### 和 Java/C 的对比重点

- Go 没有 Java 风格的 `try-catch`。
- Go 倾向显式处理错误。
- `defer` 常用于关闭资源。
- Go 的测试文件以 `_test.go` 结尾。

### 练习

1. 写一个安全除法函数，除数为 0 时返回错误。
2. 读取一个文本文件，统计行数和单词数。
3. 把用户信息保存到文件。
4. 给基础计算函数写表驱动测试。

## 第 5 周：并发基础

### 目标

理解 Go 的核心优势之一：轻量级并发。

### 内容

- goroutine。
- channel。
- buffered channel。
- `select`
- `sync.WaitGroup`
- `sync.Mutex`
- `context`
- 超时控制。

### 和 Java/C 的对比重点

- goroutine 比传统线程更轻量。
- channel 用于 goroutine 之间通信。
- Go 常见思想是通过通信共享内存。
- 后端请求链路中 `context` 非常重要。

### 练习

1. 启动 5 个 worker 并等待它们全部完成。
2. 用 channel 汇总多个 goroutine 的计算结果。
3. 写一个带超时控制的任务。
4. 模拟多个用户同时访问计数器，并用 Mutex 保证安全。

## 第 6 周：HTTP 后端入门

### 目标

用 Go 标准库写一个简单后端服务，为后续 Gin、GORM、MySQL、Redis 做准备。

### 内容

- `net/http`
- handler。
- route。
- JSON 编码和解码。
- HTTP 方法。
- 状态码。
- REST API。
- 中间件概念。
- 简单项目分层。

### 小项目：用户管理 API

先用内存存储实现：

- `GET /users`
- `GET /users/{id}`
- `POST /users`
- `PUT /users/{id}`
- `DELETE /users/{id}`

### 练习

1. 返回固定用户列表。
2. 解析 JSON 请求体创建用户。
3. 根据 ID 查询用户。
4. 添加简单日志中间件。
5. 添加请求超时控制。

## 后续进阶路线

完成 6 周基础后，可以进入 Go 后端进阶：

1. Gin 框架。
2. GORM。
3. MySQL。
4. Redis。
5. JWT 登录认证。
6. 配置管理。
7. 日志系统。
8. 单元测试和接口测试。
9. Docker 部署。
10. 简单微服务和 gRPC。

## 每节课固定协作格式

之后每次正式学习时，可以按下面格式推进：

```text
本节目标
需要新建或修改的文件
完整代码
运行命令
运行后应该看到什么
常见错误
练习题
对比 Java/C 的理解点
```

## 建议目录结构

当前学习进度比原计划更快，所以目录不再按现实周数组织，而是按主题编号组织：

```text
go/
  01-basics/
    01-functions-errors/
    02-types-formatting/
    03-slice/
    04-map/
    05-struct-method/
    06-interface/
    07-package-test/
  02-error-file-test/
  03-concurrency/
  04-http/
  05-database/
  06-projects/
```

每个小练习单独建一个目录，方便使用 `go mod` 和运行。

## 当前进度

已经完成的内容主要在：

```text
01-basics/
```

之后的新内容建议从这里开始：

```text
02-error-file-test/
  01-error-wrapping/
```

下一阶段重点是 `nil`、指针、构造函数、错误包装、文件读写和更系统的测试。
