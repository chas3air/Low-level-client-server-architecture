ARG GO_VERSION=latest
FROM golang:${GO_VERSION} AS build

WORKDIR /src

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files into the container
COPY . .

ARG TARGETARCH
# Build the executable after all files have been copied
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build -o /src/cli ./cmd/app 

FROM alpine:latest AS final

RUN apk --no-cache add ca-certificates tzdata

# Create app directory and copy source files there
RUN mkdir /app
COPY --from=build /src /app

# Move the executable to the app directory
RUN mv /app/cli /cli

RUN adduser -D appuser
USER appuser

WORKDIR /

EXPOSE 12016

ENTRYPOINT ["tail", "-f", "/dev/null"]
# ENTRYPOINT ["./cli", "--config=./app/config/local.yaml"]
