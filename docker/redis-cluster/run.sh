docker exec -it cluster-redis-1 \
redis-cli --cluster create \
10.100.97.80:7000 10.100.97.80:7001 10.100.97.80:7002 10.100.97.80:7003 10.100.97.80:7004 10.100.97.80:7005 \
--cluster-replicas 1