# timer

```bash
docker build -t timer:v1.0.0 .
```

```bash
docker run \
-d \
--name timer \
-v timer-storage:/data \
-v /etc/localtime:/etc/localtime:ro \
--restart=always \
timer:v1.0.0
```