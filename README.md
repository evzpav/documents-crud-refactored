# Documents CRUD Refactored

- Using Domain Driven Desing approach
- CRUD of documents(CPF or CNPJ), with flags type and isBlacklisted.
- Backend(server) done in Golang with Echo framework.
- Structured to run on Docker, instructions are below:

## Instructions to run:

### Pre-requisites: 
#### Server: Docker and docker-compose installed
#### Client: NodeJs installed

```bash
# Run server on docker:
make run
# Server will be running on: http://localhost:3000

# Install dependencies and run frontend:
make run-front
# Client will be running on: http://localhost:8080


```