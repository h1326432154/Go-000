# 并发(concurrency)

> Concurrency is not Parallelism.

## Goroutine

理念:

* Keep yourself busy or do the work yourself

    如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它( go func() )更简单

* Leave concurrency to the caller

    将并发逻辑交给调用者

* Never start a goroutine without knowning when it will stop

    把握goroutine生命周期

## Memory model

* [内存模型](https://www.jianshu.com/p/5e44168f47a3)

* CPU 指令重排

* 编译器的编译重排

## Package sync

## chan

## Package context

## References

如果包p引入（import）包q，那么q的init函数的结束先行发生于p的所有init函数开始 main.main函数的开始发生在所有init函数结束之后
