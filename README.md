# go-xdxml
The xdxml is go package for the XDXCT Management Library API.

## Building and Test
This only prerequisite is working installation of c-for-go from [https://github.com/xlab/c-for-go](https://github.com/xlab/c-for-go).

Build the code
```shell
$ export CGO_LDFLAGS="-ldl"
$ make clean
$ make
```

Test the code
```shell
go run examples/device/main.go
```

