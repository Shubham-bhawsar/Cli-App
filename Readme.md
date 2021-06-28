## Title
Cli-App
## Description of the project
This is CLi based app which downloads the media from specified URL,You are needed to
provide the URL and the optional path where you want to store the data.

## Application Requirement

1. golang https://golang.org/dl/
2. Cobra ,which is both a library for creating powerful modern CLI applications and a program for generating applications and batch files.
   `$ go get -u github.com/spf13/cobra/cobra`

## How to use

- Run `go install` for bulding the project.
- Run `CliApp download --url="Any URL"` for downloading the required file in default directory.
- Run `CliApp download --url="Any URL" --path="/$HOME/Downloads/xyz.jpg"` for downloading the required file in desired directory which is provide in `path` flag.

## Testing guide
### Build
-   Run `go install` for building the project.
### Test
-   Run `cd cmd` from root directory of project
-   Run `go test` from cmd directory for running the tests.
## Contributors
-   Shubham Bhawsar