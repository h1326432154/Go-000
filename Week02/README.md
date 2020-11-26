# error

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

```sql
package aync

func Go(x func()){
    if err := recover();err != nil{

    }
    go x()
}
```

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

## Error Type

* 优点：可以携带更多上下文
* 缺点：但是调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会导致和调用者产生强耦合，从而导致 API 变得脆弱。

### Opaque errors

* 最灵活的错误处理策略，因为它要求代码和调用者之间的耦合最少

> Assert errors for behaviour, not type

## Handling Error

## Go 1.13 errors

## Go 2 Error Inspection
