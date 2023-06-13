client连接server
```bash
docker run -it --rm mysql:8.0 mysql -h 10.100.97.80 -uroot -p
docker run -it --rm mongo:6.0 mongosh -host 10.100.97.80 -u root -p root
docker run -it --rm redis:7.0 redis-cli -h 10.100.97.80
```

redis sentinel check
```bash
# redis-1 client
info replication
# redis-sentinel-1 client
sentinel master mymaster
sentinel replicas mymaster
sentinel sentinels mymaster
``` 

redis cluster check
```bash
# cluster-redis-1 client
cluster info
cluster nodes
```
