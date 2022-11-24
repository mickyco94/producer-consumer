# Multi Producer / Single Consumer

This example demonstrates the case where we have multiple message producers but only one message consumer.

The main goroutine is blocked while the producers are still processing using a `WaitGroup`, which can be thought of as simply a counter, `.Add(1)` increments and `.Done()` decrements. `Wait()` will not yield until this "counter" is zero.
