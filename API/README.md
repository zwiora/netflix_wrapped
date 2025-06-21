# REST API

## Running the code

* Change the version of the go language in ```go.mod``` file to match yours
* While running for the first time: run ```go mod tidy```
* Create file `API/tmdbKey.go` (must remain in .gitignore!) and save the api key in it:

```
package main

import "os"

func setApiKey() {
	os.Setenv("TMDB_API_KEY", #PUT THE API KEY HERE#)
}
```

* Run ```go build .```
* Run ```./netflix_wrapped_api```

## Testing

```curl http://localhost:8080/generate --include --header "Content-Type: application/json" -X "POST" -d "@test_data/example.json"```
