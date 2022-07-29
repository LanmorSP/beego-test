# Just a testcase for beego

## how to run & devlop

1. container up all of service
```
docker-compose up --build
```

2. exec to the app container
```
docker exec -it app.bee sh
```

3. database migrations 
you are in container app.bee,you can run the migrate command.
```
bee migrate -driver=postgres -conn="postgres://test:test1234@localhost:5432/test?sslmode=disable"
```

4. launch the server
```
bee generate routers && bee run --runmode=dev
```

- now you can go localhost:8081 to visit App.

- and you can go localhost:8082 to visit database admin with phppgadmin and 
use test/test1234 to login.




