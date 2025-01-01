## This is a host registry which maintains registered hosts in a cache using different strategies
This is a small implementation of a heartbeat mechanism of server agent architecture. In this implementation, I have implemented a small `TCP` based protocol called the `HB` Protocol(`HeartBeat`) [specs](HB.md)

In this implementation, I have written an a TCP server which listens on a port for TCP packets and parses the packet to a heartbeat if possible. After a successful parsing it pushes the data on to a registry cache in-memory. It also contains the client code as well for testing. We will have a separate dedicated client for testing in a real scenario in `Rust`, more on that later.


### TODOs
- [x] Create a simple tcp connection which can be used to send data to the server which resembles the registry
- [x] Create a simple tcp client which can be used to bombard the registry
- [] Create a simple caching mechanism for storing the registry in memory
- [] Create multiple strategies for the caching db
