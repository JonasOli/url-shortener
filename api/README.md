# Go API

## Database migrations

To run a database migration just run:

```bash
migrate -path=internal/model/migration -database "postgresql://myuser:mypassword@localhost:5432/url-shortener?sslmode=disable" -verbose up
```
