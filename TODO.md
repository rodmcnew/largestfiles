## Must
## Should
- Basic test coverage (see https://pkg.go.dev/testing/fstest for FS mocking)
## Could
- Extensive test coverage
- "go run . -c=5 /Applications/" works but "go run . /Applications/ -c=5" ignores c (an issue with how go std lib args work)
