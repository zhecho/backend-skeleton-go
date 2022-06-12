# backend-skeleton-go

```text
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+                                           Installation                                           +
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
```

# Installing packages
```bash
    # Install httprouter package Note: Only Explicit Maches Router!
    go get -u github.com/julienschmidt/httprouter

```

# Running local dev postgres image
```bash
    # run docker && run docker container with the following cmd
    docker run  \
        --name myPostgresDb \
        -p 5455:5432 \
        -e POSTGRES_USER=postgresUser \
        -e POSTGRES_PASSWORD=postgresPW \
        -e POSTGRES_DB=go_movies \
        -d postgres
    
    # install postgresql client (osX)
    brew install libpq  && brew link --force libpq
    
    # import sql (check out and change your psql exposed port )
    psql -h localhost  -U postgresUser -p 5455 -d go_movies -f ./41_go_movies.sql
    
    # check your db with CLI or favorite GUI client  
    psql -h localhost -U postgresUser -p 5455 -d go_movies 

    # install needed go module for postgres
    go get -u github.com/lib/pq@v1.10.0


```







