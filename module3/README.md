
# 模块三作业:

## 构建httpserver本地镜像并推送到docker官方镜像仓库:
```
$ make push
```

## 本地启动httpserver镜像，并将端口映射到主机8080端口:
```
$ docker run -p 8080:80 lixixi/module3-httpserver:v1.0
```

### 在主机访问httpserver
`http://localhost:8080/readHeaders`

`http://localhost:8080/healthz`

## 查看容器IP
（由于目前docker安装在Windows系统上，无法运行nsenter命令，用了以下命令进入容器shell查看IP）
```
# 在本地后台运行httpserver
$ docker run -d -p 8080:80 lixixi/module3-httpserver:v1.0

# 查看httpserver容器ID
$ docker ps -a

# 进入容器shell
$ docker exec -it 96f8bd247bb5 sh

# 在容器shell中运行hostname命令查看IP
$ hostname -I 
```
