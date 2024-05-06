package xdxml

// xdxml.DeviceGetMinorNumber()
func DeviceGetMinorNumber(Device Device) (int, Return) {
	var minorNumber int32
	ret := xdxml_device_get_minor_number(Device, &minorNumber)
	return int(minorNumber), ret
}

func (Device Device) GetMinorNumber() (int, Return) {
	return DeviceGetMinorNumber(Device)
}

// xdxml.DeviceGetCount()
func DeviceGetCount() (int, Return) {
	var DeviceCount uint32
	ret := xdxml_device_get_count(&DeviceCount)
	return int(DeviceCount), ret
}

// xdxml.DeviceGetHandleByIndex()
func DeviceGetHandleByIndex(Index int, device *Device) Return {
	ret := xdxml_device_get_handle_by_index(uint32(Index), device)
	return ret
}

// xdxml.DeviceGetHandleByIndex()
func DeviceGetProductName(device Device, name []byte) Return {
	ret := xdxml_device_get_product_name(device, &name[0])
	return ret
}

func (Device Device) GetProductName(name []byte) Return {
	return DeviceGetProductName(Device, name)
}

// xdxml.DeviceGetUUID()
func DeviceGetUUID(Device Device) (string, Return) {
	var uuidBytes [8]byte
	ret := xdxml_device_get_uuid(Device, &uuidBytes[0])
	// var uuidStr string
	// for _, num := range Device.Handle.uuid {
	// 	uuidStr += string(num)
	// }
	uuidStr := string(uuidBytes[:])
	return uuidStr, ret
}

func (Device Device) GetUUID() (string, Return) {
	return DeviceGetUUID(Device)
}

func (Device Device) GetPciInfo() (Pci_info, Return) {
	var pci Pci_info
	ret := xdxml_device_get_pci_info(Device, &pci)
	return pci, ret
}
