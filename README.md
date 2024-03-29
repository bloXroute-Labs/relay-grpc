# Relay gRPC

This repository provides the compiled gRPC code for connecting and submitting blocks to the bloXroute Labs MEV Relays.

## Getting Started

Included in this repository are some helper files to streamline integration;

**conn.go**

- Opens and maintains connection with relay from provided host and auth token

**utils.go**

- Adds methods for converting between standard types and gRPC types

### Importing the module

Use go get to download the module:

```bash
go get github.com/bloXroute-Labs/relay-grpc
```

Import the module inside of the application you wish to add it:

```golang
import (
  grpc "github.com/bloXroute-Labs/relay-grpc"
)
```

### Opening a connection

Using the provided utility method and passing in a standard bloXroute Auth token, plus a bool to specify if gzip compression should be enabled for submitted blocks:

```golang
blockSubmissionChan, err := grpc.NewConnection("{{host}}", "{{your auth token}}", true)
```

This returns an `error` if there is an issue establishing a connection and a channel that accepts `*grpc.BlockSubmissionRequest`. You will want to store a reference to this channel somewhere that your builder will have access to during the block submission process.

If you don't want to use the provided connection method you can follow a guide online like [this one](https://grpc.io/docs/languages/go/basics/#client) and manage the connection yourself.

### Submitting blocks

Currently relays accept blocks as either JSON or SSZ encoded bytes of `spec.SubmitBlockRequest` from this module `github.com/attestantio/go-builder-client/api/spec`, which means most builders are already generating this paylaod when submitting blocks today.

The provided utility function `VersionedRequestToProtoRequest()` accepts `VersionedSubmitBlockRequest` and returns `*grpc.SubmitBlockRequest` which can then be sent to the channel returned by `grpc.NewConnection`.

```golang
body := grpc.VersionedRequestToProtoRequest(blockSubmission)
blockSubmissionChan<-body

// or

blockSubmissionChan<-grpc.VersionedRequestToProtoRequest(blockSubmission)
```

## All together

``` golang
import (
  grpc "github.com/bloXroute-Labs/relay-grpc"
)

...
// use of gzip compression is set to 'false' in this example
blockSubmissionChan, err := grpc.NewConnection("{{host}}", "{{your auth token}}", false)
if err != nil {
  log.Fatal(err)
}

...

body := grpc.VersionedRequestToProtoRequest(blockSubmission)

blockSubmissionChan<-body
```
