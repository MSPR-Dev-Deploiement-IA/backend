# Plant Care App

This is a backend service for a Plant Care application. The application allows users to manage and share their plants for others to take care of, and to view the history of their own and others' plant care.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or later
- PostgreSQL 13 or later

### Installation

1. Clone the repository
    ```
    git clone https://github.com/yourusername/plant-care-app.git
    ```
2. Change directory to the project
    ```
    cd plant-care-app
    ```
3. Install Go dependencies
    ```
    go mod tidy
    ```
4. Set up your PostgreSQL database and update the connection string in `main.go`
   
### Running the Application

To start the application, run:


The application will be running on `http://localhost:8080`.

## API Endpoints

- `GET /api/user`: Fetch the current user's information
- `PUT /api/user`: Update the current user's information
- `GET /api/cared_plants`: Fetch plants a user has cared for
- `GET /api/plants_cared_by_others`: Fetch plants that have been cared for by others
- `GET /api/plants_to_be_cared`: Fetch plants to be cared for

## Built With

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework written in Go
- [GORM](https://gorm.io/) - Fantastic ORM library for Golang

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/yourusername/your-contributing-md-url) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
