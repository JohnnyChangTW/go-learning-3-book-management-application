# Prerequisites
- Need to have MySQL installed and started
- Host: 127.0.0.1
- Port: 3306->3306
- User: root
- password: myPassword
- Database name: bookStore


## Use Docker to run MySQL container on M1 chip (arm64)
``` shell
docker run --name mysqlContainer -p 3306:3306 -e MYSQL_ROOT_PASSWORD=myPassword -d arm64v8/mysql:latest
```

## Before connecting to DB
```shell
docker exec -it mysqlContainer mysql -u root -p
```

``` MySQL Cli
CREATE TABLE bookStore
```
