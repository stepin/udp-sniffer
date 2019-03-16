FROM alpine
RUN adduser -S -D -H -h /app appuser && mkdir /app
USER appuser
COPY udp-sniffer /app
WORKDIR /app
ENTRYPOINT ["/app/udp-sniffer"]
