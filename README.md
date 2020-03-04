# folderauth

[![Build Status](https://travis-ci.org/calmandniceperson/folderauth.svg?branch=master)](https://travis-ci.org/calmandniceperson/folderauth) [![Go Report Card](https://goreportcard.com/badge/github.com/calmandniceperson/folderauth)](https://goreportcard.com/report/github.com/calmandniceperson/folderauth) [![GoDoc](https://godoc.org/github.com/calmandniceperson/folderauth?status.svg)](https://godoc.org/github.com/calmandniceperson/folderauth)

folderauth is a [Go](https://golang.org) package that wraps the [API](https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc) of the [Folder-based Authorization Strategy plugin](https://plugins.jenkins.io/folder-auth/) for Jenkins.

## How to install

`go get -u github.com/calmandniceperson/folderauth`

## Usage example

```go
c, err := folderauth.NewClient(hostname, username, token)
if err != nil {
    return err
}

// Use c here
```

## License

This package is published under the [MIT License](https://opensource.org/licenses/MIT).