#!/bin/bash

curl -i -k -L POST 'https://gigachat.devices.sberbank.ru/api/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer REQUEST_TOKEN' \
  -d '{
    "model": "GigaChat",
    "messages": [
      {
        "role": "user",
        "content": "В каком году избрался Ельцин?"
      }
    ],
    "temperature": 0.002
    "repetition_penalty": 1
  }'