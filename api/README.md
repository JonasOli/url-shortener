# Go API

## Useful links

- <https://docs.gofiber.io/guide/error-handling>
- <https://docs.docker.com/guides/golang/run-containers/>

## Database migrations

To create a migration:

```bash
migrate create -ext=sql -dir=internal/model/migration -seq migration_name
```

To run a database migration:

```bash
migrate -path=internal/model/migration -database "postgresql://myuser:mypassword@localhost:5432/url-shortener?sslmode=disable" -verbose up
```
