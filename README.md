# Go GoCD
A Go library for the GoCD API.

## TODO
- Create pipeline
    - Unit test
        - `go test -tags=unit`
    - Static code analysis
    - Conventional commit checks
    - Integration test
        - Spin up GoCD server
            - docker run -d -p8153:8153 --name gocd gocd/gocd-server:v23.1.0
        - Use `go test -tags=unit` and `go test -tags=integration`
    - Automatic release
        - https://go.dev/doc/modules/publishing
- Start work on all the other APIs...