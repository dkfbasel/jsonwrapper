# Jsonwrapper Package for Go

The package jsonwrapper provides some convenience function to work with
wrapped json data in go.

It follows the conventions outlined on http://www.restapitutorial.com.

## Description
JSON responses are wrapped in a struct containing a code (http-status-code) as integer,
a status (success, fail, error), an optional message for responses with status fail or error
and the data containing the actual payload.

*From http://www.restapitutorial.com:*
> Basically, current best practice is to wrap regular (non-JSONP) responses with the following properties:
> - **code** – contains the HTTP response status code as an integer.
> - **status** – contains the text: “success”, “fail”, or “error”. Where “fail” is for HTTP status response values from 500-599, “error” is for statuses 400-499, and “success” is for everything else (e.g. 1XX, 2XX and 3XX responses).
> - **message** – only used for “fail” and “error” statuses to contain the error message. For internationalization (i18n) purposes, this could contain a message number or code, either alone or contained within delimiters.
> - **data** – that contains the response body. In the case of “error” or “fail” statuses, this contains the cause, or exception name.

## Usage
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/tikiatua/jsonwrapper"
)

func main() {

	// get a wrapped successful response (other options are error and fail)
	response := jsonwrapper.Success([]string{"test", "data"})

	// marshal to json as byte-slice (marshalling to string included as well)
	marshalled, _ := jsonwrapper.ToJsonBytes(&response)

	// print as json string (or write to http response-writer)
	fmt.Printf("%s\n", marshalled)

}
```
## Roadmap
- [x] Add functions to marshal/unmarshal wrapped responses directly
- [ ] Add a global configuration option to toggle wrapping for all responses
- [ ] Add a convenience function to generate a http-response directly from the response (using the status code provided)

## License
MIT - feel free to use.

## Contribute
Any ideas for improvements, comments or pull-request are very welcome.
