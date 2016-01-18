#Go Channel Compendium
**Interesting ways of using Go channels by John Graham-Cumming**

---

## Overview

This is a code archive taken from the slides of John Graham-Cumming's
presentation of 'A Channel Compendium' given at GopherCon 2014.

## Video

https://www.youtube.com/watch?v=SmoM1InWXr0

## Contents

* Signalling
  * [Wait for an event](https://github.com/nomad-software/go-channel-compendium/blob/master/signalling/wait-for-an-event/main.go)
  * [Coordinate multiple goroutines](https://github.com/nomad-software/go-channel-compendium/blob/master/signalling/coordinate-multiple-goroutines/main.go)
  * [Coordinated termination of workers](https://github.com/nomad-software/go-channel-compendium/blob/master/signalling/terminate-workers/main.go)
  * [Verify termination of workers](https://github.com/nomad-software/go-channel-compendium/blob/master/signalling/verify-terminate-workers/main.go)
* Hide state
  * [Unique ID service](https://github.com/nomad-software/go-channel-compendium/blob/master/hide-state/unique-id-service/main.go)
  * [Memory recycler](https://github.com/nomad-software/go-channel-compendium/blob/master/hide-state/memory-recycler/main.go)
  * [Capped memory recycler](https://github.com/nomad-software/go-channel-compendium/blob/master/hide-state/capped-memory-recycler/main.go)
* Nil channels
  * [Disable receiving case statements](https://github.com/nomad-software/go-channel-compendium/blob/master/nil-channels/disable-receiving-case-statements/main.go)
  * [Disable sending case statements](https://github.com/nomad-software/go-channel-compendium/blob/master/nil-channels/disable-sending-case-statements/main.go)
* Timers
  * [Timeout](https://github.com/nomad-software/go-channel-compendium/blob/master/timers/timeout/main.go)
  * [Heartbeat](https://github.com/nomad-software/go-channel-compendium/blob/master/timers/heartbeat/main.go)
* Examples
  * [Network multiplexer](https://github.com/nomad-software/go-channel-compendium/blob/master/examples/network-multiplexer/main.go)
