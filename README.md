# pushshift

A Go client for the [pushshift.io reddit API](https://pushshift.io/api-parameters/).

## Installation

Install the package with

`go get -u github.com/mfdeux/pushshift`

## Authentication

The pushshift API does not require authentication. However, the API is rate-limited server-side.

## Examples

For detailed examples, check out the examples folder.

```Go
// Returns a new unauthenticated client for invoking the API
client := pushshift.NewClient("myApp/0.1.0")

```
