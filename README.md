# pushshift

[![](https://godoc.org/github.com/mfdeux/pushshift?status.svg)](http://godoc.org/github.com/mfdeux/pushshift)

[![Go Report Card](https://goreportcard.com/badge/github.com/mfdeux/pushshift)](https://goreportcard.com/report/github.com/mfdeux/pushshift)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmfdeux%2Fpushshift.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmfdeux%2Fpushshift?ref=badge_shield)

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


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmfdeux%2Fpushshift.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmfdeux%2Fpushshift?ref=badge_large)