# Этап сборки
FROM golang:1.24.5 AS builder

# Установка необходимых C-зависимостей
RUN apt-get update && apt-get install -y \
    libgl1-mesa-dev \
    xorg-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . .

# Включаем CGO
ENV CGO_ENABLED=1

# Сборка бинарника
RUN go build -o main .

# Финальный образ — минимальный
FROM debian:bookworm-slim

# Установка только рантайм-библиотек (без компилятора)
RUN apt-get update && apt-get install -y \
    libgl1 \
    libx11-6 \
    libxcursor1 \
    libxrandr2 \
    libxinerama1 \
    libxi6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]
