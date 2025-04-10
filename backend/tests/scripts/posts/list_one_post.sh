#!/bin/bash

COOKIE_STRING="$1"
PAGE_NUMBER="$2"

eval "curl -X GET \
  -H 'Cookie: $COOKIE_STRING' \
  -H 'Content-Type: application/json' \
  localhost:8080/posts?page_size=1\&page_number=$PAGE_NUMBER -v"
