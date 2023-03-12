FROM golang:1.18
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -tags prod -o /app/cleango main.go
RUN groupadd -r appgroup && useradd --no-log-init -r -g appgroup appuser
USER appuser
CMD ["/app/cleango"]
