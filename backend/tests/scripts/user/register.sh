#!/bin/bash

curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"login": "ziware", "email": "ziware@test.com", "password": "aboba"}' \
  localhost:8080/users/register -v
