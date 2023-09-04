package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func contains(messages []float64, message float64) bool {
	for _, x := range messages {
		if x == message {
			return true
		}
	}
	return false
}

func main() {
	n := maelstrom.NewNode()

	var messages []float64

	var topology map[string]any

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		response := make(map[string]any)
		response["type"] = "broadcast_ok"

		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		message := body["message"].(float64)

		// Bail early if we've already seen this one!
		if contains(messages, message) {
			return n.Reply(msg, response)
		}

		// Propagate to neighbours
		messages = append(messages, message)
		for _, node := range topology[n.ID()].([]any) {
			n.Send(node.(string), body)
		}

		return n.Reply(msg, response)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any

		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		response := make(map[string]any)
		response["type"] = "read_ok"
		response["messages"] = messages

		return n.Reply(msg, response)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any

		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		topology = body["topology"].(map[string]any)

		response := make(map[string]any)
		response["type"] = "topology_ok"

		return n.Reply(msg, response)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
