#!/bin/bash

COOKIE_STRING="$1"

eval "curl -H 'Cookie: $COOKIE_STRING' -X PUT \
  -H 'Content-Type: application/json' \
  -d '{\"name\": \"chelik\", \"email\": \"biba@zalupa.ru\", \"birth_date\": \"11-12-2004\", \"surname\": \"bobikovich\", \"phone_number\": \"+212312312\"}' \
  localhost:8080/users/profile -v"
