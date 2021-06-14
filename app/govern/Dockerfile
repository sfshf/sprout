FROM golang:1.16.3-alpine3.12 AS builder
LABEL name="govern-builder" version="1.0"
WORKDIR /sprout
COPY . .
RUN set -eo pipefail && \
    go env -w GOPROXY="https://goproxy.io,direct" && \
    go mod tidy && \
    go build -o govern ./app/govern

FROM alpine:3.12 AS binary
LABEL name="govern-binary" version="1.0"
MAINTAINER sfshf
ENV CONFIG_FILE="/app/config.toml"
ENV CASBIN_MODEL_FILE="/app/casbin_rbac.model"
ENV MONGO_URI="mongodb://govern-mongo:27017/sprout?directConnection=true"
ENV MONGO_DATABASE="sprout"
ENV REDIS_ADDRESS="govern-redis:6379"
ENV HTTP_HOST="govern"
ENV HTTP_CERT_FILE=""
ENV HTTP_CERT_KEY_FILE=""
EXPOSE 8000/tcp
EXPOSE 8010/tcp
#RUN echo "http://dl-4.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && apk --update add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /sprout/govern /sprout/app/govern/config/config.toml /sprout/app/govern/config/casbin_rbac.model ./
CMD ["./govern"]

