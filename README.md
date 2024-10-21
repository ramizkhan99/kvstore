# Distributed Key Value

A simple distributed Key-Value store implementation using GO

The storage layer uses CockroachDB's implementation of Rocksdb, [pebble](https://github.com/cockroachdb/pebble).

Clients can connect using gRPC.

As of now, the code isn't very clean, I will keep updating and cleaning up as I learn more about Golang as I am new to this language.

> Currently I have implemented only for single server and am working on the distributed application. As of now, the servers would be communicating using gRPC and for fault-tolerance Raft would be used as the consensus algorithm. As I am still studying and applying them at the same time it might take some time to complete the implementation.

