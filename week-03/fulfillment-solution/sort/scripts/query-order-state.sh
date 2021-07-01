#!/bin/bash

grpcurl \
    -d "{\"orderId\": \"1\"}" \
    -plaintext localhost:10001 fulfillment.Fulfillment.GetOrderStatusById
