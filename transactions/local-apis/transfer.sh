#! /bin/bash

curl -X POST 'http://localhost:1271/v1/local-api/transfer' -d '{"from": "aaaaaaaaaaaaaaaa", "to": "bbbbbbbbbbbb", "amount": 123}'
