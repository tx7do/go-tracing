# OpenTelemetry

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
