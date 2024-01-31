build:
	GOOS=linux GOARCH=amd64 go build -o docker/build/image/server main.go
	docker build -t ghcr.io/alma-security/go_app_service_a:8 .

	