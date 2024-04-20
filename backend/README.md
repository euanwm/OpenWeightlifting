# The Backend
Yes, it's written in Go. If you struggle with Go then you shouldn't be touching anything else in the codebase.

It was chosen for its simplicity and speed. The concept was originally written in Python and it was horrifically slow so I spent about 5 days porting (and learning) Go to see how it would fare. There was around a 20x increase in response times, along with build / startup times.

## Running Locally
Pretty simple, just run the following commands:
```bash
go build -o backend
./backend local
```

## Endpoint Documentation
The endpoint documentation is generated using Swaggo. To view the current production documentation, you can visit [this link](https://api.openweightlifting.org/swagger/index.html).
You'll need to re-run the following command to generate the documentation:
```bash
swag init --parseDependency --parseInternal -g main.go
```
If you do not have Swaggo installed, please visit the [Swaggo GitHub page](https://github.com/swaggo/swag) for installation instructions.

__All endpoints are defined in endpoints.go, and the main server is defined in  main.go.__


## Testing
When implementing new features, please ensure that you write tests for them. We like to aim for 80% test coverage however, as long any PR does not negatively impact the current test coverage, it will be accepted.
To run the tests with coverage, you can run the following command:
```bash
go test -coverprofile=coverage.out ./...
```