# golang 追踪堆栈错误包
> Golang tracks stack error package. 追踪堆栈错误包

## 输出堆栈信息
基于`github.com/pkg/errors`包进行封装,方便开发项目中使用. 

这个包本身是可以输出堆栈信息的,但是输出是原始的堆栈信息

存储在日志里查看不友好,而`github.com/yezihack/e`对堆栈信息进行封装了.

而且重新实现了系统自带的`error`接口,提供 `code` 自定义功能.

## 堆栈信息对比
|原始堆栈信息|封装后的堆栈信息|
|---|---|
|``|``|