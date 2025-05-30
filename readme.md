# 📘 Общая информация

**Название проекта:** Citatnik

**Язык:** Go

**Основной пакет:** /cmd/citatnik/main.go

## 🎯 Требования

Go версии 1.20 или выше

Операционная система: Windows 10/Linux

Права доступа: права на выполнение команд в терминале

## 🔧 Установка

### 🪟 Windows

```shell
git clone https://github.com/zhenyanesterkova/citatnik

go build -o ./cmd/citatnik/citatnik.exe ./cmd/citatnik/main.go
```

### 🐧 Linux

```bash
git clone https://github.com/zhenyanesterkova/citatnik

go build -o ./cmd/citatnik/citatnik ./cmd/citatnik/main.go
```

## 🎮 Запуск

### 🪟 Windows

```shell
./cmd/citatnik/citatnik.exe
```

#### 🛠 Запуск с использованием makefile

```shell
make run-win
```

### 🐧 Linux

```bash
chmod +x ./cmd/citatnik/citatnik

./cmd/citatnik/citatnik
```

#### 🛠 Запуск с использованием makefile

```bash
make run-linux
```

## 📝 Лицензирование

Проект распространяется без лицензии.

## 📋 Дополнительная информация

Порт по умолчанию: 8080

Уровень логирования: info

База данных: встроенная в память (для тестирования)

## 🔍 API методы

```
GET /ping

GET /quotes - получение списка цитат

GET /quotes/random - получение случайной цитаты

GET /quotes?author=Confucius - получение цитат автора

POST /quotes - добавление новой

DELETE /quotes/{id} - удаление цитаты по ID
```

### 1. Проверка работоспособности сервиса

```
   Метод: GET
   URL: /ping
```

Проверка доступности хранилища данных.

Ответы:

```
200 OK - хранилище доступно
500 Internal Server Error - ошибка при проверке хранилища
```

### 2. Получение списка цитат

```
   Метод: GET
   URL: /quotes
```

Получение списка всех цитат или цитат определенного автора.

Параметры запроса:

author (опционально) - имя автора для фильтрации

Ответы:

```
200 OK - список цитат в формате JSON
500 Internal Server Error - ошибка при получении данных
```

### 3. Получение случайной цитаты

```
   Метод: GET
   URL: /quotes/random
```

Получение случайной цитаты из базы данных.

Ответы:

```
200 OK - случайная цитата в формате JSON
500 Internal Server Error - ошибка при получении данных
```

### 4. Добавление новой цитаты

```
   Метод: POST
   URL: /quotes
```

Добавление новой цитаты в базу данных.

Тело запроса:

```
{
"quote": "текст цитаты",
"author": "автор цитаты"
}
```

Ответы:

```
200 OK - цитата успешно добавлена
500 Internal Server Error - ошибка при добавлении цитаты
```

### 5. Удаление цитаты

```
   Метод: DELETE
   URL: /quotes/{id}
```

Удаление цитаты по идентификатору.

Параметры URL:

id - идентификатор удаляемой цитаты

Ответы:

```
200 OK - цитата успешно удалена
400 Bad Request - неверный формат ID или цитата не найдена
500 Internal Server Error - ошибка при удалении цитаты
```
