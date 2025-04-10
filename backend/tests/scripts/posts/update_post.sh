#!/bin/bash

COOKIE_STRING="$1"
POST_ID="$2"

eval "curl -X PUT \
  -H 'Cookie: $COOKIE_STRING' \
  -H 'Content-Type: application/json' \
  -d '{
        \"title\": \"Updated Title\",
        \"description\": \"Updated description\",
        \"is_private\": true,
        \"tags\": [\"updated\", \"test\"]
      }' \
  localhost:8080/posts/$POST_ID -v"
