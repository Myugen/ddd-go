# syntax=docker/dockerfile:1
ARG GOLANG_VERSION=1.17-alpine

# Builder image
FROM golang:${GOLANG_VERSION} as builder

WORKDIR /ddd-go

COPY go.* ./

ENV CGO_ENABLED=0 \
    GO111MODULE="on" \
    GOOS=linux

RUN echo 'Downloading go.mod dependencies' \
    && go mod download

COPY . .
RUN go build -o /ddd-go/build/app ./cmd/main.go


# Generate clean, final image for end users
FROM scratch as app

COPY --from=builder /ddd-go/build/app /ddd-go/app

ENTRYPOINT [ "/ddd-go/app" ]
