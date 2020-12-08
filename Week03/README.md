# 并发(concurrency)

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

* 如果事件e1发生在e2之前，那么我们就可以说事件e2发生在e1之后
* 如果e1既不发生在e2之前，也不发生在e2之后，那么我们就说e1和e2是同时发生的

同步(Synchronization)
初始化(Initialization)

* CPU 指令重排

* 编译器的编译重排

## Package sync

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

## References

如果包p引入（import）包q，那么q的init函数的结束先行发生于p的所有init函数开始 main.main函数的开始发生在所有init函数结束之后
