# Usage
SSN-Generator is a command line tool for swedish social security numbers - SSN-s. SSN-s that end with 98xx or 99xx are not used by real living persons and are considered safe for testing. -check and -safe can be combined - will also return the SSN provided as a safe variant
  -check
        Will check an SSN, and return correct checksum if needed
  -female
        Generate female only
  -male
        Generate male only
  -max float
        Sets max age boundary (default 100)
  -min float
        Sets min age boundary
  -n int
        Select number of SSNs to generate (default 1)
  -safe
        Safe flag will restrict to safe 980C-999C three last digits SSN
# Go compiler
Requires [Golang](https://golang.org/) to compile
# Install
```
go get github.com/markysand/ssngenerator
```
# Run
Make sure to have your $GOPATH in your $PATH file
