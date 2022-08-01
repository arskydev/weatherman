ARG APP_NAME="weatherman"

FROM golang:latest as builder
WORKDIR /go/src
ENV CGO_ENABLED=0 
ENV GOOS="linux" 
COPY . .
RUN go get -d -v ./... && \
    go build ./cmd/api && \
    mv ./api ./weatherman

FROM alpine:latest
WORKDIR /usr/home
ENV WEATHER_API_KEY="WEATHER_KEY"
ENV IPGEO_API_KEY="IPGEO_KEY"
ENV DB_PASS="PG_PASS"
COPY --from=builder /go/src/weatherman ./weatherman
COPY --from=builder /go/src/config ./config
ENTRYPOINT ["./weatherman"]
