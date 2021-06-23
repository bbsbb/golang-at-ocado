package main

type FulfillmentFailedError struct{}

func (err *FulfillmentFailedError) Error() string {
	return "something went wrong"
}
