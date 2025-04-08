#!/bin/bash

COOKIE_STRING="$1"

eval "curl -H 'Cookie: $COOKIE_STRING' -X GET \
  -H 'Content-Type: application/json' \
  localhost:8080/users/profile -v"
