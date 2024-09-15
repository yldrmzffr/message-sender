# Notification Module

This module provides a flexible and extensible solution for sending SMS messages through different providers. It is built using the Strategy design pattern.

## Table of Contents

1. [Features](#features)
2. [Architectural Design](#architectural-design)
3. [Use of Strategy Pattern](#use-of-strategy-pattern)
4. [Installation](#installation)
5. [Usage](#usage)
6. [Providers](#providers)
7. [Error Handling](#error-handling)
8. [Development](#development)
9. [Testing](#testing)
10. [Performance Evaluation](#performance-evaluation)
11. [License](#license)

## Features

- Multi-provider support (currently GCP and Mock)
- Easily extensible structure using the Strategy pattern
- Centralized error handling
- Abstracted provider interface

## Architectural Design

The module consists of the following main components:

1. `Provider` interface: Defines the methods that all providers must implement.
2. Concrete providers: Structures like `GCPSmsProvider` and `MockProvider` that implement the `Provider` interface.
3. `Notification` structure: Holds the selected provider.
4. Factory function: `NewService` creates the desired provider.

## Use of Strategy Pattern

The Strategy pattern is used in this module for the following purposes:

1. **Flexibility**: Easily switch between different SMS sending strategies (providers).
2. **Extensibility**: Facilitate the integration of new providers into the system.
3. **Reduction of Dependencies**: Separate provider implementations from client code.
4. **Testing**: Enable easy testability using mock providers.

This pattern is implemented through the `Provider` interface and concrete classes that implement this interface.

## Installation

To include the module in your project:

```go
import "path/to/notification"
```

## Usage

1. Provider configuration:

```go
provider, err := notification.ConfigureNotificationModule("gcp")
if err != nil {
    // Error handling
}
```

2. Sending a message:

```go
response, err := provider.Send("1234567890", "Hello, world!")
if err != nil {
    // Error handling
}
```

## Providers

### GCP SMS Provider

- Sends SMS using GCP Cloud Functions.
- Makes a real HTTP request.
- Suitable for production use.

### Mock Provider

- Doesn't send a real SMS, just simulates it.
- Provides a fast and resource-efficient testing environment.
- Recommended for development and testing phases.

## Error Handling

The module uses custom error types:

- `ErrProviderNotFoundResponse`: When the specified provider is not found.
- `ErrSendingMessageResponse`: When an error occurs during message sending.

These errors are created using the `apperrors` package and include error codes and HTTP status codes.

## Development

To add a new provider:

1. Create a new file (e.g., `newprovider.go`)
2. Create a structure that implements the `Provider` interface
3. Add the new provider to the `NewService` function
4. Define new error types if necessary

Example:

```go
type NewProvider struct {
    name string
}

func NewNewProvider() *NewProvider {
    return &NewProvider{
        name: "NewProvider",
    }
}

func (p *NewProvider) Send(to string, content string) (*ProviderSuccessResponse, error) {
    // Implementation
}

// Add to NewService function
case "new":
    provider = NewNewProvider()
```

