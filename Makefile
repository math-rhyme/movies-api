a:
	@go run ./cmd/api

b:
	# server in background + saving PID
	@go run ./cmd/api & echo $$! > server.pid

stest: b
	@sleep 2
	@echo "\n\e[36m****************\e[0m"
	curl -i localhost:4000/v1/healthcheck
	@echo "\n\e[36m****************\e[0m"
	curl -X POST localhost:4000/v1/movies
	@echo "\n\e[36m****************\e[0m"
	curl localhost:4000/v1/movies/123
	@echo "\n\e[36m****************\e[0m"
	curl -i -X POST localhost:4000/v1/healthcheck
	@echo "\n\e[36m****************\e[0m"
	curl -i -X OPTIONS localhost:4000/v1/healthcheck
	@echo "\n\e[36m****************\e[0m"
	curl -i localhost:4000/v1/movies/abc
	@echo "\n\e[36m****************\e[0m"
	@-kill `cat server.pid`
	@-rm server.pid
