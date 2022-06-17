1. Запуск Docker:  
   docker-compose up 
2. migrate:  
   2.1 Создать папку с файлами для миграций     
       migrate create -ext sql -dir ./schema -seq init  
   2.2 Выполнить миграцию в соответствии с up файлом  
       migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/todo?sslmode=disable' up  
   При этом в БД создается таблица schema_migrations, отвечающая за версии миграции.  
   2.3 Откатить миграцию в сооветствии с down файлом  
       migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/todo?sslmode=disable' down
   