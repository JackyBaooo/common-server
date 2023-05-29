client连接server
```bash
docker run -it --network mysql_default --rm mysql:8.0 mysql -h mysql-db-1 -uroot -p
docker run -it --network mongo_default --rm mongo:6.0 mongosh -host mongo-db-1 -u root -p root
docker run -it --network redis_default --rm redis:7.0 redis-cli -h redis-db-1
docker run -it --network redis-sentinel_default --rm redis:7.0 redis-cli -h redis-1 --pass 123456
docker run -it --network redis-cluster_default --rm redis:7.0 redis-cli -h cluster-redis-1 -c
```

redis sentinel build
```bash
docker compose -f redis.yml up -d
docker inspect redis-1 | grep IPAddress
sh copy.sh $IPAddress
docker compose -f sentinel.yml up -d
```

redis sentinel check
```bash
# redis-1 client
info replication
# redis-sentinel-1 client
sentinel master mymaster
sentinel replication mymaster
sentinel sentinels mymaster
``` 

redis cluster check
```bash
# cluster-redis-1 client
cluster info
cluster nodes
```
