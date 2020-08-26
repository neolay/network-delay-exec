FROM python:3.7

ADD network-delay-exec /network-delay-exec
EXPOSE 2332
ENTRYPOINT ["./network-delay-exec"]
