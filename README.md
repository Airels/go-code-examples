# Code examples to get started with the [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) language

The sole purpose of this repository is to provide some help for beginners with [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) language through some examples.
Further documentation available on web to be able to really master this language.

## What's Go ?

Often referred as Golang, Go is a statically typed, compiled programming language, imperative and oriented object.
Designed at Google in 2009, its syntax is similar to C, and goal of it is to likely compete C in run-time efficiency, and being readable and usable easily, with high-performances.
Go targeted originally core system components, then being used in all various domains, even used as script language due to its quick compilation.

## List of given files :

- *hello.go* : A simple Hello world which gets the name of the user session (on linux or windows)
- *isPrimeNumber.go* : A function declaration who calculates if given number n is prime number or not (return boolean)
- *findPrimeNumbers.go* : A program using concurrent functions to get prime numbers between two numbers. When function instance ends its calculations, will return a boolean in a [Go channel](https://tour.golang.org/concurrency/2)
- *user.go* : An implementation of structures in Go
- *net.go* : A TCP server who sends Hello World at connection and copy all received data into standard output.

## How to execute ? (minimal help)

1) Install golang compiler. On Debian-based GNU/Linux distributions :
   `sudo apt install golang-go`

2) To only run go file :
   `go run [yourfile]`

3) To build executable :
   `go build [yourfile]`

3-bis) To build executable for Windows (example) :
`env GOOS=windows GOARCH=amd64 go build hello.go`



Further information about compilation explained [here](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)
