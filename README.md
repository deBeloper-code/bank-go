# Bank Go

## Estructura del Proyecto

```plaintext
/cmd
  /app
    main.go
/internal
  /api
    handler.go
  /domain
    model.go
    repository.go
    service.go
  /adapter
    /http
      handler_adapter.go
    /database
      repository_adapter.go
```

## Directory Descriptions

### `/cmd`

This directory holds application-specific logic, such as initialization and configuration.

- **/app**: This directory contains the main entry point of the application (`main.go`). This file initializes the application and configures dependencies.

### `/internal`

This directory contains the internal code of the project, organized into layers.

- **/api**: Here, you'll find HTTP controllers that interact with the outside world. These controllers handle HTTP requests and invoke domain services to perform operations.
- **/domain**: It holds the core business logic of the application.
  - **model.go**: Defines fundamental data structures for your domain, such as user models.
  - **repository.go**: Defines interfaces for repositories handling storage and data retrieval.
  - **service.go**: Contains logic for domain services that implement specific business operations.
- **/adapter**: Contains adapters that connect external layers with internal layers.
  - **/http**: Contains HTTP adapters that handle HTTP requests and translate them into calls to domain services.
    - **handler_adapter.go**: Implements HTTP controllers that interact with the outside world.
  - **/database**: Contains adapters to interact with the database.
    - **repository_adapter.go**: Implements repositories defined in the domain package using a specific storage system (e.g., a database).
