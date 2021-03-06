FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -a --installsuffix cgo --ldflags="-s" -o main

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main .

EXPOSE 8080

ENTRYPOINT ["./main"]
