---
title: 1sl4nds/moses
---

1sl4nds/moses is a software component of the Nostr network, a decentralized and extensible social protocol. This piece of software serves as a bridge between Nostr-compatible applications, offering both relay and client functionalities to ensure smooth communication within the network.

As an integral part of the Nostr network, 1sl4nds/moses aids in the creation, management, and communication of events between clients and relays. It provides core functionalities and supports selected NIPs (Nostr Improvement Proposals), simplifying the development process of Nostr-compatible applications and contributing to the overall growth of the network.

## Getting Started

The 1sl4nds project is structured into three main sections:

1. **Go internal executable and dependencies**: These are the core elements that run the 1sl4nds software.
2. **Go external package**: This includes essential components required for building custom clients and relays.
3. **Angular web application**: This facilitates the administration of the client and interaction with third-party clients and relays.

Here's a simple 1sl4nds usage example using the Go package to create a simple Nostr client:

```go
package main

import (
 "context"
 "flag"
 "fmt"
 "os"

 "github.com/1sl4nds/moses/client"
 "github.com/1sl4nds/moses/message"
 "github.com/1sl4nds/moses/message/requestmessage"
 "github.com/1sl4nds/moses/subscriptionid"
)

var (
 relayURL string = "wss://relay.damus.io"
 subscriptionID string = subscriptionid.New()
)

func init() {
 flag.StringVar(&relayURL, "u", relayURL, "Relay to connect to")
 flag.StringVar(&subscriptionID, "sid", subscriptionID, "Subscription ID to use for connection")
}

func main() {
 flag.Parse()

 ctx := context.Background()

 cl := client.New(&client.Options{
  ReadLimit: 2e6,
 })

 cl.Connect(ctx, relayURL)

 cl.HandleErrorFunc(func(err error) {
  fmt.Println(err.Error())
 })

 cl.HandleMessageFunc(func(msg message.Message) {
  fmt.Println(msg)
 })

 cl.SendMessage(ctx, requestmessage.New(subscriptionID, &requestmessage.Filter{}))

 if err := cl.Listen(ctx); err != nil {
  fmt.Print(err.Error())
  os.Exit(1)
 }
}
