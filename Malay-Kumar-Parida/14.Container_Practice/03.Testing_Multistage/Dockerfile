FROM golang:1.17 as build-env

WORKDIR /

COPY go.mod ./

COPY *.go .

RUN go mod download

RUN go build -o /app

FROM gcr.io/distroless/static

COPY --from=build-env /app /
CMD ["/app"]
