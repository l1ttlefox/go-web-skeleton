#go-web-skeleton Installation/Structure guide

##Installation
* Clone this repository
* Install [glide package manager](https://glide.sh/)
* go to the root of application and run ```glide install```
* Set up config.json file from /config folder
* Try to run ``` go run main.go ```

##Structure

```
This project has the following structure:

 main.go          # Main file where application start
 glide.yaml       # File where project dependencies are saved
 glide.lock       # File where version of dependencies is locked
 /config          # Folder where dev/prod configurations are stored
 /vendor          # Folder where are stored our dependencies
 /app             # Main folder where is stored the code of whole app
 /app/models      # Database queries
 /app/controllers # Requests logic organized by HTTP methods (GET, POST, PUT, DELETE)
 /app/shared      # packages for database, cryptography, json etc.
 /app/routes      # Routes information and middlewares
```
##External libraries

```
github.com/gin-gonic/gin      - Web framework
github.com/jmoiron/sqlx       - Postgresql general purpose extensions   
github.com/lib/pq             - Postgresql driver
golang.org/x/crypto/bcrypt    - Cryptohraphy library
```
