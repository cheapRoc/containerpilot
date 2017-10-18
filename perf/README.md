This directory contains the infrastructure and scripting required to perform
performance and memory testing of ContainerPilot. The goal is to have something
we can easily setup (or leave running) to continually ensure we haven't
introduced new memory or goroutine leaks into the application.

Much of this work was done to support issue [#464][issue]. This directory
includes that knowledge into the project.

## Setup

In order to run the included tests the ContainerPilot binary needs to include
`net/http/pprof`. This standard library provides us a way to snapshot live heap
profiles. Currently, this requires an optional build time parameter for enabling
the code within the application at run time.

## Tests

- [ ] `config-test` is a soak test across various CP configuration files that
  looks at pinpointing the location within the app that could be presenting a
  leak.
- [ ] `app-test` is a soak test that mimicks an actual application, including
  some commonly found patterns used with CP. This is a broad/general test that
  attempts to discern common usage patterns. Take it with a grain of salt.
- [ ] `reload-test` is a soak test that constantly reloads CP over and over
  looking for any increasing data points.
- [ ] `signal-test` is a soak test that constantly barrages CP with UNIX
  signals, hoping to create a leak of some kind.

## Utilities

- `bin/heaps-to-csv.py` process Go heap profiles into a CSV report. Useful for
  graphing output memory stats.
- `bin/profile-heap.sh` grabs a heap profile from a running ContainerPilot
  instance.
- `bin/dump-heap.sh` grabs an entire heap dump from a running ContainerPilot
  instance. Only useful in extreme cases where the profile is not enough
  information.
- `bin/snapshot-pmap.sh` records a snapshot of pmap.

## Environments

- `Dockerfile` describes a container which can be used to run ContainerPilot. It
  includes all of the aforementioned scripts and tests.
- `docker-host.tf` describes an Ubuntu Linux host instance setup to run native
  Docker.

## Links

- https://medium.com/golangspec/goroutine-leak-400063aef468

[issue]: https://github.com/joyent/containerpilot/issues/464
