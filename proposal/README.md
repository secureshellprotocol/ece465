# Cloud Computing Final Proposal
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

Currently:
March 3rd: develop a roadmap on where to go, reading papers on migration & 
checkpoint/restore. This includes a brush up on some lower level operating 
systems concepts, as well as a deeper dive into the Linux scheduler.

## Prerequisites to run

This work has only been tested on Ubuntu Server 24.04 LTS, either on bare-metal
or on a virtual machine.

Vagrant must be installed, you can find the instructions
[here.](https://developer.hashicorp.com/vagrant/install)

## Steps to run

Pull down the repo and save it in a safe place
```
# apt install git
$ cd ~
$ git clone https://github.com/secureshellprotocol/ece465.git
$ cd ece465/provisioning        # this is our PWD
```

## Running the demo

`vagrant up --provision`

## Reseting the Vagrant Environment

To destroy our worker VM's, and re-provision them to the intended state:
```
$ vagrant destroy
$ vagrant up --provision
```


## Troubleshooting

```
ERROR! the role 'geerlingguy.docker' was not found in <directory list>
```

Ensure that you have run the `configure.sh` script in the repo root directory

## References:

[Jeff Geerling, "Ansible for DevOps"](https://www.ansiblefordevops.com/)

[S. Nadgowda, S. Suneja, N. Bila and C. Isci, "Voyager: Complete Container State
Migration,"](https://ieeexplore.ieee.org/abstract/document/7980161)

[George Washington University's Distributed Systems website.](https://gwdistsys20.github.io/project/#milestone-1-select-a-topic)


