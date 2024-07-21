## Docker CLI Commands

1. You can manage the database via `psql` with the command:
```bash
docker-compose exec db psql -U <DB_USER (from the environment variables)>
```
