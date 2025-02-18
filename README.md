# Cloud Computing @ Cooper Union
James Ryan, advised by Prof. Rob Marano

## Course Goals

Our course goals involve creating an MVP which follows our four design goals:

1. Make Resources Easily Accessible
2. Mask the fact that these resources are distributed over a network
3. Build each component to be easily interchangable
4. Have the ability to scale based on usage

## Proposal

As a rough cut, 
I'd like to investigate current approaches to container migration techniques 
across distributed systems, and research what heuristics matter when migrating
resources between nodes. I plan to utilize [QEMU](https://www.qemu.org/), [CRIU](https://criu.org/Main_Page),
as well as some original code in (probably) Python, I'll reimplement techniques
 reviewed in past & current publications, looking to make optimizations where I
 can.

Primarily, these methods will center around the course goals of transparency 
and scalability of distributed resources, attempting to make this system work 
independent of the actual software on each container, and properly mask the 
migrations occuring on the overall container network

This topic will evolve over time.

Currently:
March 3rd: develop a roadmap on where to go, reading papers on migration & 
checkpoint/restore. This includes a brush up on some lower level operating 
systems concepts, as well as a deeper dive into the Linux scheduler.

E.g [https://ir.lib.uwo.ca/etd/10358/](https://ir.lib.uwo.ca/etd/10358/), [https://ieeexplore.ieee.org/abstract/document/7980161](https://ieeexplore.ieee.org/abstract/document/7980161)

I took a lot of inspiration first by getting direction from the broad list of 
research topics presented on [George Washington University's Distributed Systems website.](https://gwdistsys20.github.io/project/#milestone-1-select-a-topic),
 and then reading the various papers linked on that site. 
