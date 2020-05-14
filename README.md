## tsquare
### 项目架构图
![项目架构图](https://github.com/zbrechave/tsquare/raw/master/docs/项目架构图.png)
### 项目启动顺序
  1. 启动服务
  ```bash
    # 启动consul
    docker run --name consul -d -p 8500:8500/tcp consul:latest agent -server -ui -bootstrap-expect=1 -client=0.0.0.0
    
    # 启动hystrix监控
    docker run --name hystrix-dashboard -d -p 8081:9002 mlabouardy/hystrix-dashboard:latest
    # 访问地址:  http://localhost:8081/hystrix.stream
    # 输入ip访问监控: http://{ip}:81/hystrix.stream {ip}为本机ip
    
    # 启动mysql
    docker run --name mysql -e MYSQL_ROOT_PASSWORD=123 -d -p 3306:3306 mysql
    
    # 启动jaeger链路追踪
    docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:1.6
    
    # 启动prometheus监控
    docker run --name prometheus -d -p 0.0.0.0:9090:9090 -v ~/tmp/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus     
  ```
  ```yml
  # prometheus.yml
  global:
    scrape_interval: 15s
    scrape_timeout: 10s
    evaluation_interval: 15s
  alerting:
    alertmanagers:
    - static_configs:
      - targets: []
      scheme: http
      timeout: 10s
  scrape_configs:
  - job_name: APIGW
    honor_timestamps: true
    scrape_interval: 15s
    scrape_timeout: 10s
    metrics_path: /metrics
    scheme: http
    static_configs:
    - targets:
      - 10.104.34.106:8080   #10.104.34.106为本机ip， 本机127.0.0.1在容器中无法访问到

   ```
  1. 启动api-gateway
  ```bash
    cd api
    make run
  ```
  2. 启动各个srv
  ```bash
    // 启动认证服务
    cd srv/auth-srv
    make run
    
    // 启动用户服务
    cd srv/user-srv
    make run
    
    ...
  ```
    
### auth-srv

### user-srv

### order-srv

### payment-srv

### inventory-srv

### settlement-srv
