FROM golang as builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go build -o wake-on-lan

FROM scratch

COPY --from=builder /app/wake-on-lan .
COPY index.html .

CMD ["./wake-on-lan"]
