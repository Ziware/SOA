#!/bin/bash

COOKIE_STRING="$1"
POST_ID="$2"

eval "curl -X GET \
  -H 'Cookie: $COOKIE_STRING' \
  -H 'Content-Type: application/json' \
  localhost:8080/posts/$POST_ID -v"
