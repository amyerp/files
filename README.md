# files Microservice

Full API Documentation in docs/ folder

## Build Microservoce

```
docker build --no-cache -t files:latest -f Dockerfile .
```
or
```
docker build -t files:latest -f Dockerfile .
```


## Run Microservice in Docker (in case if it in the same area with API Gateway)

```
docker run --name files \
--restart=always \
-v $(pwd)/config:/var/gufo/config \
-v $(pwd)/lang:/var/gufo/lang \
-v $(pwd)/templates:/var/gufo/templates \
-v $(pwd)/logs:/var/gufo/log \
-v $(pwd)/files:/var/gufo/files \
--network="gufo" \
-d files:latest
```

Before run microservice need to add in API Gateway config next lines

```
[microservices]
[microservices.files]
type = 'server'
host = 'files'
port = '5300'
entrypointversion = '1.0.0'
cron = 'false'
```
