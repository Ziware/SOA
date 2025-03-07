#!/bin/bash

curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"login": "ziware", "password": "aboba"}' \
  localhost:8080/users/login -v
