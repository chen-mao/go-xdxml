package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

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
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get uuid of device at uuid %d: %v", i, ret)
		}
		fmt.Printf("len(uuid): %v, uuid: %v\n", len(uuid), uuid)

		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get uuid of device at index %d: %v", i, ret)
		}

		fmt.Printf("UUID: %s\n", uuid)

		ID, ret := device.GetMinorNumber()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get id of device at index %v: %v", ID, ret)
		}

		fmt.Printf("ID: %v\n", ID)

		pci_info, ret := device.GetPciInfo()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get id of device at index %v: %v", ID, ret)
		}
		fmt.Printf("PciBus: %d\n", pci_info.Bus)
		fmt.Printf("PciBusFunc: %d\n", pci_info.Func)
		pciBusID := getBusID(pci_info)
		fmt.Printf("PciBusID: %s\n", pciBusID)
		drmDeviceNodes, err := GetDeviceNodesByBusID(pciBusID)
		if err != nil {
			fmt.Errorf("failed to determine DRM devices for %v: %v", pciBusID, err)
		}
		fmt.Printf("demDeviceNodes: %s\n", drmDeviceNodes)
		// fmt.Printf("Test: %s\n", TestgetBusID())

		memory_st, ret := device.GetMemoryInfo()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to get mem of device at index %v: %v", ID, ret)
		}
		fmt.Printf("memory: %v MB\n", memory_st.Fb_total/(1024*1024))
	}
}

func getBusID(p xdxml.Pci_info) string {
	busStr := fmt.Sprintf("%02x", p.Bus)
	deviceStr := fmt.Sprintf("%02x", p.Device)
	funcStr := fmt.Sprintf("%1x", p.Func)

	id := fmt.Sprintf("0000:%s:%s.%s", busStr, deviceStr, funcStr)
	id = strings.ToLower(id)
	return id
}

func GetDeviceNodesByBusID(busID string) ([]string, error) {
	drmRoot := filepath.Join("/sys/bus/pci/devices", busID, "drm")
	matches_card, err := filepath.Glob(fmt.Sprintf("%s/card*", drmRoot))
	if err != nil {
		return nil, err
	}
	matches_renderD, err := filepath.Glob(fmt.Sprintf("%s/renderD*", drmRoot))
	if err != nil {
		return nil, err
	}

	matches := append(matches_card, matches_renderD...)
	var drmDeviceNodes []string
	for _, m := range matches {
		drmDeviceNode := filepath.Join("/dev/dri", filepath.Base(m))

		drmDeviceNodes = append(drmDeviceNodes, drmDeviceNode)
	}

	return drmDeviceNodes, nil
}
