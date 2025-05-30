# Quoter App

### Описание
Quoter App - простое CRUD приложение для создания, просмотра и удаления цитат

### Установка
Для того чтобы запустить это приложение локально на вашем устройстве вам нужно:
1. Скопировать данный репозиторий к себе на устройство ```git clone https://github.com/JulyInSummer/quoter.git```
2. У вас должен быть установлен Docker & Docker-Compose так как в нем нужно будет развернуть базу данных ```make run-local-docker-up```, а если у вас ее установлен ```make``` то, нужно выполнить ```docker-compose -f docker-compose.local.yaml up --build -d```
> Настоятельно рекомендуется установить утилиту ```make``` так как это облегчить выполнение разных команд

> Перед 3 шагом нужно установить утилиту ```migrate```.

для Ubuntu
```bash
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
    sudo apt-get update
    sudo apt-get install migrate
```

для MacOS
```bash
    brew install golang-migrate
```
3. Далее нужно накатить миграции для базы данных: ```make migrate``` или ```migrate create -ext sql -dir ${CURRENT_DIR}/migrations -seq -digits 8 {name_of_migration}```
4. Далее для запуска приложения выполнить команду ```make run``` или ```go run cmd/main.go```

### Тесты
Для запуска тестов и проверки работоспособности приложения выполните команду

```bash
    make test
```
или
```bash
    go run -p 1 ./...
```


### Описание маршрутов
Создание цитаты
```bash
    curl --request POST \
  --url http://localhost:8080/quotes \
  --header 'Content-Type: application/json' \
  --data '{
	"author": "Mr. White",
	"quote": "Im not in danger. I'\''M THE DANGER!"
}'
```

Получение всех цитат
```bash
    curl --request GET \
  --url 'http://localhost:8080/quotes'
```
или с фильтром по автору
```bash
    curl --request GET \
  --url 'http://localhost:8080/quotes?author=Mr.%20White'
```

Получение случайной цитаты
```bash
    curl --request GET \
  --url http://localhost:8080/quotes/random 
```

Удаление цитаты
```bash
    curl --request DELETE \
  --url http://localhost:8080/quotes/3 
```