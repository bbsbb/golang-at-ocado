# sort-system-v1

This is the first version of the [SORT system](https://www.youtube.com/watch?v=BQDliV7w7_8) we'll be building in the course.

## How to run the project
 * `make grpc-compile` to generate all grpc-related files in the `gen/` folder
 * Enter `sorting-service` and type `go run *.go`

## (Optional) Name your project the way you like
 * Modify the following files and change the repository reference from `github.com/bbsbb/go-at-ocado/sort-vX` to your own repo:
   * `go.mod`
   * `Makefile.GRPC`
   * all files in the `idl` directory 
   * `sorting-service/go.mod`, `sorting-service/main.go`, `sorting-service/service.go`

## Assignment
In this part of the project, we'll be building the initial version of the fulfillment service which controls the sorting robot.
The responsibility of the service is to receive a set of items, which are associated to orders and dispatch operations to the robot in order to sort the items in appropriate cubbies for the orders.

