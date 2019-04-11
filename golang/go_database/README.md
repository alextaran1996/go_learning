## Info

This is simple program to make basic ops with mysql DB

## Usage

1. Build docker image `docker build . -t mysql`
2. Run docker container `docker run -d --name=mysql --rm  mysql`
3. Connect to mysql DB `mysql -h 172.17.0.2 -u root -p'root'`
4. Create DB `create database golang;`
5. Create Table `create table golang.users (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY, `name` varchar(255) NOT NULL);`
6. Go to src folder and run `go run *.go  -c ../config.yml -n Aliaksandr` to create your first user