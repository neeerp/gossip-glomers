# Echo

This first challenge corresponds to [echo](https://fly.io/dist-sys/1/) on the
Glossip Glomers site.

## Specification

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

## Testing

### Prerequisites

Make sure you have `go` installed.

Make sure to install the
[prerequisites](https://github.com/jepsen-io/maelstrom/blob/main/doc/01-getting-ready/index.md#prerequisites)
for the `maelstrom` library. Afterwards, follow the [installation
instructions](https://github.com/jepsen-io/maelstrom/blob/main/doc/01-getting-ready/index.md#installation).

### Running echo

Compile the program by running `go install .`; it should appear in your
`$GO_PATH` (in my case `~/go/bin/`).

Now invoke `maelstrom` with our binary:

```sh
./maelstrom test -w echo --bin ~/go/bin/maelstrom-echo --node-count 1 --time-limit 10
```

If it terminates with a happy looking message, everything's good!
