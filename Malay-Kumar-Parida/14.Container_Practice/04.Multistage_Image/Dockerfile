FROM golang:1.17 as build-env

WORKDIR /

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go .

RUN go build -o /app


FROM gcr.io/distroless/base-debian11


COPY --from=build-env /app /

CMD [ "/app" ]