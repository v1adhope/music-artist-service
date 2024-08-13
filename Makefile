# Docker image version
export version="v0.0.1"
# Migration number
export n="1"

.PHONY: tests
tests:
	go test --race ./...

.PHONY: test-by-name
test-by-name:
	go test --race -run ${name} ./...

.PHONY: generate-protos
generate-protos:
	./scripts/generate_protos.sh

.PHONY: dockerfile
dockerfile:
	docker build -t v1adhope/music-artist-service:${version} .

.PHONY: dockerfile-debug
dockerfile-debug:
	docker build --no-cache --progress=plain -t v1adhope/music-artist-service:${version} .

.PHONY: compose-up
compose-up:
	docker-compose up -d

.PHONY: compose-up-dev
compose-up-dev: dockerfile compose-up

.PHONY: compose-down
compose-down:
	docker-compose down

.PHONY: migrate-up
migrate-up:
	docker run -v $(shell pwd)/db/migrations:/migrations\
		--network music-artist-service_default\
		migrate/migrate\
		-path migrations\
		-database "postgres://rat:secret@postgres:5432/music_artist_service?sslmode=disable"\
		up

.PHONY: migrate-down
migrate-down:
	echo "y" | \
	docker run -v $(shell pwd)/db/migrations:/migrations\
		--network music-artist-service_default\
		migrate/migrate\
		-path migrations\
		-database "postgres://rat:secret@postgres:5432/music_artist_service?sslmode=disable"\
		down ${n}

.PHONY: migrate-force
migrate-force:
	docker run -v $(shell pwd)/db/migrations:/migrations\
		--network music-artist-service_default\
		migrate/migrate\
		-path migrations\
		-database "postgres://rat:secret@postgres:5432/music_artist_service?sslmode=disable"\
		force ${n}
