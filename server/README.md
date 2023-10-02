第一次搭建elk时，需要进入 kafka 容器创建 topic 
```shell
 docker exec -it kafka /bin/sh
 cd /opt/kafka/bin
 ./kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic zpshop-log
```

## 目录结构
```shell
app # 服务总目录
  cart # 购物车服务
  order # 订单服务
  pay # 支付服务
  product # 商品服务
  user # 用户服务
  system # 系统管理
common # 公用目录
  xjwt # jwt 鉴权
  result # 统一返回
  ......
data # docker 数据挂载目录
deploy # docker 容器配置相关
tpl # go-zero 自定义代码生成模板
```
app 文件夹下的每个目录都有如下文件结构：
```shell
cmd 
  api # 暴露 http 接口
  rpc # 提供 rpc 服务
model # orm 模型
```
端口映射
```shell
system:
  api: 2220
  rpc: 2221
order:
  api: 3330
  rpc: 3331
cart:
  api: 4440
  rpc: 4441
user:
  api: 5550
  rpc: 5551
product:
  api: 6660
  rpc: 6661  
```