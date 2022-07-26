# Just a testcase for beego

## how to run

```
docker-compose up --build
```

- you can go localhost:8081 to visit App.

- you can go localhost:8082 to visit database admin with phppgadmin and 
use test/test1234 to login.

## database migration

you can run migrate to generate base database when you docker-compose up and bash in app container.

1. exec to the app container
```
docker exec -it app.bee sh
```

2. and you are in container app.bee,you can run the migrate command.
```
bee migrate -driver=postgres -conn="postgres://test:test1234@postgres-db:5432/test?sslmode=disable"
```
