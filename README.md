## Dependencies
- Be sure to have Go (1.20+)
- Be sure to have postgres running locally
___
## Database
- Create db user with
  ```
  CREATE USER bt_data_server_user;
  ```
- Set up development db with
  ```
  createdb bt_data_server_dev
  ```
- Run migration:
  ```
  psql -U bt_data_server_user -d bt_data_server_dev -a -f internal/storage/migrations/schema.sql
  ```
___
## Usage
### Development
```
go run cmd/server/main.go
```
___
## Testing
Standard:
```
go test ./...
```

