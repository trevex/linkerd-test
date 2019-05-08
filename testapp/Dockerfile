FROM golang:1.12.4 as builder
WORKDIR /testapp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o testapp main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates bash
COPY --from=builder /testapp/testapp /usr/bin/testapp

CMD ["/usr/bin/testapp"]

