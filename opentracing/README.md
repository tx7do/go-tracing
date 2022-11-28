# OpenTracing

## 数据模型

- Trace （调用链）：一个 Trace 代表一个事务或者流程在（分布式）系统中的执行过程。例如来自客户端的一个请求从接收到处理完成的过程就是一个 Trace。
- Span（跨度）：Span 是分布式追踪的最小跟踪单位，一个 Trace 由多段 Span 组成。可以被理解为一次方法调用, 一个程序块的调用, 或者一次 RPC/数据库访问。只要是一个具有完整时间周期的程序访问，都可以被认为是一个 Span。
- SpanContext（跨度上下文）：分布式追踪的上下文信息，包括 Trace id，Span id 以及其它需要传递到下游服务的内容。一个 OpenTracing 的实现需要将 SpanContext 通过某种序列化协议 (Wire Protocol) 在进程边界上进行传递，以将不同进程中的 Span 关联到同一个 Trace 上。对于 HTTP 请求来说，SpanContext 一般是采用 HTTP header 进行传递的。

## Docker部署服务器

### Jaeger

```shell
docker pull jaegertracing/all-in-one:latest

docker run -d \
    --name jaeger \
    -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -e COLLECTOR_OTLP_ENABLED=true \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 4317:4317 \
    -p 4318:4318 \
    -p 14250:14250 \
    -p 14268:14268 \
    -p 14269:14269 \
    -p 9411:9411 \
    jaegertracing/all-in-one:latest
```

- API：<http://localhost:14268/api/traces>
- Zipkin API：<http://localhost:9411/api/v2/spans>
- 后台: <http://localhost:16686>

### Zipkin

```shell
docker pull openzipkin/zipkin:latest

docker run -d \
    --name zipkin \
    -p 9411:9411 \
    openzipkin/zipkin:latest
```

- API：<http://localhost:9411/api/v2/spans>
- 后台: <http://localhost:9411>

### SkyWalking

```bash
docker pull apache/skywalking-oap-server:latest
docker pull apache/skywalking-ui:latest

# 11800端口用于skywalking将应用的服务监控信息收集端口。
# 12800端口用于skywalking对UI提供查询接口。
docker run -itd \
    --name skywalking-oap-server \
    -e TZ=Asia/Shanghai \
    -p 12800:12800 \
    -p 11800:11800 \
    --link elasticsearch \
    -e SW_STORAGE=elasticsearch \
    -e SW_STORAGE_ES_CLUSTER_NODES=elasticsearch:9200 \
    apache/skywalking-oap-server:latest

docker run -itd \
    --name skywalking-ui \
    -e TZ=Asia/Shanghai \
    -p 8080:8080 \
    --link skywalking-oap-server \
    -e SW_OAP_ADDRESS=skywalking-oap-server:12800 \
    apache/skywalking-ui:latest
```

- 后台: <http://localhost:8080>
