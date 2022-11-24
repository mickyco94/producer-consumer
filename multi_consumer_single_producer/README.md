# Single Producer / Multi Consumer

This example demonstrates the case where we have a single message producer, with multiple consumers receiving.

Here we do not need to use a `WaitGroup` to keep exits in sync, all consumers will block as long as the consuming channel is open. Only one of the consumers need to indicate to the main goroutine that consuming has been completed, allowing the application to exit.
