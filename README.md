# Cloud Computing Proposal
James Ryan, advised by Prof. Rob Marano

## Course Goals

Our course goals involve creating an MVP which follows our four design goals:

1. Make Resources Easily Accessible
2. Mask the fact that these resources are distributed over a network
3. Build each component to be easily interchangable
4. Have the ability to scale based on usage

## Proposal

My main proposal is to design a Distributed information system (see Distributed
Systems, ch 1.3.2). I would use this framework to design a communication/social
network service similar to [Matrix](https://matrix.org/), primarily focusing on
secure, consistent communication between clients, allowing for any user to host
their own server/instance, allowing for seamless communication between clients
on any home server. This is commonly known as a [Federated service](https://en.wikipedia.org/wiki/Federated_architecture).

So, the service aims to be a seamless instant messenger which users can interact
with on any client, as long as it adheres to the API. And, users can host their
own home-server for clients to connect to & interact with the overall network.
