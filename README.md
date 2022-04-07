# boilerplate
Golang service builderplate

This service created as a template for future development and as example how service layout may look like

http router
gorilla mux + new relic integration

database driver 
pgx + new relice integration

test
testify + mockery

## Requirements and tips
### install mockery
go install github.com/vektra/mockery/v2@latest

### regenerate mocks
go generate ./... (in poject root)

### run tests with coverage
go test ./... -cover
