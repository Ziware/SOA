#!/bin/bash

COOKIE_STRING="$1"

eval "curl -X GET \
  -H 'Cookie: $COOKIE_STRING' \
  -H 'Content-Type: application/json' \
  localhost:8080/posts?page_size=100 -v"
