FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY bin/notify /app/notify
COPY notify/etc /app/etc

CMD ["./notify", "-f", "etc/notify.yaml"]
