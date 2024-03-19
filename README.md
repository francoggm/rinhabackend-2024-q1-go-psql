# Implementação para a rinha de backend 2024 - Q1
- [Repositório da rinha](https://github.com/zanfranceschi/rinha-de-backend-2024-q1)

## Tecnologias usadas
- **Go**
- **Postgres (pool)** 
- **Docker compose**

### Gatling para os testes de estresse

![Gatling](./img/gatling%20test.png)

### Comando para rodar os containers do projeto

- Subir toda infra

```
docker compose up --build
```

- Subir apenas postgres para testar local

```
docker compose -f docker-compose.debug.yml up --build
```

#### Suporte para Makefile

```
make up
```

```
make up-debug
```