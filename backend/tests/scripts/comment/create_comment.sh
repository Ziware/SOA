#!/bin/bash

COOKIE_STRING="$1"
POST_ID="$2"
COMMENT_TEXT="$3"

eval "curl -X POST \
  -H 'Cookie: $COOKIE_STRING' \
  -H 'Content-Type: application/json' \
  -d '{\"text\": \"$COMMENT_TEXT\"}' \
  localhost:8080/posts/$POST_ID/comments -v"
