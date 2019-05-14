# documents-crud-refactored

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
