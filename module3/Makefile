export tag=v1.0
build:
	echo "building module3-go enhttpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building module3-httpserver container"
	docker build -t lixixi/module3-httpserver:${tag} .

push: release
	echo "pushing lixixi/module3-httpserver"
	docker push lixixi/module3-httpserver:${tag}
