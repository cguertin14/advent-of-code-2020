test:
	go test -v ./... --bench --benchmem --cover --covermode=count --coverprofile=coverage.out