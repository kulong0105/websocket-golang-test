#!/bin/bash

set -x

echo {\"a\": \"hello\", \"b\": 100}  | websocat -n  ws://127.0.0.1:8080/json
echo asdf |websocat -n  ws://127.0.0.1:8080/text


timeout 3 curl -i -vvv -N -H "Connection: Upgrade" \
              -H "Upgrade: websocket"  \
              -H "Sec-WebSocket-Version: 13" \
              -H "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ=="  http://127.0.0.1:8080/text


timeout 3 curl -i -vvv -N -H "Connection: Upgrade" \
              -H "Upgrade: websocket" \
              -H "Sec-WebSocket-Version: 13" \
              -H "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ=="  http://127.0.0.1:8080/json
