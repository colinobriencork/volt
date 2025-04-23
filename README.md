# Volt

Volt is a lightweight distributed data processing framework written in Go.

## Features

- Resilient Distributed Datasets (RDD) implementation
- Lazy evaluation of transformations
- Fault-tolerant processing
- Simple, efficient API
- Pure Go implementation

## Quick Start

```go
package main

import "github.com/colinobriencork/volt"

func main() {
    // Create a new RDD
    data := []interface{}{1, 2, 3, 4, 5}
    rdd := volt.NewRDD(data)

    // Transform data
    result := rdd.
        Map(func(x interface{}) interface{} {
            return x.(int) * 2
        }).
        Filter(func(x interface{}) bool {
            return x.(int) > 5
        }).
        Collect()
}
```

## Installation

```bash
go get github.com/yourusername/volt
```

## Developing

```bash
go mod tidy
```
- It ensures that your go.mod file includes all the necessary dependencies for your project.
- It removes any unused dependencies from go.mod.
- It generates or updates the go.sum file with checksums for all required modules (both direct and indirect).

```bash
go run ./cmd/volt
```
- For development, this is the most common and efficient approach for running the project.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Apache License 2.0