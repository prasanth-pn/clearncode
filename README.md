# clearncode
This is the documentation for the avua backend, which employs a monolithic architecture. The backend follows a clean architecture design pattern.

## Services

Following are the main services:

- applicant

Each of the above services and their sub modules follow a similar pattern as listed below:

- A repository layer whose primary purpose is to interact with the database.
- A service layer where all the business logic is implemented.
- A http_handler layer to handle the request and the response.
- A contracts file, responsible for request contracts.
- A constants file, used to store all the constants that are local to that particular module.
- A middleware file, to handle the middlewares such as auth middlewares etc.
- A models file, consisting of DSOs for each of the schemas in the database.
- Finally, a view layer, that consists of DSOs/response structure.


Along with the services, there are other utilities (folders) as listed below:

- `cmd/api`: This is where the main file resides.
- `boot`: Responsible for server setup. This is part of the Gin framework.
- `clients`: Responsible for interacting with external clients or services.
- `config`: Where all the necessary project configurations are stored and loaded.
- `config/loadev`: Where environment files are stored.
- `db`: Where the connections to the database are established and where the migrations are stored.
- `dependencies`: Used to inject dependencies during server boot.
- `jwt`: Manages JWT tokens.
- `utils`: Contains helper functions.
- `vendor`: Used for package management.