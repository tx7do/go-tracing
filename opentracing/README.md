# OpenTracing

## 数据模型

- Trace （调用链）：一个 Trace 代表一个事务或者流程在（分布式）系统中的执行过程。例如来自客户端的一个请求从接收到处理完成的过程就是一个 Trace。
- Span（跨度）：Span 是分布式追踪的最小跟踪单位，一个 Trace 由多段 Span 组成。可以被理解为一次方法调用, 一个程序块的调用, 或者一次 RPC/数据库访问。只要是一个具有完整时间周期的程序访问，都可以被认为是一个 Span。
- SpanContext（跨度上下文）：分布式追踪的上下文信息，包括 Trace id，Span id 以及其它需要传递到下游服务的内容。一个 OpenTracing 的实现需要将 SpanContext 通过某种序列化协议 (Wire Protocol) 在进程边界上进行传递，以将不同进程中的 Span 关联到同一个 Trace 上。对于 HTTP 请求来说，SpanContext 一般是采用 HTTP header 进行传递的。
