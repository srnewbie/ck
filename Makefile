SERVER_PKG:=ck
SERVER_BIN=ck_server
APP_PKG:=ck/app
APP_BIN=ck_app

all:
	go build -o bin/$(SERVER_BIN) $(SERVER_PKG)
	go build -o bin/$(APP_BIN) $(APP_PKG)

test:
	mkdir -p .gen/mocks
	mockery -dir=models/queue -name=Queue -case=underscore -output=.gen/mocks/
	mockery -dir=models/pq -name=PQ -case=underscore -output=.gen/mocks/
	mockery -dir=models/cron -name=Cron -case=underscore -output=.gen/mocks/
	go test ./... -cover -coverprofile .gen/cover.out
	go tool cover -html=.gen/cover.out -o .gen/cover.html
	open .gen/cover.html
