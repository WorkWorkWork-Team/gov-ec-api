run:
	ENV=dev go run main.go

unit-test:
	ginkgo -r
	
start-dev-db-linux:
	docker run --name mysql-ec-dev -d \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=P@ssw0rd \
        --restart unless-stopped \
        -v ${PWD}/database:/docker-entrypoint-initdb.d \
        mysql:latest

#not try yet
start-dev-db-window:
	docker run --name mysql-ec-dev -d \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=P@ssw0rd \
        --restart unless-stopped \
        -v %cd%/database:/docker-entrypoint-initdb.d \
        mysql:latest

start-dev-php-admin:
	docker run --name phpmyadmin-ec -d --link mysql-ec-dev:db -p 8080:80 phpmyadmin/phpmyadmin