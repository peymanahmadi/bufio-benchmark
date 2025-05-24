package bufio_bench

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	fileName = "10mb.log"
	fileSize = 10 * 1024 * 1024 // 10 MB
)

// TestMain sets up the 10MB file before running benchmarks.
func TestMain(m *testing.M) {
	// Create the file if it does not exist
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		buffer := make([]byte, 1024)
		written := 0
		for written < fileSize {
			// Fill buffer with random data
			r.Read(buffer)
			toWrite := min(fileSize - written, len(buffer))
			// Write to file
			if _, err := f.Write(buffer[:toWrite]); err != nil {
				panic(err)
			}
			written += toWrite
		}
	}

	// Run benchmarks
	os.Exit(m.Run())
}

// BenchmarkUnbufferedRead reads the file one byte at a time without buffering.
func BenchmarkUnbufferedRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			b.Fatalf("failed to open file: %v", err)
		}

		buf := make([]byte, 1)
		for {
			if _, err := file.Read(buf); err != nil {
				if err == io.EOF {
					break
				}
				b.Fatalf("read error: %v", err)
			}
		}
		file.Close()
	}
}

// BenchmarkBufferedRead reads the file one byte at a time using bufio.Reader.
func BenchmarkBufferedRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			b.Fatalf("failed to open file: %v", err)
		}

		reader := bufio.NewReader(file)
		buf := make([]byte, 1)
		for {
			if _, err := reader.Read(buf); err != nil {
				if err == io.EOF {
					break
				}
				b.Fatalf("buffered read error: %v", err)
			}
		}
		file.Close()
	}
}

// BenchmarkUnbufferedRead4K reads the file with 4KB chunks without buffering.
func BenchmarkUnbufferedRead4K(b *testing.B) {
	buf := make([]byte, 4096)
	for i := 0; i < b.N; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			b.Fatalf("failed to open file: %v", err)
		}

		for {
			_, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				b.Fatalf("read error: %v", err)
			}
		}
		file.Close()
	}
}

// BenchmarkBufferedRead4K reads the file with 4KB chunks using bufio.Reader.
func BenchmarkBufferedRead4K(b *testing.B) {
	buf := make([]byte, 4096)
	for i := 0; i < b.N; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			b.Fatalf("failed to open file: %v", err)
		}

		reader := bufio.NewReader(file)
		for {
			_, err := reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				b.Fatalf("buffered read error: %v", err)
			}
		}
		file.Close()
	}
}