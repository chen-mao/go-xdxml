package main

/*
#cgo LDFLAGS: -L/opt/xdxgpu/lib/x86_64-linux-gnu -lxdxgpu-ml -Wl,--unresolved-symbols=ignore-in-object-files

#include <stdlib.h>
#include <dlfcn.h>

#define XDXML_DEVICE_NAME_BUFFER_SIZE 256

typedef enum {
    XDXML_SUCCESS = 0,
    XDXML_FAILURE = 1
} xdxml_return_t;

typedef void* xdx_device_t;

xdxml_return_t xdxml_init(void);
xdxml_return_t xdxml_device_get_count(unsigned int*);
xdxml_return_t xdxml_device_get_handle_by_index(unsigned int, xdx_device_t *);
xdxml_return_t xdxml_device_get_product_name(xdx_device_t, char*);
xdxml_return_t xdxml_device_get_minor_number(xdx_device_t, int *);
xdxml_return_t xdxml_shutdown(void);
*/
import "C"
import (
	"fmt"
)

func main() {
	var deviceCount C.uint
	var result C.xdxml_return_t
	var minor_number C.int
	name := make([]C.char, C.XDXML_DEVICE_NAME_BUFFER_SIZE)

	result = C.xdxml_init()
	if result != C.XDXML_SUCCESS {
		fmt.Println("Failed to initialize.")
		return
	}

	result = C.xdxml_device_get_count(&deviceCount)
	if result != C.XDXML_SUCCESS {
		fmt.Println("Failed to get device count.")
		return
	}

	fmt.Printf("\nmain:device count is: %d\n", deviceCount)
	fmt.Println("Listing devices:")

	for i := C.uint(0); i < deviceCount; i++ {
		var device C.xdx_device_t
		result = C.xdxml_device_get_handle_by_index(i, &device)
		if result != C.XDXML_SUCCESS {
			fmt.Println("Failed to get handle.")
			return
		}
		result = C.xdxml_device_get_product_name(device, &name[0])
		if result != C.XDXML_SUCCESS {
			fmt.Println("Failed to get name of device.")
			return
		}
		fmt.Printf("\nproduct_name is : %s\n", C.GoString(&name[0]))
		result = C.xdxml_device_get_minor_number(device, &minor_number)
		if result != C.XDXML_SUCCESS {
			fmt.Println("Failed to get minor_number of device.")
			return
		}
		fmt.Printf("\nminor_number is : %d\n", minor_number)
	}

	result = C.xdxml_shutdown()
	if result != C.XDXML_SUCCESS {
		fmt.Println("Failed to shutdown.")
		return
	}
}
