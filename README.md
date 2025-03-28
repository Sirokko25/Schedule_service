## Стек
Golang, PostgreSQL, Gin, BunORM, Docker

## Установка

### Склонируйте репозиторий

```
git clone https://github.com/Sirokko25/Schedule_service.git
cd Sсhedule_service
go mod tidy
```

## Настройка
### Параметры окружения
Путь: ./env/local.sh
Можно установить свои параметры окружения

## Инициализация базы данных

Приложение автоматически создаст файл базы данных и необходимые таблицы при первом запуске, если файл базы данных не существует.

## Запуск приложения
### Запуск сервера и базы данных

Перейти в корневую директорию и выполнить в терманле команду task build ENV=local.

Сервер будет доступен по адресу http://0.0.0.0:8020
База данных будет доступна по адресу http://0.0.0.0:5432

### Тестирование
Для тестирования имеется Postman коллекция.

### Структура проекта
- сmd/: Директория содержит главный файл приложения - точка входа сервера.
- server/: Директория в которой запускается сервер, определяются роутеры, запускается подключение к бд.
- docker/: Директория с конфигурациями Docker контейнеров
    - local-docker-compose.yml: Файл конфигурации для управления Docker-контейнерами
- Taskfile.yaml: Скрипты командной оболочки для сборки Docker контейнеров 
- Dockerfile: Набор инструкций для автоматизированной сборки образа Docker
- postman/: Директория с postman коллекцией тестовых запросов
- internal/: Директория с логикой проекта:
    - errorlist/: Пакет c набором ошибок.
    - handlers/: Пакет с обработчиками API запросов.
        - AppendShedule.go: Ручка для добавления расписания в бд
        - GetShedule.go: Ручка для получения расписания по user_id и shedule_id
        - GetShedules.go: Ручка получения расписаний для пользователя по user_id
        - GetNextTakings.go: Ручка для получения информации о времени приема лекарств для пользователя по user_id
    - storage/: Пакет для инициализации и работы с базой данных
        - base.go: Хранит структуры для создания таблиц и интерфейс для взаимодействия с бд
        - Сonnection.go: Создает таблицы если они не созданы, и подключение к бд
        - AppendShedule.go: Метод для добавления расписания в бд
        - FindShedule.go: Метод для поиска расписания по user_id и shedule_id
        - FindShedules.go: Метод для поиска расписаний для пользователя по user_id
        - NextTakings.go: Метод для поиска информации о времени приема лекарств для пользователя по user_id
        - PingDB.go: Метод для проверки подключения к бд.
        - Actualize.go: Метод для проверки актуальности хранимых данных, и перемещению устаревших в таблицу истории
    -  models/: Пакет с необходимыми для работы структурами
    -  helpers/: Пакет со вспомогательными методами
        - CalculateIntervals.go: Для расчёта интервалов приема лекарств
        - CreateIntervals.go: Для создания интервалов приема лекарств
        - EndDateCalculate.go: Для расчёта конечной даты приема лекарств
        - ValidateShedule.go: Для валидации поступающих расписаний
        - CheckChart.go: Для проверки времени ближайшего приёма лекарств
        - CreateResponceString.go: Создание строки ответа поиска расписаний для пользователя по user_id
