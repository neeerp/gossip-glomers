# Broadcast

## Specification

Nodes will handle three RPC message types, `broadcast`, `read`, and `topology`.
The node will store the set of integer values it sees in `broadcast`s so that
they can be later returned via `read`s.

### Message types

#### Broadcast

This message requests that a value is sent to all nodes in the cluster. The
value is always a unique integer, and the body looks like this:

```json
{
  "type": "broadcast",
  "message": 1000
}
```

The response should be a simple acknowledgement:

```json
{
  "type": "broadcast_ok"
}
```

#### Read

This message requests a node to return all values it has seen so far. The body
looks like this:

```json
{
  "type": "read"
}
```

The response should have all values previously seen:

```json
{
  "type": "read_ok",
  "messages": [1, 8, 72, 85]
}
```

#### Topology

This message informs the node of its neighbours. It looks as follows:

```json
{
  "type": "topology",
  "topology": {
    "n1": ["n2", "n3"],
    "n2": ["n1"],
    "n3": ["n1"]
  }
}
```

Your response should be a simple acknowledgement:

```json
{
  "type": "topology_ok"
}
```

## Testing

Run `maelstrom` with the following command.

```sh
./maelstrom test -w broadcast --bin ~/go/bin/broadcast --time-limit 20 --rate 10 --node-count 1
```

This runs a single node for 20 seconds, sending 10 messages per second. It
simply validates that all values sent by broadcasts are returned via read.

If it works, great. That's the easy part :)
