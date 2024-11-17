# Сервис для сбора метрик проекта из Jira

## Сборка приложения для запуска на ВМ (RHEL/SberLinux)
```
CGO_ENABLED=0 go build -o performance-dashboard -ldflags="-s -w" pkg/main.go
```