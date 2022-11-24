# Graceful Shutdown

This example demonstrates graceful shutdown pattern that is closer to reality, in which a `SIGINT` instruction is sent to the application and multiple running goroutines need to gracefully exit before closing.

Using the [multi_producer_consumer](../multi_producer_consumer/) as a base, this example adds a usage of `context.WithCancel()` that allows an interrupt signal to be broadcast to multiple subscribers by receiving on `ctx.Done()`.
