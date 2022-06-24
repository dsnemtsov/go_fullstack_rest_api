# Go Fullstack App

## 1. Запуск docker
`docker-compose up`
## 2. Migrate
- Создать папку с файлами для миграций  
`migrate create -ext sql -dir ./schema -seq init` 
- Выполнить миграцию в соответствии с up файлом  
`migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/todo?sslmode=disable' up`  
При этом в БД создается таблица schema_migrations, отвечающая за версии миграции.  
- Откатить миграцию в сооветствии с down файлом  
`migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/todo?sslmode=disable' down`
   