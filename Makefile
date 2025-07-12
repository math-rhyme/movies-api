a:
	go run ./cmd/api

b:
	# server in background + saving PID
	go run ./cmd/api & echo $$! > server.pid

stest: b
	sleep 2
	curl -i localhost:4000/v1/healthcheck
	-kill `cat server.pid`
	-rm server.pid
