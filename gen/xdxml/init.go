package xdxml

import (
	"fmt"

	"github.com/chen-mao/go-xdxml/pkg/dl"
)

import "C"

const (
	xdxmlLibraryName      = "libxdxgpu-ml.so"
	xdxmlLibraryLoadFlags = dl.RTLD_LAZY | dl.RTLD_GLOBAL
)

var xdxml *dl.DynamicLibrary

// xdml_init()
func Init() Return {
	lib := dl.New(xdxmlLibraryName, xdxmlLibraryLoadFlags)
	err := lib.Open()
	if err != nil {
		return ERROR_LIBRARY_NOT_FOUND
	}
	xdxml = lib

	return xdxml_init()
}

// xdxml_shutdown()
func Shutdown() Return {
	ret := xdxml_shutdown()
	if ret != SUCCESS {
		return ret
	}

	err := xdxml.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing %s: %v", xdxmlLibraryName, err))
	}

	return ret
}
