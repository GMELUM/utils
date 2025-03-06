# My Personal Go Utils Collection

Welcome to my personal collection of Go packages, designed to streamline backend development for web-oriented projects such as VKMA, TGA, and more.

## Overview

This repository contains a set of utility packages that I have developed and refined over time, aimed at facilitating the backend development process. Each package is built with a focus on simplicity, efficiency, and ease of integration into various web-oriented projects.

## List of Utilities

Here is a list of available packages and their purposes:

- [**Package `tonsub`**](https://github.com/GMELUM/utils/tree/main/tonsub)
  - Provides features to subscribe to TON blockchain transactions, process Jetton, NFT, and TON transfer events.

- [**Package `wallet`**](https://github.com/GMELUM/utils/tree/main/wallet)
  - A package for managing transactions on the TON blockchain.

- [**Package `search`**](https://github.com/GMELUM/utils/tree/main/wallet)
  - Implementation for searching interlocutors with mandatory parameters and interests.

- [**Package `callback`**](https://github.com/GMELUM/utils/tree/main/callback)
  - Implementation of an internal callback system for dispatching notifications.

## Getting Started

To start using these packages, install the repository using `go get`:

```bash
go get github.com/gmelum/utils
```

Then import the desired package into your Go project:

```go
import (
    "github.com/gmelum/utils/{package}" // replace {package} with the specific package you want to use
)
```

## Additional Information

For detailed usage instructions and examples, please visit the specific package directories in the repository.