FROM golang as builder

WORKDIR /app

COPY . .

RUN go build -o wake-on-lan

FROM scratch

COPY --from=builder /app/wake-on-lan .

CMD ["./wake-on-lan"]
