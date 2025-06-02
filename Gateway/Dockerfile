# Этап сборки
FROM golang:1.24 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum, чтобы сначала установить зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем бинарный файл
RUN go build -o main ./cmd/main.go

# Этап минимального запуска
FROM debian:bookworm-slim

# Создаем непривилегированного пользователя
RUN useradd -m appuser

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/main /main
COPY keys /app/keys

# Меняем владельца файлов, чтобы appuser имел доступ
RUN chown -R appuser:appuser /app/keys

# Меняем пользователя
USER appuser
EXPOSE 8080
# Устанавливаем команду по умолчанию
ENTRYPOINT ["/main"]