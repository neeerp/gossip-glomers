# Unique ID Generation

## Specification

Nodes will now receive a `generate` request message body that looks like this:

```json
{
  "type": "generate",
  "msg_id": 1
}
```

The nodes should reply with a `generate_ok` response and a unique ID:

```json
{
  "type": "generate_ok",
  "msg_id": 1,
  "in_reply_to": 1,
  "id": 123
}
```

IDs may be of any data type, including strings, booleans, integers, floats,
arrays, etc.

## Testing

Run `maelstrom` with the following command.

```sh
./maelstrom test -w unique-ids --bin ~/go/bin/unique --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
```

This runs a 3 node cluster for 30 seconds, and requests new IDs at a rate of
1000 RPS. The test verifies that all IDs generated are unique and that the system
is totally available even in the face of network partitions/interruptions.

If you get a happy response, then you'll know you're done :)

## Stuff I tried

### Simple counter

What if we naively have a counter variable in each node?

```go
    a := 0
	n.Handle("generate", func(msg maelstrom.Message) error {
        // ...
		a++
		return n.Reply(msg, body)
    }
```

This obviously won't work... Every node is going to be counting from the same
starting point, so there'll necessarily be overlap.
