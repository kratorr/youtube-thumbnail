#  gRPC прокси-сервис для загрузки thumbnail’ов

## Сборка и запуск сервера
Склонировать проект
```bash
git clone git@github.com/kratorr/youtube-thumbnail
```

Перейти в директорию с проектом 
```bash
cd youtube-thumbnail 
```

Сборка
```bash
go build -o server cmd/main.go 
```
Запуск 
```bash
./server 
```
Конфигурационный файл должен лежать в одной директории с исполняем файлом
config.json
```
{
    "ip": "127.0.0.1",
    "port": 8085
}
```

## Сборка и запуск клиента

Сборка
```bash
go build -o client_th client/main.go
```
Запуск синхронный режим
```bash
./client_th --ip 127.0.0.1 --port 8081 https://www.youtube.com/watch\?v\=LxJLuW5aUDQ https://www.youtube.com/watch\?v\=7JLefviLqek
```
```bash
./client_th --ip 127.0.0.1 --port 8081 --async https://www.youtube.com/watch\?v\=LxJLuW5aUDQ https://www.youtube.com/watch\?v\=7JLefviLqek
```