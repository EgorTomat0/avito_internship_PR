# 🚀 Сервис назначения ревьюеров для Pull Request'ов

## 🎯 Quick start

### Установка и запуск

```bash
git clone https://github.com/EgorTomat0/PR_dz

cd avito_internship_PR
```
```bash
docker compose up -d
```
Или (если уже установлен task)
```bash
task up:docker-compose
```

После выполнения этих команд приложение будет доступно по адресу: http://localhost:8080

### Локальное использование
Для этого можно применять Taskfile
Гайд по установке: https://taskfile.dev/docs/installation#official-package-managers
````bash
#Для запуска проекта
task up:local
````
````bash
#Для запуска с помощью docker compose
task up:docker-compose
````