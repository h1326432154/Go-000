# error

## 目录
  
  [作业](#作业)

## Error vs Exception

|语言|错误处理方式|
|--|--|
|C|单返回值，一般通过传递指针作为入参，返回值为 int 表示成功还是失败。|
|C++|引入了 exception，但是无法知道被调用方会抛出什么异常。|
|java|引入了 checked exception，方法的所有者必须申明，调用者必须处理。在启动时抛出大量的异常是司空见惯的事情，并在它们的调用堆栈中尽职地记录下来。Java 异常不再是异常，而是变得司空见惯了。它们从良性到灾难性都有使用，异常的严重性由函数的调用者来区分。|
|go|如果一个函数返回了 value, error，你不能对这个 value 做任何假设，必须先判定 error。唯一可以忽略 error 的是，如果你连 value 也不关心。|

* Go painc

1. main 强依赖 初始化不成功 painc
2. 配置错误  init

* error 优点

* 简单
* 考虑失败，而不是成功
* 没有隐藏的控制流
* 完全交给程序员控制
* Error are values

### sentinel error

* 全局预定义的特定错误（错误码）

* 包之间尽量避免sentinel error（减少依赖）

```sql
if err == ErrSomething{}
```

### Error Type

Error type 是实现了 error 接口的自定义类型。

* 优点：可以携带更多上下文
* 缺点：但是调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会导致和调用者产生强耦合，从而导致 API 变得脆弱。

### Opaque errors

* 最灵活的错误处理策略，因为它要求代码和调用者之间的耦合最少

> Assert errors for behaviour, not type

## Handling Error

* you should only handle errors once. Handling an error means inspecting the error value, and making a single decision.
* 日志记录与错误无关且对调试没有帮助的信息应被视为噪音，应予以质疑。记录的原因是因为某些东西失败了，而日志包含了答案。
* 在你的应用代码中，使用 errors.New 或者 errros.Errorf 返回错误。
* 如果调用其他的函数，通常简单的直接返回。
* 如果和其他库进行协作，考虑使用 errors.Wrap 或者 errors.Wrapf 保存堆栈信息。
* 直接返回错误，而不是每个错误产生的地方到处打日志。
* 在程序的顶部或者是工作的 goroutine 顶部(请求入口)，使用 %+v 把堆栈详情记录。
* 使用 errors.Cause 获取 root error，再进行和 sentinel error 判定。
* 选择 wrap error 是只有 applications 可以选择应用的策略。
* 如果函数/方法不打算处理错误，那么用足够的上下文 wrap errors 并将其返回到调用堆栈中。
* 一旦确定函数/方法将处理错误，错误就不再是错误。如果函数/方法仍然需要发出返回，则它不能返回错误值。

## Wrap errors 错误包装

如果错误没有就地处理，需要向调用者输出，最终在调用栈的根部需要处理错误，这时将错误输出， 打印出来的只有基本的错误信息，缺少错误生成时的 file:line 信息、没有调用堆栈。

为了追踪错误，有一种做法是使用 fmt.Errorf 以原 err 加一些描述信息生成新的 error 抛出， 但这种模式与 sentinel errors 或 type assertions 的使用不兼容： 破坏了原始错误，导致等值判定失败。

## Go 1.13 errors

Unwrap、Is、As

## Go 2 Error Inspection

[Go 2 Error Inspection](https://go.googlesource.com/proposal/+/master/design/29934-error-values.md)

## 作业

Q1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A1. 需要 wrap err, 抛给上层

[示例代码](./main.go)
