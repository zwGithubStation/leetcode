
Features
--------
- Create a [Fast](#benchmarks) custom Redis compatible server in Go
- Simple interface. One function `ListenAndServe` and two types `Conn` & `Command`
- Support for pipelining and telnet commands
- Works with Redis clients such as [redigo](https://github.com/garyburd/redigo), [redis-py](https://github.com/andymccurdy/redis-py), [node_redis](https://github.com/NodeRedis/node_redis), and [jedis](https://github.com/xetorthio/jedis)
- [TLS Support](#tls-example)
- Compatible pub/sub support
- Multithreaded
