# 并发(concurrency)

## 目录
  
  [作业](#作业)

> Don’t communicate by sharing memory; share memory by communicating.

## Goroutine

理念:

* Keep yourself busy or do the work yourself

    如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它( go func() )更简单

* Leave concurrency to the caller

    将并发逻辑交给调用者

* Never start a goroutine without knowning when it will stop

    把握goroutine生命周期

## Memory model [内存模型](https://www.jianshu.com/p/97a345f47cfd)

> Go内存模型指定条件，在该条件下，可以保证一个goroutine中的变量读取可以观察到不同goroutine写入同一个变量而产生的值

注：任何一个Go程序中都不会只有一个goroutine的存在，即使你没有显示声明过（go关键字声明），程序在启动时除了有一个main的goroutine存在之外，至少还会隐式的创建一个goroutine用于gc，使用runtime.NumGoroutine()可以得到程序中的goroutine的数量

### Happens Before原则

> 如果包p引入（import）包q，那么q的init函数的结束先行发生于p的所有init函数开始 main.main函数的开始发生在所有init函数结束之后

* 如果事件e1发生在e2之前，那么我们就可以说事件e2发生在e1之后
* 如果e1既不发生在e2之前，也不发生在e2之后，那么我们就说e1和e2是同时发生的

同步(Synchronization)
初始化(Initialization)

* CPU 指令重排

* 编译器的编译重排

## Package sync

* 竞态条件:一旦数据被多个线程共享，那么就很可能会产生争用和冲突的情况。这种情况也被称为竞态条件（race condition）
* 临界区:多个并发运行的线程对这个共享资源的访问是完全串行的。只要一个代码片段需要实现对共享资源的串行化访问，就可以被视为一个临界区（critical section）
* 同步工具:临界区总是需要受到保护的，否则就会产生竞态条件。施加保护的重要手段之一，就是使用实现了某种同步机制的工具，也称为同步工具。

### 同步工具

#### Mutex(互斥锁)

* 不要重复锁定互斥锁
* 不要忘记解锁互斥锁，必要时使用defer语句
* 不要对尚未锁定或者已解锁的互斥锁解锁
* 不要在多个函数之间直接传递互斥锁。

> 这种由 Go 语言运行时系统自行抛出的 panic 都属于致命错误，都是无法被恢复的，调用recover函数对它们起不到任何作用。也就是说，一旦产生死锁，程序必然崩溃。

锁饥饿

Barging: 提高了吞吐量，但不公平
Hands-off: 吞吐量有所降低，但公平
Spinning: 性能开销大
Go 1.8 使用了Barging和Spinning结合实现，自旋几次后就会park
Go 1.9 添加了饥饿模式，如果等待锁1ms, unlock会hands-off把锁丢给第一个等待者,此时同样代码g1:57 g2:10

#### RWMutex(读写锁)

* 在写锁已被锁定的情况下再试图锁定写锁，会阻塞当前的 goroutine
* 在写锁已被锁定的情况下试图锁定读锁，也会阻塞当前的 goroutine
* 在读锁已被锁定的情况下试图锁定写锁，同样会阻塞当前的 goroutine
* 在读锁已被锁定的情况下再试图锁定读锁，并不会阻塞当前的 goroutine

#### Atomic(原子操作)

sync/atomic包中的函数可以做的原子操作有：加法（add）、比较并交换（compare and swap，简称 CAS）、加载（load）、存储（store）和交换（swap）。

64位操作系统一次刷数据8byte，刷数据操作便是一个不可分割的原子操作

#### errGroup

* 核心原理 利用sync.WaitGroup管理并执行goroutine
* 主要功能
  * 并行工作流
  * 处理错误 或者 优雅降级
  * context 传播与取消
  * 利用局部变量+闭包

* 设计缺陷 --- 改进
  * 没有捕获panic，导致程序异常退出 --- 改进 加defer recover
  * 没有限制goroutine数量，存在大量创建goroutine --- 改进 增加一个channel用来消费func
  * WithContext 返回的context可能被异常调用，当其在errgroup中被取消时，影响其它函数 --- 改进 代码内嵌context

#### sync.Pool

* 保存与复用临时对象
* 降低GC压力
* 不能放链接类型，有可能导致链接泄漏

Copy-On-Write

## chan

channel通信是goroutine同步的主要方法。每一个在特定channel的发送操作都会匹配到通常在另一个goroutine执行的接收操作。

> 一个通道相当于一个先进先出（FIFO）的队列。也就是说，通道中的各个元素值都是严格地按照发送的顺序排列的，先被发送通道的元素值一定会先被接收。元素值的发送和接收都需要用到操作符<-。我们也可以叫它接送操作符。一个左尖括号紧接着一个减号形象地代表了元素值的传输方向。

* 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
* 发送操作和接收操作中对元素值的处理都是不可分割的。
* 发送操作在完全完成之前会被阻塞。接收操作也是如此。

元素值从外界进入通道时会被`复制`。更具体地说，进入通道的并不是在接收操作符右边的那个元素值，而是它的副本。
另一方面，元素值从通道进入外界时会被`移动`。这个移动操作实际上包含了两步，第一步是生成正在通道中的这个元素值的副本，并准备给到接收方，第二步是删除在通道中的这个元素值。

>注:如果通道关闭时，里面还有元素值未被取出，那么接收表达式的第一个结果，仍会是通道中的某一个元素值，而第二个结果值一定会是true。

所以除非有特殊的保障措施，不要让接收方关闭通道，而应当让发送方做这件事。

## Package context

Request-scoped context

* 实现传递数据，搞定超时控制，或者级联取消(显示传递)
* context集成到API
  * 函数首参为context
  * 创建对象时携带context对象: WithContext

Don't store Contexts inside a struct type

* 不要把context放到结构体里，然后再把结构体当参数传输

context.WithValue

* 从子向父递归查询key-value
* Background、TODO
* Debugging or tracing data is safe to pass in a Context
* context.WithValue 只读、安全 --- 染色、API重要性、Trace
* 禁止在context中挂载与业务逻辑耦合的东西，不能放一些奇奇怪怪的东西进去
* 如果有必要修改context的内容，请使用COW:
  * 从源ctx获取到v1
  * 复制v1到v2
  * 修改v2
  * 将v2重新挂载到ctx,产生ctx2
  * 将ctx2向下传递
* gin的context.Next有缺陷，应参考grpc的middleware
* 计算密集型耗时短，一般不处理超时。
* go标准网络库可被托管，~~吊打其它语言业务、中间件，~~不会因为超时导致oom。kratos案例
* 当一个context被cancel时，所有子context都会被cancel
* 一定要cancel 否者context会泄漏

## References

## 作业

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

[提交地址](https://github.com/Go-000/Go-000/issues/69)

```go
func main() {
    eg := errgroup.Group{}
    serverErr := make(chan error, 1)
    sigC := make(chan os.Signal, 1)

    s := http.Server{Addr: ":8080"}

    eg.Go(func() error {
        go func() {
            serverErr <- s.ListenAndServe()
        }()
        select {
            case err := <-serverErr:
            close(sigC)
            close(serverErr)
            return err
        }
    })

    eg.Go(func() error {
        signal.Notify(sigC,
            syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
        <-sigC
        return s.Shutdown(context.TODO())
    })

    log.Println(eg.Wait())
}
```
