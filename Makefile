build:
	docker build . -t pursuit-mortalkin-dock

run:
	docker run --net pursuit_network -v ~/pursuit/mortalkin/resource/snapshot/:/resource/snapshot -p 5004:5004 pursuit-mortalkin-dock

pretty:
	go fmt `go list ./...`

test:
	go test `go list ./... | grep -v cmd | grep -v vendor`
