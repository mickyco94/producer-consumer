# Single Producer/Consumer

This example demonstrates the simplest concurrent model with only 3 goroutines, our main goroutine and one each for the data consumer and producer respectively.

Here the main goroutine is blocked by a simple `chan bool` that is kept open by the consuming goroutine until all messages in the `jobs` channel have been received.
