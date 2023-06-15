lang:
	easyi18n generate --pkg=locales ./locales ./locales/locales.go

api:
	goctl api go -api notify/notify.api -dir notify -style gozero --home template

run:
	go run notify/notify.go -f notify/etc/notify.yaml -env notify/etc/.env

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/notify_service notify/notify.go