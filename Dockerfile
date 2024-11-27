FROM golang:1.22.6-alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -tags musl -o main

FROM gcr.io/distroless/base-debian11:nonroot

WORKDIR /app

COPY --from=build /app/main .

USER nonroot:nonroot

CMD ["/app/main"]
