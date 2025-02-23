# Ind Study ECE 465 - HW 2
## Due February 24, 2025 11:59 PM

# How to run and build

First, make sure the Go toolchain is installed on your system.

Refer to the [Go installation docs](https://go.dev/doc/install) for instructions 
on how to build for your system.

For Fedora Linux:
```
$ sudo dnf install golang
```

For Ubuntu Linux:
```
$ sudo apt install golang
```

To run:

```
$ go build
$ ./hw2
```

# Instructions

You are to provide a solution to the Dining Philosophers problem. It's a classic illustration of concurrency and some of its fundamental challenges like deadlock and starvation.


Imagine a scenario with five philosophers seated around a circular table. Between each pair of adjacent philosophers lies a single chopstick. Each philosopher alternates between thinking and eating. To eat, a philosopher needs two chopsticks – the ones to their immediate left and right.


The problem arises because if all philosophers simultaneously decide to eat and each grabs their left chopstick, they'll all be stuck. No one can eat because everyone is waiting for the chopstick held by their neighbor. This is a classic deadlock situation. Everyone is blocked, waiting for a resource that will never become available.


Here's a breakdown of the key elements and challenges:

    Philosophers: Represent processes or threads competing for resources.
    Chopsticks: Represent shared resources. They are a limited quantity.
    Eating: Represents a process holding multiple resources (two chopsticks) to perform some work.
    Thinking: Represents a process releasing resources and not needing them.

The Dining Philosophers problem highlights several important concurrency concepts:

    Deadlock: As explained above, a situation where two or more processes are blocked indefinitely, waiting for each other to release resources.
    Starvation: A situation where a process is repeatedly denied access to a resource, even though the resource is periodically available. For example, one philosopher might consistently be unlucky and never get both chopsticks, even if deadlock is avoided.
    Mutual Exclusion: Only one philosopher can use a chopstick at a time. This is a fundamental requirement for shared resources.

Several solutions exist to address the Dining Philosophers problem. Which one will you use to solve this.

Your solution will run only on one single computer equipped with a multi-processor, e.g., multi-core x86_64/amd64 or arm64 running either Windows, MacOS, or Linux.
