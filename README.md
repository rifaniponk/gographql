## Requirement

* docker version 19.03 or higher
* docker-compose version 1.25.5 or higher

## Getting Started

```sh
docker-compose build && docker-compose up
./go.sh sql-migrate up
```

Graphql Playground
http://localhost/

Graphql Query
http://localhost/query

pgweb
http://localhost:8070/ 


## Graphql Generator

```sh
./go.sh gqlgen
```

## migration

```sh
./go.sh sql-migrate down      # Undo a database migration
./go.sh sql-migrate new       # Create a new migration
./go.sh sql-migrate redo      # Reapply the last migration
./go.sh sql-migrate skip      # Sets the database level to the most recent version available, without running the migrations
./go.sh sql-migrate status    # Show migration status
./go.sh sql-migrate up        # Migrates the database to the most recent version available
```

