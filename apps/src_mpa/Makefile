run:
	docker compose up -d
stop:
	docker compose down
	#docker rmi src-app
	#docker stop postgres
run_tests: 
	docker-compose up -d
	docker exec -it app /bin/bash -c "go test -v ./internal/service/tests/"
	docker exec -it app /bin/bash -c "go test -v ./internal/data_access/postgres/tests/"
	docker compose down
	#docker rmi src-app
stop_tests:
	docker compose down
	docker rmi src-app
lint:
	golangci-lint run ./... --config golangci.yaml -D gci -D govet -D testifylint -D forbidigo -D gofumpt
generate_mocks:
	./mockery --dir "./internal/service/repository/" --all --output "./internal/service/repository/mocks/"
run_service_tests_v:
	go test -v ./internal/service/tests/
run_service_tests:
	go test ./internal/service/tests/
run_da_tests_v:
	go test -v ./internal/data_access/postgres/tests/
run_da_tests:
	go test ./internal/data_access/postgres/tests/