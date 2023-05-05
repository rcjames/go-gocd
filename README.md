# Go GoCD
A Go library for the GoCD API.

## TODO
- Create pipeline for PR
    - Unit test
        - `go test -tags=unit`
    - Static code analysis
    - Conventional commit checks
    - Integration test
        - Spin up GoCD server
            - docker run -d -p8153:8153 --name gocd gocd/gocd-server:v23.1.0
            - Wait for startup
        - Use `go test -tags=integration`
- Create pipeline for merge to main
    - Conentional commits changelog
        - https://github.com/marketplace/actions/conventional-changelog-action
- Create pipeline for release being cut
    - Automatic release
        - https://go.dev/doc/modules/publishing
- Start work on all the other APIs...