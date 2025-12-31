SWAGGO_BIN=swag
SWAGGO_CMD=$(SWAGGO_BIN) init

.PHONY: swaggo
swaggo:
	@echo "Generating Swagger docs with swaggo..."
	$(SWAGGO_CMD)

.PHONY: install-swaggo
install-swaggo:
	@echo "Installing swaggo..."
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: dev
dev: 
	go run main.go server rest

.PHONY: migrateup
migrateup: 
	go run main.go migrate up

.PHONY: migratedown
migratedown: 
	go run main.go migrate down

.PHONY: mock
mock:
	@echo "Generating mocks for all files in repository directory..."
	@for file in internal/app/repository/*.go; do \
		mockgen -source=$$file -destination=internal/app/mocks/mock_$$(basename $$file .go).go -package=mocks; \
	done

.PHONY: test
test:
	godotenv -f ./.env.test go test ./internal/... -cover