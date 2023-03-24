# postgres
### CREATE DATABASE USING POSTGRES
> Create database "**postgres**"

create file **.env** then fill these to configure Makefile:
>DB_PORT=5431
DB_NAME=postgres
DB_HOST=127.0.0.1
DB_USERNAME=postgres
DB_PASSWORD=postgres


### INSTAL GOLANG MIGRATE
[Release Downloads](https://github.com/golang-migrate/migrate/releases)
```shell
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```

### RUN Makefile
Create two files with up and down at the suffix of the created sql files.
```shell
make migrate
```
Execute migratedown command
```shell
make migratedown
```
Execute migrateup command
```shell
make migrateup
```

