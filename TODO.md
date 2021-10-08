Must:
- Write at least a few tests
- Improve -h screen to talk about passing the path
- look for @TODOs
Should:
- Move most logic outside the cmd dir
- Consider separating out code more into funcs
- Extensive test coverage
- "go run . -c=5 /Applications/" works but "go run . /Applications/ -c=5" ignores c