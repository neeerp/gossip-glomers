# Specification

Our nodes will receive an `echo` message that looks as follows:

```json
{
  "src": "c1",
  "dest": "n1",
  "body": {
    "type": "echo",
    "msg_id": 1,
    "echo": "Please echo 35"
  }
}
```

In turn, the node must reply with the same body, but with message type `echo_ok`
and with an `in_reply_to` field referencing the original `msg_id`:

```json
{
  "src": "n1",
  "dest": "c1",
  "body": {
    "type": "echo_ok",
    "msg_id": 1,
    "in_reply_to": 1,
    "echo": "Please echo 35"
  }
}
```
