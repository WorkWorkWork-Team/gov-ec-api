run:
	ENV=dev go run main.go

unit-test:
	ginkgo -r --label-filter="unit"

integration-test:
	ginkgo -r --label-filter="integration"

start-dev-db-linux:
	docker start mysql-dev || docker run --name mysql-dev -d \
	-p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=P@ssw0rd \
	--restart unless-stopped \
	-v ${PWD}/database:/docker-entrypoint-initdb.d \
	mysql:latest

#not try yet
start-dev-db-window:
	docker start mysql-dev || docker run --name mysql-dev -d \
	-p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=P@ssw0rd \
	--restart unless-stopped \
	-v %cd%/database:/docker-entrypoint-initdb.d \
	mysql:latest

start-dev-php-admin:
	docker start phpmyadmin || docker run --name phpmyadmin -d --link mysql-dev:db -p 8080:80 phpmyadmin/phpmyadmin

mockgen:
	mockgen -destination=./test/mock_repository/mock_submitmp.go -source=./repository/submitmp.go -package=mock_repository
	mockgen -destination=./test/mock_repository/mock_population.go -source=./repository/population.go -package=mock_repository
	mockgen -destination=./test/mock_service/mock_submitmp.go -source=./service/submitmp.go -package=mock_service
	mockgen -destination=./test/mock_service/mock_population.go -source=./service/population.go -package=mock_service