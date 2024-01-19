FROM docker.io/library/golang:1.21.5 as builder
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0
RUN go build -o kvdbstore cmd/main/main.go

FROM gcr.io/distroless/static-debian12:nonroot
EXPOSE 8080
COPY --from=builder /app/kvdbstore /app/kvdbstore
ENTRYPOINT ["/app/kvdbstore"]
