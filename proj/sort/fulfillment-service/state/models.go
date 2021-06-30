package state

import "github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"

type itemStatus struct {
	item    *gen.Item
	inCubby bool
}

type ItemInfo struct {
	OrderId string
	Index   int
	CubbyId string
}

type OrderInfo struct {
	State gen.OrderState
	Order *gen.Order
	Cubby *gen.Cubby
}
