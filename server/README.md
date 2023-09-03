第一次搭建elk时，需要进入 kafka 容器创建 topic 
```shell
 docker exec -it kafka /bin/sh
 cd /opt/kafka/bin
 ./kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic zpshop-log
```