# Single Process Graceful Shutdown

This example demonstrates implementing a graceful shutdown on a single process. It is implemented here only for the consumer portion of the application, using a channel passed to the consumer that is listening for the `SIGINT` OS signal.
