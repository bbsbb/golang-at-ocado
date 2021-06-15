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
In this part of the project, we'll be building the initial version of the sorting service.

Implement the following:
 * LoadItems - loads an input array of items in the service. E.g. ["tomatoes", "cucumber", "potato", "cheese"]
 * SelectItem -> Choose an item at random from the remaining ones in the array. E.g. choose "tomatoes" at random && remove item from existing array
 * MoveItem -> Move the selected item in the input cubby. Simply return "Success" here.

Return an error in any of the following cases:
 * SelectItem is invokes but there are no items in input bin
 * MoveItem is invoked but no item is selected yet
 * SelectItem is invoked when an item is already selected

