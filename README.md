# Демонстрационный сервис с Kafka, PostgreSQL, кешем

## Запуск kafka и postgres
```bash
docker compose up -d
```

## Запуск миграций
```bash
go run ./cmd/migrate/main.go
```

## Запуск приложения
```bash
go run ./cmd/app/main.go
```

## Отправка сообщения (order.json) в kafka
```bash
go run ./producer/producer.go
```
