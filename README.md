# Spotlas_test_backend_dev

## Installation for local tests/development
1. Run the docker-compose file to set up a local postgres db.
2. Run the spots.sql file to create the table and fill it with data.

## Tasks solutions
### 1. Creating a query

In the task1 folder you can find two files:

a) spots_domains_counts.sql  
returns all the spots with their domains and counts for domains as per task 1;

b) only_domains_and_counts.sql
returns all the domains and counts for domains;

### 2. Creating an API endpoint
In the task2 folder you can find the implementation of a grpc server.
The config folder uses viper to read the config file and set the environment variables. The default values are:

DB_ADDRESS=postgres://spots_db_user:spots_db_pass@localhost:45678/spots_db?sslmode=disable;

GRPC_ADDRESS=0.0.0.0:9090;

Prebuilt binaries are available for linux os and windows os
in the task2 folder. To run the server you can use the following commands:

linux os:

./close-spots-service.exe serve

windows os:

./close-spots-service serve

In cmd/main.go you can find the main function that starts the application.
Cobra is used to create the cli commands. The serve command starts the grpc server.