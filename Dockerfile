FROM golang:1.21.7-alpine AS builder

WORKDIR /app

COPY --chown=app:app . .

RUN go build -ldflags="-s -w" -o bin/tgbot-watermark

FROM jrottenberg/ffmpeg:4.4-alpine

COPY --from=builder /app/bin /
COPY --from=builder /app/locale /locale

RUN chmod +x /tgbot-watermark

ENTRYPOINT ["/tgbot-watermark"]