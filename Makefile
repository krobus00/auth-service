launch_args=
test_args=-coverprofile cover.out && go tool cover -func cover.out
cover_args=-cover -coverprofile=cover.out `go list ./...` && go tool cover -html=cover.out

tidy:
	go mod tidy

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
  		--go-grpc_opt=paths=source_relative pb/auth/*.proto
	ls pb/auth/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

lint:
	golangci-lint run --disable-all -E errcheck -E misspell -E revive -E goimports

run-dev:
# make run-dev server, make run-dev worker
ifeq (server, $(filter server,$(MAKECMDGOALS)))
	$(eval launch_args=server $(launch_args))
else ifeq (worker, $(filter worker,$(MAKECMDGOALS)))
	$(eval launch_args=worker $(launch_args))
endif
	air --build.cmd "go build -o bin/auth-service main.go" --build.bin "./bin/auth-service $(launch_args)"
	
run:
# make run server, make run worker
ifeq (server, $(filter server,$(MAKECMDGOALS)))
	$(eval launch_args=server $(launch_args))
else ifeq (worker, $(filter worker,$(MAKECMDGOALS)))
	$(eval launch_args=worker $(launch_args))
endif
	./bin/auth-service $(launch_args)

build:
	# build binary file
	go build -ldflags "-s -w" -o bin/auth-service main.go
ifeq (, $(shell which upx))
	$(warning "upx not installed")
else
	# compress binary file if upx command exist
	upx -9 bin/auth-service
endif

test:
ifeq (, $(shell which richgo))
	go test ./... $(test_args)
else
	richgo test ./... $(test_args)
endif
	

cover: test
ifeq (, $(shell which richgo))
	go test $(cover_args)
else
	richgo test $(cover_args)
endif

%:
	@: