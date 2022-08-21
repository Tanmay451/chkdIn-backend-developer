# chkdIn-backend-developer
This project is to demonstrate CRUD operations using Gin framework and postgres as database.

Design patterns: MVC

project requirement:
1. Golang
2. Postgres
3. Migrate [https://github.com/golang-migrate/migrate]

Project Setup:
1. Make sure you have installed above requirement.
2. To setup database, use "make db_up" command. This will use Migrate to setup database.
3. Setup environment file. Use ".env.sample", make a copy and rename it as ".env". Fill up proper values.
4. Run "make get", to install or update all requirement.
5. To run project in your local system, use command "make run" 
    or to make a build use command "make build". 
    In case you need to deploy this build on linux based arch, use command "make build_linux".