# Web-Analyzer

This is a simple web application that takes a URL as input, scrapes the web page, and provides details such as the HTML version, title, headings, internal and external links, inaccessible links, login form presence, and accessibility of URLs.

## Features

- Scrapes a given URL for the following details:

  - HTML version
  - Page title
  - Number of headings by level
  - Internal and external links
  - Inaccessible links
  - Presence of a login form
  - Accessibility of URLs

- Displays results in a user-friendly format, including a table for URL accessibility.
- Swagger documentation for the API.

## Prerequisites

- Go (version 1.21+)
- Gin web framework
- `golang.org/x/net/html` package for HTML parsing
- Swaggo for Swagger documentation

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/venusaran/web-analyzer.git
   cd web-analyzer
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Install `swag` for generating Swagger documentation:

   ```sh
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

4. Generate Swagger documentation:
   ```sh
   swag init
   ```

## Running the Application

1. Start the Go server:

   ```sh
   go run cmd/main.go
   ```

2. Open your web browser and navigate to `http://localhost:8080`

## Usage

1. Enter a URL in the input box.
2. Click the "Submit" button.
3. Wait for the spinner to disappear, indicating that the data has been fetched.
4. View the results, including a table of accessible URLs.

## Accessing Swagger Documentation

Once the server is running, you can access the Swagger documentation at:
[Swagger Docs](http://localhost:8080/docs/index.html#/)

## Project Structure

```
.
|
api/
└── rest/
|   └── controller/
|      └── analyzer/
|         └── analyzer.go
router/
└── router.go
cmd/
└── main.go
docs/
└── docs.go
└── swagger.json
└── swagger.yaml
internal/
└── service/
|   └── scraper/
|   └── helper.go
|   └── scraper.go
pkg/
└── constants/
|   └── constants.go
└── interfaces/
|   └── interfaces.go
util/
└── utility.go
static/
└── index.html
vendor/
.gitignore
go.mod
go.sum
README.md
```

## Example

Here's a sample output you can expect from the application:

```
HTML Version: HTML 5
Title: Example Title
Headings: {"h1": 2, "h2": 3, "h3": 1}
Internal Links: 5
External Links: 10
Inaccessible Links: 2
Login Form Present: true
Accessible URLs:

https://example.com: Accessible
https://example.org: Inaccessible
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP web framework for Go
- [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html) - HTML parsing library
- [Swaggo](https://github.com/swaggo/swag) - Automatically generate RESTful API documentation with Swagger 2.0 for Go.
