package main

import ()

// This will be an orchestrator. It runs API routes that can receive events
// Then it sends it to all workers that are subscribed to this event
// For now we'll consider that they're subscribed to all events. Later we'll add filtering
func main() {
	Start()
}
