run:
	docker-compose up --build fibonacci
getSequence:
	docker exec -it fibonacci_fibonacci_1 /client $(start) $(end)
test:
	go test -v ./pkg/repository