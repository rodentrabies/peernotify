# Peernotify: pseudonymous notification service


## Introduction
In the last few years, there is a lot of movement towards decentralization
of the common services and tools for economic activity. Now, few years after
distributed hash tables decentralized file sharing and bitcoin decentralized
money, there are technologies that attempt to decentralize the biggest
application of the Internet itself: the Web. One of the most interesting
projects of that kind is [*IPFS*][ipfs] [^1], which combines distributed hash tables
(*DHT*) like [*Kademlia*][kademlia] [^2] and content-addressed file system like
[*Git*][git] [^3] to create a global hypermedia environment that is based on
what it contains and not on who owns it.

An interesting consequence of using *DHT* as transport is persistency of the
information resources within the platform, which allows for *offline services
model*: owner of the resource uploads it to the DHT where clients can get it
without connecting to the owner directly. Such model greatly reduces
maintenance costs, as offline services only have to be connected to the
network for an amount of time sufficient to upload new data, thus the name.
What is even more important is that the owner of such service can develop a
pseudonymous identity, which hides his personal information while allowing him
to be a trusted entity on the network. Great example of such system is the
[*OpenBazaar 2.0*][ob] [^4] platform, which is basically a network of pseudonymous
offline (well, not necessarily) stores.

The problem with such design is that it lacks interactivity, which is crucial
for most of traditional web-services. If the service is offline, clients cannot
directly contact it, and that makes such a service basically useless. There are
two ways to solve this problem:
- provide contact information like email or phone number alongside service's
identity information;
- "roll your own messaging".

The first one goes against the very idea of pseudonymous offline services by
binding personal contact information such as email or phone number to your
pseudonym within the network. The second one is, well, not exaclty a convenient
means of communication, as clients need to have instances of messaging apps etc,
while email/SMS is what we call "at hand" communication - everyone carries it
on his smartphone.


## Peernotify service

**Peernotify** is a service that provides pseudonymous notifications mechanism
on top of offline services model. It combines both ways highlighted above into
a intermediary layer between handy, personal communication data and
pseudonymous identities within a peer-to-peer service network. Service allows
each client to register sets of
*contacts*, where:
- *contact* is a pair *(protocol, address)*;
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
- `forward` - send notification to the generator of the one-time token,
described below.

`register` and `verify` functions are meant for clients and represent two phases
of registration process, while `forward` is an endpoint for pers

#### Registration
To register, user must submit his contact data to the service. Contact data is
simply a list of tuples of the form `(<medium>, <address>)`, where `<medium>`
is a short string that describes some common communication protocol, like SMTP
or Signal, and `<address>` is an abstract identifier of the user within that
protocol, like actual email address or phone number, for SMTP or Signal
respectively.

Example of registration request data in JSON format:
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

#### Verification
Verification process must be performed to confirm that the user has access
to each personal identifier withing the specified communication protocol
(it is supposed that immediate access to the communication system under that
identifier is a sufficient confirmation). Verification is done via sending
confirmation link to each address and requiring user to "click on dat link".
The link itself consists of a randomly generated 512-bit number in a *base58*
encoding. Temporary data storage maps this number into contact data
and once HTTP request is performed, data at that key is moved to permanent
storage and at this point it becomes valid and can be used to forward ping
messages.

#### Forwarding
Forwarding is done via submitting a pair `(<token>, <message>)` to the service.
The token is verified according to the protocol described below and if the
verification is successful, services dispatches the message to each address
known for the client who generated the token.

Example of forwarding request data in JSON format:
```json
{
    "token": "C9W3LaZPskHKNqLncZc9ZaS8By4a82kM55U83rMzFs3kXgMP7oBu4gd",
    "message": "Anonymous user says \"where's my money???\""
}
```


## Peernotify protocol
TODO

## References
[^1]: https://ipfs.io/
[^2]: https://pdos.csail.mit.edu/~petar/papers/maymounkov-kademlia-lncs.pdf
[^3]: https://git-scm.com/
[^4]: https://www.openbazaar.org/

[ipfs]: https://ipfs.io/
[kademlia]: https://pdos.csail.mit.edu/~petar/papers/maymounkov-kademlia-lncs.pdf
[git]: https://git-scm.com/
[ob]: https://www.openbazaar.org/
