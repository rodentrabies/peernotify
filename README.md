# Peernotify: pseudonymous notification service

## Introduction
In the last few years, there is a lot of movement towards decentralization
of the common services and tools for economic activity. Now, few years after
distributed hash tables decentralized file sharing and bitcoin decentralized 
money, there are technologies that attempt to decentralize the biggest 
application of the Internet itself: the Web. One of the most interesting 
projects of that kind is [*IPFS*][1] [1], which combines distributed hash tables
(DHTs) like [*Kademlia*][2] [2] and content-addressed file system like 
[*Git*][3] [3] create a global hypermedia environment that works with what it
contains, not who owns it.

An interesting consequence of using DHT as transport is persistency of the 
information resources within the platform, which allows for *offline services
model*: owner of the resource uploads it to the DHT where clients can get it
without connecting to the owner directly. Such model greatly reduces maintenance
costs, as offline services only have to be connected to the network for an 
amount of time sufficient to upload new data, thus the name.


## Peernotify service

**Peernotify** is a service that provides pseudonymous notifications mechanism
for offline services owners and clients. It allows to register sets of contacts, 
where:
- contact is a pair *(protocol, address)*;
- *protocol* is an identifier of the communication protocol to be used;
- *address* is a unique identifier of an entity withing the communication system
built on top of the given protocol.
After registration is confirmed via verification link, client receives a root 
key, that can be used to generate one-time tokens, each of which can be used to 
notify the client of some event or forward a message while using convenient 
ways of communication like email or SMS. Thus client provides a way for a third
party to contact him at his personal environment in case his service is offline 
while keeping his personal contact data private.

**Peernotify** API consists of three main functions:
- `register` - upload contact data to the service;
- `verify` - confirm access to the accounts represented by contact data;
- `forward` - send notification to the generator of the one-time token, described
below.

Register and Verify functions are meant for client and represent two phases
of registration process, while Forward is an endpoint for pers

### Registration
To register, user must submit his *contact* data to the service. This can be
done via JSON API or via webform. Contact data is simply a list of tuples of
the form `(<medium>, <address>)`, where `<medium>` is a short string that 
describes some common communication protocol, like SMTP or Signal, and 
`<address>` is an abstract identifier of the user within that protocol, like
actual email address or phone number, for SMTP or Signal respectively.

Example of contact data in JSON format:
```json
{
    "methods": [
        {
            "protocol":   "smtp",
            "address": "me.here@example.net"
        },
        {
            "id":   "signal",
            "addr": "+380930000000"
        }
    ]
}
```

After contact data is submitted, it is placed in a temporary data storage and
verification requests are sent to each of the addresses according to the 
specified protocols. From now on contact data is invalid until verified.

### Verification
Verification process must be performed by the user to confirm that for each 
identifier and communication protocol specified, it truly belongs to him 
(we suppose that immediate access to the communication system under that 
identifier is a sufficient confirmation). Verification is done via requesting
the user to perform an HTTP request (with GET method) to the specific path at
the service's verification endpoint. Path consists of verification dispatch
subpath and a string that encodes a randomly generated 256-bit key in a 
*base58* encoding. Temporary data storage maps this key to the contact data
and once HTTP request is performed, data at that key is moved to permanent
storage and at this point it becomes valid and can be used to forward ping
messages.


### Forwarding
TODO


## Analysis
TODO

## Conclusion
TODO

## References
TODO


[1]: https://ipfs.io/
[2]: https://blockstack.org/
