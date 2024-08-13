from golang:1.22.6 as build
label org.opencontainers.image.authors="Vladislav Gardner <vladislavgardner@gmail.com>"

workdir /app

copy go.mod go.sum .env ./
run mkdir -p cmd internal pkg
copy cmd/ ./cmd
copy internal ./internal
copy pkg ./pkg

run go mod download
run go mod verify
run CGO_ENABLED=0 GOOS=linux go build -o ./app ./cmd/service/main.go

run rm -rf go.mod go.sum cmd internal pkg

from busybox:1.36.1

workdir /service

copy --from=build ./app/app .
copy --from=build ./app/.env .

# TODO: fix nobody
user 1000:1000

cmd ["./app"]
