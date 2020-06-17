# Restful API Template
### Keys: Golang, Echo, Repository Pattern

#### Setup
* Create config file: `cp config-sample.yaml config.yaml`
* Config database info
* Install require packages: `go mod vendor`

#### Startup
* Run App: `go run -mod=vendor main.go`

## Module structure
Each module is spitted by repository pattern
##### Repository Pattern
![alt text](https://imgur.com/VWxivOX.png "Repository Pattern")


##### Module api structure
* `routers/`: define api url, request body, params
* `extensions/`: setup base configuration
* `helpers/`: define helper function
* `models/`: define orm model  
* `repositories/`: define repository to access data
* `services/`: handle business logic
* `tests/`: app test script

### 📙 Resource

#### 📙 Libraries
- Echo: https://echo.labstack.com/

### Contributing
If you want to contribute to this boilerplate, clone the repository and just start making pull requests.
