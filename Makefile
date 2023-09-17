dev-up:
	docker-compose up --build -d

dev-down:
	docker-compose down

lint:
	goimports -w . && \
	gofumpt -w .