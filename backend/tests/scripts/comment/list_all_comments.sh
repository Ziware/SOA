#!/bin/bash

COOKIE_STRING="$1"
POST_ID="$2"
PAGE_NUMBER="$3"

eval "curl -X GET \
  -H 'Cookie: $COOKIE_STRING' \
  -H 'Content-Type: application/json' \
  localhost:8080/posts/$POST_ID/comments?page_size=100\&page_number=$PAGE_NUMBER -v"
