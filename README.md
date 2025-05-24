# File I/O Benchmarking in Go
This repository contains Go code for benchmarking file I/O operations, comparing unbuffered and buffered reading methods. The code is designed to measure the performance of reading a 10MB file using different approaches, specifically reading one byte at a time and reading in 4KB chunks, both with and without buffering using the `bufio` package.

The code is referenced in the following article on Dev.to: [Introduction to `bufio` in Go: Why Buffered I/O Matters](https://dev.to/peymanahmadi/introduction-to-bufio-in-go-why-buffered-io-matters-58n).

## Overview
The benchmarks create a 10MB file (`10mb.log`) filled with random data and then measure the performance of reading this file in four different ways:

- **Unbuffered Read (1 byte at a time)**: Reads the file directly using `os.File.Read` with a 1-byte buffer.
- **Buffered Read (1 byte at a time)**: Reads the file using `bufio.Reader` with a 1-byte buffer.
- **Unbuffered Read (4KB chunks)**: Reads the file directly using `os.File.Read` with a 4KB buffer.
- **Buffered Read (4KB chunks)**: Reads the file using `bufio.Reader` with a 4KB buffer.

The goal is to demonstrate the performance impact of using buffered I/O (`bufio.Reader`) compared to unbuffered I/O and the effect of different buffer sizes.

## Prerequisites
- **Go**: Ensure you have Go installed (version 1.13 or later recommended).
- A system with sufficient disk space to create a 10MB file for testing.

## Installation
- Clone this repository or copy the code into a Go package.
- Ensure the code is in a directory named `bufio_bench` (or adjust the package name accordingly).
- Run the benchmarks using the instructions below.

## Usage
To run the benchmarks, navigate to the directory containing the code and execute:

```
bash

go test -bench=.
```


This will:
- Generate a 10MB file named `10mb.log` filled with random data (if it doesn't already exist).
- Run the four benchmark tests (`BenchmarkUnbufferedRead`, `BenchmarkBufferedRead`, `BenchmarkUnbufferedRead4K`, `BenchmarkBufferedRead4K`).
- Output the benchmark results, showing the time taken for each approach.

## Example Output
The output will look something like this (results vary based on hardware and system load):
```
goos: linux
goarch: amd64
pkg: bufio_bench
BenchmarkUnbufferedRead-8        1000             1234567 ns/op
BenchmarkBufferedRead-8          2000              654321 ns/op
BenchmarkUnbufferedRead4K-8      5000              234567 ns/op
BenchmarkBufferedRead4K-8        8000              123456 ns/op
PASS
ok      bufio_bench     6.789s
```

- The numbers indicate the time taken per operation (in nanoseconds) for each benchmark.
- Lower times indicate better performance.

## Code Structure
- **TestMain**: Sets up the 10MB test file (`10mb.log`) with random data before running the benchmarks.
- **BenchmarkUnbufferedRead**: Measures unbuffered reading of the file one byte at a time.
- **BenchmarkBufferedRead**: Measures buffered reading of the file one byte at a time using `bufio.Reader`.
- **BenchmarkUnbufferedRead4K**: Measures unbuffered reading of the file in 4KB chunks.
- **BenchmarkBufferedRead4K**: Measures buffered reading of the file in 4KB chunks using `bufio.Reader`.

## Notes
- The file `10mb.log` is created only once and reused for subsequent runs unless deleted.
- The benchmarks use a fixed file size of 10MB, defined by the `fileSize` constant.
- The `min` function used in TestMain is assumed to be defined elsewhere (e.g., `func min(a, b int) int { if a < b { return a } return b }`).

## Contributing
Feel free to fork this repository, make improvements, or add additional benchmarks. Pull requests and issues are welcome!

## References
- Go `bufio` package documentation
- Go `testing` package documentation
- Dev.to article: [Introduction to `bufio` in Go: Why Buffered I/O Matters](https://dev.to/peymanahmadi/introduction-to-bufio-in-go-why-buffered-io-matters-58n)