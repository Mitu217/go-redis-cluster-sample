# go-redis-cluster-sample

## Test

```bash
# Connect 127.0.0.1:7000
redis-cli -c -h localhost -p 7000

localhost:7000> set hoge fuga
localhost:7000> get hoge
"fuga"

# Connect 127.0.0.1:7001
redis-cli -c -h localhost -p 7001

localhost:7001> get hoge
-> Redirected to slot [1525] located at 127.0.0.1:7000
"fuga"
```
