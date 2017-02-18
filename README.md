# Peernotify: pseudonymous notification service

## Introduction
In the last few years, there is a lot of movement towards decentralization
of the Internet services. Implementations of the ideas from 90's cypherpunk 
movement, namely persistent, immutable, content-addressed information 
resources, zero-trust economic activity tools like electronic cash and 
peer-to-peer marketplaces etc., have unexpectedly emerged within the 
cryptographic community of the last decade. Two of most prominent technologies
of that kind are distributed hash tables which are already widely adopted for
file sharing in the *BitTorrent* protocol, and more recent cryptocurrencies
like Bitcoin, which builds on ideas of Wei Dai, Adam Back and David Chaum and
is considered one of the greatest inventions since the Internet itself.

Mixing this technologies spawns a new wave of Internet services that 
eliminate trust in many areas of our life. We no longer have to rely on 
trustworthiness of a central authority to handle economic activities such 
as financial operations and trading, because there are cryptographic protocols 
for doing it peer-to-peer, without external control, while providing much
higher levels of security and confidentiality. Peer-to-peer hypermedia
protocols like [*IPFS*][1], [*Blockstack*][2], etc allow service providers
to securely and persistently distribute content in the network, effectively
no longer requiring them to be "online" all the time.

The problem inherent here is that possibility to be offline most of the time,
as your data resides in persistent global hypermedia environment, removes
interactivity from it. Assuming pseudonymous identities, it is impossible to 
reach the originator of the content, as mapping from his unique ID to his
real identity is valid only within the peer-to-peer system at hand. 
**Peernotify** protocol proposes a solution based on one-time tokens that can
be non-interactively generated by clients and shared with their peers, who in
turn can use them to both pay the service and send ping messages.

## Solution

## Peernotify protocol

## Analysis

## Conclusion

## References


[1]: https://ipfs.io/
[2]: https://blockstack.org/
[3]: https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
