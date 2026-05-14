APP=be-summer-store
PORT=8080

build:
	docker build -t $(APP) .

run:
	docker run -d -p $(PORT):$(PORT) -v $(PWD)/.env:/app/.env $(APP)

dev:
	docker run -d -p $(PORT):$(PORT) -v $(PWD)/.env:/app/.env --env GIN_MODE=debug $(APP)

stop:
	docker stop $$(docker ps -q --filter ancestor=$(APP))

restart: stop run

logs:
	docker logs -f $$(docker ps -q --filter ancestor=$(APP))

re: build run

clean:
	docker rmi $(APP)

.PHONY: build run dev stop restart logs re clean