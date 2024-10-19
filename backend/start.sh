#!/bin/bash
if [ ! "$(docker ps -a | grep tinyWebSql)" ]; then
  docker run -d --name tinyWebSql -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 mysql:latest
else
  docker start tinyWebSql
fi

go run main.go