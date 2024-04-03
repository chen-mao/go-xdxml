package main

import (
	"fmt"
	"log"

	"github.com/chen-mao/go-xdxml/pkg/xdxml"
)

func main() {

	ret := xdxml.Init()
	if ret != xdxml.SUCCESS {
		log.Fatalf("Unable to initialize XDXML: %v", ret)
		return
	}
	defer func() {
		ret := xdxml.Shutdown()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to shutdown XDXML: %v", ret)
		}
	}()

	count, ret := xdxml.DeviceGetCount()
	if ret != xdxml.SUCCESS {
		log.Fatalf("Unable to get device count: %v", ret)
	}
	fmt.Printf("GPU Count: %v\n", count)

	for i := 0; i < count; i++ {

		var device xdxml.Device
		name := make([]byte, 64)

		ret = xdxml.DeviceGetHandleByIndex(i, &device)

		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get device at index %d: %v", i, ret)
		}

		ret = device.GetProductName(name)
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get uuid of device at index %d: %v", i, ret)
		}

		fmt.Printf("name: %s\n", name)

		uuid, ret := device.GetUUID()
		fmt.Printf("len(uuid): %v\n", len(uuid))

		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get uuid of device at index %d: %v", i, ret)
		}

		fmt.Printf("UUID: %s\n", uuid)

		ID, ret := device.GetMinorNumber()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get id of device at index %v: %v", ID, ret)
		}

		fmt.Printf("ID: %v\n", ID)

	}
}
