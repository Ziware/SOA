#!/bin/bash

COOKIE_STRING="$1"

eval "curl -X POST \
  -H 'Cookie: $COOKIE_STRING' \
  -H 'Content-Type: application/json' \
  -d '{
        \"title\": \"New Post\",
        \"description\": \"This is a new test post\",
        \"is_private\": false,
        \"tags\": [\"test\", \"api\"]
      }' \
  localhost:8080/posts/create -v"
