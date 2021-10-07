start:
	docker-compose -p runtime_metrics up -d

stop:
	docker-compose -p runtime_metrics down -v

restart: stop start

run:
	go run -v .

.PHONY: start stop restart run-tester
