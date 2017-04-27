// Package tap represents the VPP binary API of the 'tap' VPP module.
// DO NOT EDIT. Generated from 'bin_api/tap.api.json' on Fri, 21 Apr 2017 17:10:06 CEST.
package tap

import "github.com/fdio-stack/govpp/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x1aedb9f2

// TapConnect represents the VPP binary API message 'tap_connect'.
// Generated from 'bin_api/tap.api.json', line 6:
//
//        ["tap_connect",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "use_random_mac"],
//            ["u8", "tap_name", 64],
//            ["u8", "mac_address", 6],
//            ["u8", "renumber"],
//            ["u32", "custom_dev_instance"],
//            ["u8", "ip4_address_set"],
//            ["u8", "ip4_address", 4],
//            ["u8", "ip4_mask_width"],
//            ["u8", "ip6_address_set"],
//            ["u8", "ip6_address", 16],
//            ["u8", "ip6_mask_width"],
//            ["u8", "tag", 64],
//            {"crc" : "0x91720de3"}
//        ],
//
type TapConnect struct {
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
	IP4AddressSet     uint8
	IP4Address        []byte `struc:"[4]byte"`
	IP4MaskWidth      uint8
	IP6AddressSet     uint8
	IP6Address        []byte `struc:"[16]byte"`
	IP6MaskWidth      uint8
	Tag               []byte `struc:"[64]byte"`
}

func (*TapConnect) GetMessageName() string {
	return "tap_connect"
}
func (*TapConnect) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TapConnect) GetCrcString() string {
	return "91720de3"
}
func NewTapConnect() api.Message {
	return &TapConnect{}
}

// TapConnectReply represents the VPP binary API message 'tap_connect_reply'.
// Generated from 'bin_api/tap.api.json', line 24:
//
//        ["tap_connect_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xf47feac1"}
//        ],
//
type TapConnectReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapConnectReply) GetMessageName() string {
	return "tap_connect_reply"
}
func (*TapConnectReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TapConnectReply) GetCrcString() string {
	return "f47feac1"
}
func NewTapConnectReply() api.Message {
	return &TapConnectReply{}
}

// TapModify represents the VPP binary API message 'tap_modify'.
// Generated from 'bin_api/tap.api.json', line 31:
//
//        ["tap_modify",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "use_random_mac"],
//            ["u8", "tap_name", 64],
//            ["u8", "mac_address", 6],
//            ["u8", "renumber"],
//            ["u32", "custom_dev_instance"],
//            {"crc" : "0x8abcd5f3"}
//        ],
//
type TapModify struct {
	SwIfIndex         uint32
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
}

func (*TapModify) GetMessageName() string {
	return "tap_modify"
}
func (*TapModify) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TapModify) GetCrcString() string {
	return "8abcd5f3"
}
func NewTapModify() api.Message {
	return &TapModify{}
}

// TapModifyReply represents the VPP binary API message 'tap_modify_reply'.
// Generated from 'bin_api/tap.api.json', line 43:
//
//        ["tap_modify_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x00aaf940"}
//        ],
//
type TapModifyReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapModifyReply) GetMessageName() string {
	return "tap_modify_reply"
}
func (*TapModifyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TapModifyReply) GetCrcString() string {
	return "00aaf940"
}
func NewTapModifyReply() api.Message {
	return &TapModifyReply{}
}

// TapDelete represents the VPP binary API message 'tap_delete'.
// Generated from 'bin_api/tap.api.json', line 50:
//
//        ["tap_delete",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xe99d41c1"}
//        ],
//
type TapDelete struct {
	SwIfIndex uint32
}

func (*TapDelete) GetMessageName() string {
	return "tap_delete"
}
func (*TapDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TapDelete) GetCrcString() string {
	return "e99d41c1"
}
func NewTapDelete() api.Message {
	return &TapDelete{}
}

// TapDeleteReply represents the VPP binary API message 'tap_delete_reply'.
// Generated from 'bin_api/tap.api.json', line 57:
//
//        ["tap_delete_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x0e47d140"}
//        ],
//
type TapDeleteReply struct {
	Retval int32
}

func (*TapDeleteReply) GetMessageName() string {
	return "tap_delete_reply"
}
func (*TapDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TapDeleteReply) GetCrcString() string {
	return "0e47d140"
}
func NewTapDeleteReply() api.Message {
	return &TapDeleteReply{}
}

// SwInterfaceTapDump represents the VPP binary API message 'sw_interface_tap_dump'.
// Generated from 'bin_api/tap.api.json', line 63:
//
//        ["sw_interface_tap_dump",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xbc6ddbe2"}
//        ],
//
type SwInterfaceTapDump struct {
}

func (*SwInterfaceTapDump) GetMessageName() string {
	return "sw_interface_tap_dump"
}
func (*SwInterfaceTapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceTapDump) GetCrcString() string {
	return "bc6ddbe2"
}
func NewSwInterfaceTapDump() api.Message {
	return &SwInterfaceTapDump{}
}

// SwInterfaceTapDetails represents the VPP binary API message 'sw_interface_tap_details'.
// Generated from 'bin_api/tap.api.json', line 69:
//
//        ["sw_interface_tap_details",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "dev_name", 64],
//            {"crc" : "0x0df07bc3"}
//        ]
//
type SwInterfaceTapDetails struct {
	SwIfIndex uint32
	DevName   []byte `struc:"[64]byte"`
}

func (*SwInterfaceTapDetails) GetMessageName() string {
	return "sw_interface_tap_details"
}
func (*SwInterfaceTapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceTapDetails) GetCrcString() string {
	return "0df07bc3"
}
func NewSwInterfaceTapDetails() api.Message {
	return &SwInterfaceTapDetails{}
}
