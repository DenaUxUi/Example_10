# этап 1: билдим бинарник
from golang:1.21 as builder

workdir /app

copy . .

run go mod init myapp && go mod tidy

run go build -o server

# этап 2: минимальный контейнер
from debian:bookworm-slim

# переменные окружения (если хочешь задать порт)
env port=8080

# копируем бинарник из билдера
copy --from=builder /app/server /server

# указываем порт (документально, не обязательно)
expose 8080

# запуск сервера
entrypoint ["/server"]
