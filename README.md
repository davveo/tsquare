## tsquare
### 说明
这个项目暂时停止开发, 目前先更新完单体版的服务, 项目地址为:
https://github.com/davveo/singleTsquare
单体版后端预计8月底更新完毕.

等到单体版更新完毕后, 会继续更新服务版，欢迎大家star关注.



### 项目架构图
![项目架构图](https://github.com/zbrechave/tsquare/blob/master/docs/项目架构图.jpg)
### 项目启动
```bash
> cd bin
> sh build.sh
```
### 项目规划
 - 四周左右时间完善底层基础服务: 如: 短信服务, 用户服务, 认证服务, UUID服务, 问题/回答服务...
 - 两周左右时间完成api服务: 如: user-api, answer-api...
 - 一周左右时间完成上述调试
 
### 项目技术点
 - 分布式session解决方案
 - 分布式UUID解决方案
 - 分布书搜索解决方案
 - 分布式缓存解决方案
 - 分布式配置管理
 - 分布式链路追踪
 - 持续集成/容器化技术
 - 消息队列(rabbitmq/kafka)等
 
 ### 服务涉及技术点总结
 #### 配置服务(conf-srv)
 #### 用户服务(user-srv)
 #### 短信服务(sms-srv)
 #### 认证服务(auth-srv)
 #### 敏感词服务(senword-srv)
 #### 问题服务(answer-srv)
 #### 回答服务(question-srv)
 
### 说明
#### 分布式session解决方案
  > 解决多服务之间, 用户身份识别。

#### 分布式UUID解决方案
  > 解决用户ID生成, 问题回答ID生成, 日志ID生成等。

#### 分布书搜索解决方案
  > 解决问题回答搜索, 招聘信息搜索等

#### 分布式缓存解决方案

#### 分布式配置管理
 > 管理服务配置

#### 分布式链路追踪
