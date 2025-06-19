# REST API

## Running the code

* Change the version of the go language in ```go.mod``` file to match yours
* While running for the first time: run ```go mod tidy```
* Run ```go build .```
* Run ```./netflix_wrapped_api```

## Testing

* curl http://localhost:8080/test
* curl http://localhost:8080/test --include --header "Content-Type: application/json" -X "POST" -d "@test_data/test.txt"
