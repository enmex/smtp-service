FROM golang:1.18-alpine as builder

WORKDIR /srv
ENV CONFIG_PATH=${CONFIG_PATH}
COPY ./go.mod ./ ./go.sum ./

RUN go mod download \
    && go mod verify

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a \
    -ldflags '-extldflags "-static"' \
    -o smtp ./cmd/main.go
RUN echo $CONFIG_PATH > ./config.yml

WORKDIR /srv

FROM scratch
COPY --from=builder srv/smtp /usr/local/bin/smtp
COPY --from=builder srv/config.yml /usr/local/bin/config.yml
ENTRYPOINT ["/usr/local/bin/smtp"]