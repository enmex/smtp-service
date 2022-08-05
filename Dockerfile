FROM golang:1.18-alpine as builder

WORKDIR /srv

ARG config_path

COPY ./go.mod ./ ./go.sum ./
COPY $config_path ./

RUN go mod download \
    && go mod verify
RUN cp $config_path config.yml

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a \
    -ldflags '-extldflags "-static"' \
    -o smtp ./cmd/main.go

FROM scratch
COPY --from=builder srv/smtp /usr/local/bin/smtp
COPY --from=builder srv/config.yml /usr/local/bin/config.yml
ENTRYPOINT ["/usr/local/bin/smtp"]