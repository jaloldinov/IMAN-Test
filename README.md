#### Тестовое задание

## Open every service seperately otherwise you will get error
<br />

#### FIRST OF ALL YOU NEED TO SET UP AN ENV ACCORDING TO .ENV_EXAMPLE FILE
<br />

- create database
        
        CREATE DATABASE IF NOT EXISTS iman;
- then run with 

       make migration-up:
       make run


### After running all service you can see the result in the following link:

 http://localhost:8080/swagger/index.html#/ 


<br />
<br />
<br />

# Тестовое задание


### Необходимо создать маленькое микросервисное приложение, состоящее из 3 микросервисов. <br /> <br />
### Сервис №1 <br />
Задача сервиса собрать 50 страниц постов из открытого API - https://gorest.co.in/public/v1/posts
Собранные данные необходимо сохранить в ДБ (Любую на выбор).
Будет плюсом если данные будут собираться в несколько потоков.  <br /> <br />
### Сервис №2  <br /> 
Сервис должен реализовать логику GRUD для собранных ранее постов: <br />
    • Возможность получить несколько постов. <br />
    • Возможность получить конкретный пост <br />
    • Возможность удалить пост <br />
    • Возможность изменить пост  <br /> <br />
### Сервис №3  <br />
Сервис должен являться API Gateway (REST API) и предоставить методы для выполнения операций сервиса №1 и сервиса №2.  
    • Запуск процесса сбора данных и возможности проверки окончания процесса
    • Методы GRUD сервиса №2

Взаимодействие между сервисами должно осуществляться по gRPC.
Задание нужно расположить в любом Git репозитории и предоставить ссылку.

