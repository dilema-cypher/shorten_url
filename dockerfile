# Estágio 1: Build
# Use uma versão específica ao invés de 'latest' para garantir previsibilidade
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependência primeiro para aproveitar o cache do Docker
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest  

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

EXPOSE 8282

CMD ["./main"]