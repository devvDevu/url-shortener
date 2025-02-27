![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Go version](https://img.shields.io/badge/Golang-1.23.6-blue)


# URL Shortener Service

🚀 Сервис для сокращения URL-адресов с возможностью отслеживания и управления сокращенными ссылками

## Особенности
- Генерация коротких кодов фиксированной длины
- Валидация входных URL (регулярные выражения + net/url)
- Многослойная архитектура (Handler -> UseCase -> Service)
- Подробное логирование операций
- Полная документация API через автоматическую генерацию роутов
- Поддержка интеграционных и unit-тестов


## Технологический стек
- **Язык**: Go 1.21+
- **Маршрутизация**: Gorilla Mux
- **Валидация**: go-playground/validator
- **Логирование**: logrus
- **Тестирование**: testify + gomock
- **Конфигурация**: встроенные типы с валидацией

## Быстрый старт

### Установка
```bash
go mod download
go install go.uber.org/mock/gomock@latest
go install go.uber.org/mock/mockgen@v0.3.0
```

### Запуск
```bash
go run cmd/main.go
```

### Примеры запросов

Создание короткой ссылки:
```bash
curl -X POST http://localhost:8080/url \
  -H "Content-Type: application/json" \
  -d '{"original_url":"https://www.google.com"}'
```

Получение оригинального URL:
```bash
curl -X GET http://localhost:8080/url \
  -H "Content-Type: application/json" \
  -d '{"code":"uO2_-PFG"}'
```

## Тестирование
Unit-тесты:
```bash
go test -v ./internal/...
```

Интеграционные тесты:
```bash
go test -v ./tests/integration/...
```

## Архитектура

### Слои приложения
1. **Handler** - обработка HTTP запросов
   - Валидация входных данных
   - Преобразование DTO
   - Логирование ошибок

2. **UseCase** - бизнес-логика
   - Генерация коротких кодов
   - Координация работы сервисов
   - Обработка ошибок

3. **Service** - работа с данными
   - Взаимодействие с хранилищем
   - Кеширование
   - Транзакции

### Модели данных
- **Url** - основная сущность:
  ```go
  type Url struct {
      Id       UrlId
      Original UrlOriginal
      Code     UrlCode
  }
  ```

## Конфигурация
Параметры настройки через типы:
```go
type UrlCode string // формат: 8 символов
type UrlOriginal string // валидный URL
```

## Разработка
### Генерация моков
```bash
mockgen -source=internal/usecase/url_usecase/usecase.go -destination=mocks/mock_service.go -package=mocks
```

