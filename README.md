# Documents CRUD Refactored

- Using Domain Driven Desing approach
- CRUD of documents(CPF or CNPJ), with flags type and isBlacklisted.
- Backend(server) done in Golang with Echo framework.
- Structured to run on Docker, instructions are below:

##Instructions to run:

## Method 1 - Docker
###Pre-requisites: docker and docker-compose installed
```bash
# Run docker compose:
sudo docker-compose up

#Server will be running on: http://localhost:3000

```

## Method 2 - Run locally
###Pre-requisites: make, go and docker installed
```bash
# Run Server:
make all

#Server will be running on: http://localhost:1323
```



docker run --name postgres_local -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:alpine

docker exec -i postgres_local psql -U postgres
CREATE DATABASE documents;


docker run -v $PWD/internal/storage/postgres/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/documents?sslmode=disable goto 1


docker run -v $PWD/internal/storage/postgres/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/documents?sslmode=disable version
