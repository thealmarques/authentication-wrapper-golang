# start server in background and wait for 1 sec
redis-server --daemonize yes && sleep 1 

redis-cli < /scripts/seed.redis 

redis-cli save 

redis-cli shutdown 

redis-server 