#! /bin/bash

curl -X POST http://localhost:8080/a/create \
  -H "Content-Type: application/json" \
  -d '{
    "slug": "my-new-article",
    "title": "My New Article",
    "summary": "This is a summary of my new article.",
    "content": "This is the content of my new article."
  }'