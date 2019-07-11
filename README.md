TODO API
--------
The project is a demo CRUD app, it exposes a POST endpoint and a GET
endpoint for creating some to do list items and getting list of created
items.

Each to do item has an `id`, a `title` and a `done` field.


In order to setup db and run migrations of project, run

```
DB_PASSWORD=<some pass> docker-compose up
```

If you want to drop db for any reason run

```
DB_PASSWORD=<some pass> docker-compose down
```

the in order to run the project run
```
GO111MODULE=on DB_PASSWORD=<some pass> go run main.go
```
