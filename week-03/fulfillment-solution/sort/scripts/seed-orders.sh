#!/bin/bash

items_1='[{"code": "123", "label": "tomato"},{"code": "456", "label": "cucumber"}]'
items_2='[{"code": "420", "label": "glass"},{"code": "222", "label": "fork"}]'
# items_3='[{"code": "111", "label": "english breakfast"},{"code": "333", "label": "beans in a can"}]'
# items_4='[{"code": "666", "label": "peaches"},{"code": "667", "label": "oranges"}]'
# items_5='[{"code": "501", "label": "headphones"},{"code": "502", "label": "keyboard"},{"code": "503", "label": "cat in a box"}]'
# items_6='[{"code": "601", "label": "book 1"},{"code": "602", "label": "book 2"},{"code": "603", "label": "book 3"},{"code": "604", "label": "book 4"}]'
# items_7='[{"code": "401", "label": "water bottle"},{"code": "402", "label": "wataaa"}]'
# items_8='[{"code": "301", "label": "juice"}]'
# items_9='[{"code": "201", "label": "toy"},{"code": "202", "label": "teddy bear"},{"code": "203", "label": "dinosaur"},{"code": "204", "label": "dog"},{"code": "205", "label": "mug"}]'
# items_10='[{"code": "101", "label": "laptop"},{"code": "102", "label": "mouse"}]'

grpcurl -d "{\"items\":$items_1}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
grpcurl -d "{\"items\":$items_2}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_3}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_4}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_5}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_6}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_7}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_8}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_9}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems
# grpcurl -d "{\"items\":$items_10}" -plaintext localhost:10000 sorting.SortingRobot.LoadItems

grpcurl \
    -d "{\"orders\":[{\"id\":\"1\", \"items\":$items_1},{\"id\":\"2\", \"items\":$items_2}]}" \
    -plaintext localhost:10001 fulfillment.Fulfillment.LoadOrders


    #-d "{\"orders\":[{\"id\":\"1\", \"items\":$items_1},{\"id\":\"2\", \"items\":$items_2},{\"id\":\"3\", \"items\":$items_3},{\"id\":\"4\", \"items\":$items_4},{\"id\":\"5\", \"items\":$items_5},{\"id\":\"6\", \"items\":$items_6},{\"id\":\"7\", \"items\":$items_7},{\"id\":\"8\", \"items\":$items_8},{\"id\":\"9\", \"items\":$items_9},{\"id\":\"10\", \"items\":$items_10}]}" \
