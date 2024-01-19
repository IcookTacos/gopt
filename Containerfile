FROM docker.io/library/golang:1.21.5 as builder
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=1
RUN go build -o kvdbstore cmd/main/main.go

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/kvdbstore /usr/bin/kvdbstore
EXPOSE 8080
ENTRYPOINT ["/usr/bin/kvdbstore"]
