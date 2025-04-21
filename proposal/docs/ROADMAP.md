# Timeline for April-May:

## Starting with a reproducible build environment.

I am selecting Ubuntu 24.04.02 LTS as the only supported environment for running
a demo of the project. Included will be an [autoinstall
configutaion](https://canonical-subiquity.readthedocs-hosted.com/en/latest/howto/autoinstall-quickstart.html#autoinstall-quick-start)
which can be used to kickstart an image with the required libraries.

That will be used in tandem with an ansible script to get the rest of the
requisite software in place, such as podman configuration.

Instructions will be included under a `docs/` directory

## The demo

My goal is to design a re-implementation of the 
[Voyager](https://ieeexplore.ieee.org/abstract/document/7980161) paper's method
of moving a container from one machine to another. 

### Goal 1: basic election

The main goal is to have a basic election between nodes, where one is the
controller, and 

Some choices:

- Choice 1
-- The demo will at first target being able to select a controller through some
election mechanism. I will likely be using OpenAFS's election mechanism for
selecting a leader, where the server with the lowest-valued IP address is
selected as a leader. This is closer to a distributed model

- Choice 2
-- We act as a service, and have one controller which we connect to by default.
This is given a list of nodes to worry about. This would be more straightforward
to implement

### Goal 2: basic container checkpoint-restore onto another machine

The demo should be able to checkpoint a container from one VM, dump its memory,
and start it again. First i need to implement this
[demo](https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/8/html/building_running_and_managing_containers/assembly_creating-and-restoring-container-checkpoints#proc_creating-and-restoring-a-container-checkpoint-locally_assembly_creating-and-restoring-container-checkpoints)
and try to get it to work between Vagrant vms.

### Goal 3: tie these together with a userspace tool

Any bash script would do -- just something to access the controller and tell it
to move one vm to another.

# Future work beyond the course

Some future work could involve making a demo which uses an algorithm to move
containers based on the resources being used by each monitored VM. 
