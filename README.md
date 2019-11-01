# timer

```bash
docker build -t timer:v1.0.0 .
```

```bash
docker run \
-d \
--name timer \
-v /etc/localtime:/etc/localtime:ro \
--restart=always \
timer:v1.0.0
```