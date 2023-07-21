// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs types.go

package xdxml

type Node_path struct {
	Card               [128]int8
	Hwmon              [128]int8
	Name               [256]int8
	Dbdf               [256]int8
	Current_link_speed [256]int8
	Current_link_width [256]int8
	Max_link_speed     [256]int8
	Max_link_width     [256]int8
}

type Device struct {
	Node_path          Node_path
	Index              uint32
	Architecture       [64]int8
	Pci_dbdf           [64]int8
	Uuid               [10]int8
	Domain             uint64
	Bus                uint64
	Device             uint64
	Func               uint64
	Bdfid              uint64
	Curr_link_speed    uint32
	Curr_link_width    uint32
	Mx_link_speed      uint32
	Mx_link_width      uint32
	Curr_link_gen      uint32
	Mx_link_gen        uint32
	Dev                *_Ctype_struct_pci_device
	Device_id          uint64
	Xdx_gddr_phys_freq int32
	Xdx_gfx_freq       int32
	Xdx_acpu_freq      int32
	Xdx_cmcu_freq      int32
	Xdx_ndf_freq       int32
	Name               [64]int8
	Vbios_version      [64]int8
	Fw_ver             [64]int8
	Fb_total           uint64
}