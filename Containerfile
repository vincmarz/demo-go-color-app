FROM registry.access.redhat.com/ubi9/go-toolset:latest AS builder
WORKDIR /opt/app-root/src
COPY go.mod .
COPY main.go .
RUN go build -o app main.go

FROM registry.access.redhat.com/ubi9/ubi-minimal
WORKDIR /opt/app
COPY --from=builder /opt/app-root/src/app /opt/app/app
EXPOSE 8080
CMD ["/opt/app/app"]
