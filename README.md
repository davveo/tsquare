### tsquare
#### 项目启动顺序
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
