dev-up:
	docker compose up -d

dev-down:
	docker compose down

purge:
	docker compose exec kafka kafka-topics --delete --topic foo-bar-topic --bootstrap-server localhost:9092

tail:
	docker compose exec kafka kafka-console-consumer --topic foo-bar-topic --from-beginning --bootstrap-server localhost:9092

proto:
	protoc --go_out=. proto/*.proto

.PHONY: dev-up dev-down purge tail proto
