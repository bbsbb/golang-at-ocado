# sort

This is an educational version of the [SORT system](https://www.youtube.com/watch?v=BQDliV7w7_8) we'll be building in the course.

## How to run the project
 * `make grpc-compile` to generate all grpc-related files in the `gen/` folder
 * Enter `sorting-service` and type `go run *.go`
 * Enter `fulfillment-service` and type `go run *.go`

## (Optional) Name your project the way you like
 * Modify the following files and change the repository reference from `github.com/bbsbb/go-at-ocado/...` to your own repo:
   * `go.mod`
   * `Makefile.GRPC`
   * all files in the `idl` directory 
   * `sorting-service/go.mod`, `sorting-service/main.go`, `sorting-service/service.go`
   * `fulfillment-service/go.mod`, `sorting-service/main.go`, `sorting-service/service.go`

## Assignment
In this part of the project, we'll be building the initial version of the fulfillment service which controls the sorting robot.
The responsibility of the service is to receive a set of items, which are associated to orders and dispatch operations to the robot in order to sort the items in appropriate cubbies for the orders.

 * Implement the Fulfillment/LoadOrders rpc routine. It accepts a list of `Order`s and returns a `list of order to cubby mappings`
 * Don't just return the correct mapping, control the robot so that the items are **actually placed in the correct cubby**

**Implementation details:**
 * The total number of cubbies is always `10` for this part of the project
 * Cubby IDs are in the range [1..10]
 * All items of a given order should be associated to the same `cubby`
   * To determine the cubby for a given order, use the [ordertocubby library](https://github.com/preslavmihaylov/ordertocubby)
   * In case of collisions, e.g. two orders are mapped to cubby `5`, remap the second one using the same library.

To test your implementation, use the `./scripts/seed-orders.sh` script.

A starting point for setting up a grpc client - https://grpc.io/docs/languages/go/basics/
