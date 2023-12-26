# decode-json-from-same-reader-2-times

Example of reading and decoding JSON from the same `io.Reader` 2 times.

## How to Read and Decode JSON data from the Same io.Reader 2 Times
* Use [io.TeeReader](https://pkg.go.dev/io#TeeReader) to Duplicate the IO Stream

## Example
* Laptop is an Interface
* MacBook and WindowsLaptop Types Implement the Laptop Interface
* MacBook or WindowsLaptop Contains Specified Field(e.g. `WithAppleSilicon` for MacBook Type)
* LoadFromJSON creates a Laptop Interface(`*MacBook` or `*WindowsLaptop`) by Loading the JSON Data

  It needs the `io.Reader` can be read 2 times.
  * 1st Read and Decode: Decode the JSON Data into a Base struct value which contains the Manufacturer
  * 2nd Read and Decode: Decode the JSON Data into MacBook or WindowsLaptop According to the Manufacturer

## How to Run
```
go run main.go

// Output:
========== Laptop: 0 ==========
MacBook Pro 160in.(M3 Max)
with Apple Silicon: true
===============================

========== Laptop: 1 ==========
Lenovo Y7000
Office Pre-Installed: true
===============================
```

## References
* [How to read multiple times from same io.Reader](https://stackoverflow.com/questions/39791021/how-to-read-multiple-times-from-same-io-reader)
