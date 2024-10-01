/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.2 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// InterfaceTypeLabel the model 'InterfaceTypeLabel'
type InterfaceTypeLabel string

// List of Interface_type_label
const (
	INTERFACETYPELABEL_VIRTUAL InterfaceTypeLabel = "Virtual"
	INTERFACETYPELABEL_BRIDGE InterfaceTypeLabel = "Bridge"
	INTERFACETYPELABEL_LINK_AGGREGATION_GROUP__LAG InterfaceTypeLabel = "Link Aggregation Group (LAG)"
	INTERFACETYPELABEL__100_BASE_FX__10_100_ME_FIBER InterfaceTypeLabel = "100BASE-FX (10/100ME FIBER)"
	INTERFACETYPELABEL__100_BASE_LFX__10_100_ME_FIBER InterfaceTypeLabel = "100BASE-LFX (10/100ME FIBER)"
	INTERFACETYPELABEL__100_BASE_TX__10_100_ME InterfaceTypeLabel = "100BASE-TX (10/100ME)"
	INTERFACETYPELABEL__100_BASE_T1__10_100_ME_SINGLE_PAIR InterfaceTypeLabel = "100BASE-T1 (10/100ME Single Pair)"
	INTERFACETYPELABEL__1000_BASE_T__1_GE InterfaceTypeLabel = "1000BASE-T (1GE)"
	INTERFACETYPELABEL__1000_BASE_TX__1_GE InterfaceTypeLabel = "1000BASE-TX (1GE)"
	INTERFACETYPELABEL__2_5_GBASE_T__2_5_GE InterfaceTypeLabel = "2.5GBASE-T (2.5GE)"
	INTERFACETYPELABEL__5_GBASE_T__5_GE InterfaceTypeLabel = "5GBASE-T (5GE)"
	INTERFACETYPELABEL__10_GBASE_T__10_GE InterfaceTypeLabel = "10GBASE-T (10GE)"
	INTERFACETYPELABEL__10_GBASE_CX4__10_GE InterfaceTypeLabel = "10GBASE-CX4 (10GE)"
	INTERFACETYPELABEL_SFP__100_ME InterfaceTypeLabel = "SFP (100ME)"
	INTERFACETYPELABEL_GBIC__1_GE InterfaceTypeLabel = "GBIC (1GE)"
	INTERFACETYPELABEL_SFP__1_GE InterfaceTypeLabel = "SFP (1GE)"
	INTERFACETYPELABEL_SFP__10_GE InterfaceTypeLabel = "SFP+ (10GE)"
	INTERFACETYPELABEL_XFP__10_GE InterfaceTypeLabel = "XFP (10GE)"
	INTERFACETYPELABEL_XENPAK__10_GE InterfaceTypeLabel = "XENPAK (10GE)"
	INTERFACETYPELABEL_X2__10_GE InterfaceTypeLabel = "X2 (10GE)"
	INTERFACETYPELABEL_SFP28__25_GE InterfaceTypeLabel = "SFP28 (25GE)"
	INTERFACETYPELABEL_SFP56__50_GE InterfaceTypeLabel = "SFP56 (50GE)"
	INTERFACETYPELABEL_QSFP__40_GE InterfaceTypeLabel = "QSFP+ (40GE)"
	INTERFACETYPELABEL_QSFP28__50_GE InterfaceTypeLabel = "QSFP28 (50GE)"
	INTERFACETYPELABEL_CFP__100_GE InterfaceTypeLabel = "CFP (100GE)"
	INTERFACETYPELABEL_CFP2__100_GE InterfaceTypeLabel = "CFP2 (100GE)"
	INTERFACETYPELABEL_CFP2__200_GE InterfaceTypeLabel = "CFP2 (200GE)"
	INTERFACETYPELABEL_CFP2__400_GE InterfaceTypeLabel = "CFP2 (400GE)"
	INTERFACETYPELABEL_CFP4__100_GE InterfaceTypeLabel = "CFP4 (100GE)"
	INTERFACETYPELABEL_CXP__100_GE InterfaceTypeLabel = "CXP (100GE)"
	INTERFACETYPELABEL_CISCO_CPAK__100_GE InterfaceTypeLabel = "Cisco CPAK (100GE)"
	INTERFACETYPELABEL_DSFP__100_GE InterfaceTypeLabel = "DSFP (100GE)"
	INTERFACETYPELABEL_SFP_DD__100_GE InterfaceTypeLabel = "SFP-DD (100GE)"
	INTERFACETYPELABEL_QSFP28__100_GE InterfaceTypeLabel = "QSFP28 (100GE)"
	INTERFACETYPELABEL_QSFP_DD__100_GE InterfaceTypeLabel = "QSFP-DD (100GE)"
	INTERFACETYPELABEL_QSFP56__200_GE InterfaceTypeLabel = "QSFP56 (200GE)"
	INTERFACETYPELABEL_QSFP_DD__200_GE InterfaceTypeLabel = "QSFP-DD (200GE)"
	INTERFACETYPELABEL_QSFP112__400_GE InterfaceTypeLabel = "QSFP112 (400GE)"
	INTERFACETYPELABEL_QSFP_DD__400_GE InterfaceTypeLabel = "QSFP-DD (400GE)"
	INTERFACETYPELABEL_OSFP__400_GE InterfaceTypeLabel = "OSFP (400GE)"
	INTERFACETYPELABEL_OSFP_RHS__400_GE InterfaceTypeLabel = "OSFP-RHS (400GE)"
	INTERFACETYPELABEL_CDFP__400_GE InterfaceTypeLabel = "CDFP (400GE)"
	INTERFACETYPELABEL_CPF8__400_GE InterfaceTypeLabel = "CPF8 (400GE)"
	INTERFACETYPELABEL_QSFP_DD__800_GE InterfaceTypeLabel = "QSFP-DD (800GE)"
	INTERFACETYPELABEL_OSFP__800_GE InterfaceTypeLabel = "OSFP (800GE)"
	INTERFACETYPELABEL__1000_BASE_KX__1_GE InterfaceTypeLabel = "1000BASE-KX (1GE)"
	INTERFACETYPELABEL__2_5_GBASE_KX__2_5_GE InterfaceTypeLabel = "2.5GBASE-KX (2.5GE)"
	INTERFACETYPELABEL__5_GBASE_KR__5_GE InterfaceTypeLabel = "5GBASE-KR (5GE)"
	INTERFACETYPELABEL__10_GBASE_KR__10_GE InterfaceTypeLabel = "10GBASE-KR (10GE)"
	INTERFACETYPELABEL__10_GBASE_KX4__10_GE InterfaceTypeLabel = "10GBASE-KX4 (10GE)"
	INTERFACETYPELABEL__25_GBASE_KR__25_GE InterfaceTypeLabel = "25GBASE-KR (25GE)"
	INTERFACETYPELABEL__40_GBASE_KR4__40_GE InterfaceTypeLabel = "40GBASE-KR4 (40GE)"
	INTERFACETYPELABEL__50_GBASE_KR__50_GE InterfaceTypeLabel = "50GBASE-KR (50GE)"
	INTERFACETYPELABEL__100_GBASE_KP4__100_GE InterfaceTypeLabel = "100GBASE-KP4 (100GE)"
	INTERFACETYPELABEL__100_GBASE_KR2__100_GE InterfaceTypeLabel = "100GBASE-KR2 (100GE)"
	INTERFACETYPELABEL__100_GBASE_KR4__100_GE InterfaceTypeLabel = "100GBASE-KR4 (100GE)"
	INTERFACETYPELABEL_IEEE_802_11A InterfaceTypeLabel = "IEEE 802.11a"
	INTERFACETYPELABEL_IEEE_802_11B_G InterfaceTypeLabel = "IEEE 802.11b/g"
	INTERFACETYPELABEL_IEEE_802_11N InterfaceTypeLabel = "IEEE 802.11n"
	INTERFACETYPELABEL_IEEE_802_11AC InterfaceTypeLabel = "IEEE 802.11ac"
	INTERFACETYPELABEL_IEEE_802_11AD InterfaceTypeLabel = "IEEE 802.11ad"
	INTERFACETYPELABEL_IEEE_802_11AX InterfaceTypeLabel = "IEEE 802.11ax"
	INTERFACETYPELABEL_IEEE_802_11AY InterfaceTypeLabel = "IEEE 802.11ay"
	INTERFACETYPELABEL_IEEE_802_11BE InterfaceTypeLabel = "IEEE 802.11be"
	INTERFACETYPELABEL_IEEE_802_15_1__BLUETOOTH InterfaceTypeLabel = "IEEE 802.15.1 (Bluetooth)"
	INTERFACETYPELABEL_IEEE_802_15_4__LR_WPAN InterfaceTypeLabel = "IEEE 802.15.4 (LR-WPAN)"
	INTERFACETYPELABEL_OTHER__WIRELESS InterfaceTypeLabel = "Other (Wireless)"
	INTERFACETYPELABEL_GSM InterfaceTypeLabel = "GSM"
	INTERFACETYPELABEL_CDMA InterfaceTypeLabel = "CDMA"
	INTERFACETYPELABEL_LTE InterfaceTypeLabel = "LTE"
	INTERFACETYPELABEL__4_G InterfaceTypeLabel = "4G"
	INTERFACETYPELABEL__5_G InterfaceTypeLabel = "5G"
	INTERFACETYPELABEL_OC_3_STM_1 InterfaceTypeLabel = "OC-3/STM-1"
	INTERFACETYPELABEL_OC_12_STM_4 InterfaceTypeLabel = "OC-12/STM-4"
	INTERFACETYPELABEL_OC_48_STM_16 InterfaceTypeLabel = "OC-48/STM-16"
	INTERFACETYPELABEL_OC_192_STM_64 InterfaceTypeLabel = "OC-192/STM-64"
	INTERFACETYPELABEL_OC_768_STM_256 InterfaceTypeLabel = "OC-768/STM-256"
	INTERFACETYPELABEL_OC_1920_STM_640 InterfaceTypeLabel = "OC-1920/STM-640"
	INTERFACETYPELABEL_OC_3840_STM_1234 InterfaceTypeLabel = "OC-3840/STM-1234"
	INTERFACETYPELABEL_SFP__1_GFC InterfaceTypeLabel = "SFP (1GFC)"
	INTERFACETYPELABEL_SFP__2_GFC InterfaceTypeLabel = "SFP (2GFC)"
	INTERFACETYPELABEL_SFP__4_GFC InterfaceTypeLabel = "SFP (4GFC)"
	INTERFACETYPELABEL_SFP__8_GFC InterfaceTypeLabel = "SFP+ (8GFC)"
	INTERFACETYPELABEL_SFP__16_GFC InterfaceTypeLabel = "SFP+ (16GFC)"
	INTERFACETYPELABEL_SFP28__32_GFC InterfaceTypeLabel = "SFP28 (32GFC)"
	INTERFACETYPELABEL_SFP__32_GFC InterfaceTypeLabel = "SFP+ (32GFC)"
	INTERFACETYPELABEL_QSFP__64_GFC InterfaceTypeLabel = "QSFP+ (64GFC)"
	INTERFACETYPELABEL_SFP_DD__64_GFC InterfaceTypeLabel = "SFP-DD (64GFC)"
	INTERFACETYPELABEL_SFP__64_GFC InterfaceTypeLabel = "SFP+ (64GFC)"
	INTERFACETYPELABEL_QSFP28__128_GFC InterfaceTypeLabel = "QSFP28 (128GFC)"
	INTERFACETYPELABEL_SDR__2_GBPS InterfaceTypeLabel = "SDR (2 Gbps)"
	INTERFACETYPELABEL_DDR__4_GBPS InterfaceTypeLabel = "DDR (4 Gbps)"
	INTERFACETYPELABEL_QDR__8_GBPS InterfaceTypeLabel = "QDR (8 Gbps)"
	INTERFACETYPELABEL_FDR10__10_GBPS InterfaceTypeLabel = "FDR10 (10 Gbps)"
	INTERFACETYPELABEL_FDR__13_5_GBPS InterfaceTypeLabel = "FDR (13.5 Gbps)"
	INTERFACETYPELABEL_EDR__25_GBPS InterfaceTypeLabel = "EDR (25 Gbps)"
	INTERFACETYPELABEL_HDR__50_GBPS InterfaceTypeLabel = "HDR (50 Gbps)"
	INTERFACETYPELABEL_NDR__100_GBPS InterfaceTypeLabel = "NDR (100 Gbps)"
	INTERFACETYPELABEL_XDR__250_GBPS InterfaceTypeLabel = "XDR (250 Gbps)"
	INTERFACETYPELABEL_T1__1_544_MBPS InterfaceTypeLabel = "T1 (1.544 Mbps)"
	INTERFACETYPELABEL_E1__2_048_MBPS InterfaceTypeLabel = "E1 (2.048 Mbps)"
	INTERFACETYPELABEL_T3__45_MBPS InterfaceTypeLabel = "T3 (45 Mbps)"
	INTERFACETYPELABEL_E3__34_MBPS InterfaceTypeLabel = "E3 (34 Mbps)"
	INTERFACETYPELABEL_X_DSL InterfaceTypeLabel = "xDSL"
	INTERFACETYPELABEL_DOCSIS InterfaceTypeLabel = "DOCSIS"
	INTERFACETYPELABEL_BPON__622_MBPS___155_MBPS InterfaceTypeLabel = "BPON (622 Mbps / 155 Mbps)"
	INTERFACETYPELABEL_EPON__1_GBPS InterfaceTypeLabel = "EPON (1 Gbps)"
	INTERFACETYPELABEL__10_G_EPON__10_GBPS InterfaceTypeLabel = "10G-EPON (10 Gbps)"
	INTERFACETYPELABEL_GPON__2_5_GBPS___1_25_GBPS InterfaceTypeLabel = "GPON (2.5 Gbps / 1.25 Gbps)"
	INTERFACETYPELABEL_XG_PON__10_GBPS___2_5_GBPS InterfaceTypeLabel = "XG-PON (10 Gbps / 2.5 Gbps)"
	INTERFACETYPELABEL_XGS_PON__10_GBPS InterfaceTypeLabel = "XGS-PON (10 Gbps)"
	INTERFACETYPELABEL_NG_PON2__TWDM_PON__4X10_GBPS InterfaceTypeLabel = "NG-PON2 (TWDM-PON) (4x10 Gbps)"
	INTERFACETYPELABEL__25_G_PON__25_GBPS InterfaceTypeLabel = "25G-PON (25 Gbps)"
	INTERFACETYPELABEL__50_G_PON__50_GBPS InterfaceTypeLabel = "50G-PON (50 Gbps)"
	INTERFACETYPELABEL_CISCO_STACK_WISE InterfaceTypeLabel = "Cisco StackWise"
	INTERFACETYPELABEL_CISCO_STACK_WISE_PLUS InterfaceTypeLabel = "Cisco StackWise Plus"
	INTERFACETYPELABEL_CISCO_FLEX_STACK InterfaceTypeLabel = "Cisco FlexStack"
	INTERFACETYPELABEL_CISCO_FLEX_STACK_PLUS InterfaceTypeLabel = "Cisco FlexStack Plus"
	INTERFACETYPELABEL_CISCO_STACK_WISE_80 InterfaceTypeLabel = "Cisco StackWise-80"
	INTERFACETYPELABEL_CISCO_STACK_WISE_160 InterfaceTypeLabel = "Cisco StackWise-160"
	INTERFACETYPELABEL_CISCO_STACK_WISE_320 InterfaceTypeLabel = "Cisco StackWise-320"
	INTERFACETYPELABEL_CISCO_STACK_WISE_480 InterfaceTypeLabel = "Cisco StackWise-480"
	INTERFACETYPELABEL_CISCO_STACK_WISE_1_T InterfaceTypeLabel = "Cisco StackWise-1T"
	INTERFACETYPELABEL_JUNIPER_VCP InterfaceTypeLabel = "Juniper VCP"
	INTERFACETYPELABEL_EXTREME_SUMMIT_STACK InterfaceTypeLabel = "Extreme SummitStack"
	INTERFACETYPELABEL_EXTREME_SUMMIT_STACK_128 InterfaceTypeLabel = "Extreme SummitStack-128"
	INTERFACETYPELABEL_EXTREME_SUMMIT_STACK_256 InterfaceTypeLabel = "Extreme SummitStack-256"
	INTERFACETYPELABEL_EXTREME_SUMMIT_STACK_512 InterfaceTypeLabel = "Extreme SummitStack-512"
	INTERFACETYPELABEL_OTHER InterfaceTypeLabel = "Other"
)

// All allowed values of InterfaceTypeLabel enum
var AllowedInterfaceTypeLabelEnumValues = []InterfaceTypeLabel{
	"Virtual",
	"Bridge",
	"Link Aggregation Group (LAG)",
	"100BASE-FX (10/100ME FIBER)",
	"100BASE-LFX (10/100ME FIBER)",
	"100BASE-TX (10/100ME)",
	"100BASE-T1 (10/100ME Single Pair)",
	"1000BASE-T (1GE)",
	"1000BASE-TX (1GE)",
	"2.5GBASE-T (2.5GE)",
	"5GBASE-T (5GE)",
	"10GBASE-T (10GE)",
	"10GBASE-CX4 (10GE)",
	"SFP (100ME)",
	"GBIC (1GE)",
	"SFP (1GE)",
	"SFP+ (10GE)",
	"XFP (10GE)",
	"XENPAK (10GE)",
	"X2 (10GE)",
	"SFP28 (25GE)",
	"SFP56 (50GE)",
	"QSFP+ (40GE)",
	"QSFP28 (50GE)",
	"CFP (100GE)",
	"CFP2 (100GE)",
	"CFP2 (200GE)",
	"CFP2 (400GE)",
	"CFP4 (100GE)",
	"CXP (100GE)",
	"Cisco CPAK (100GE)",
	"DSFP (100GE)",
	"SFP-DD (100GE)",
	"QSFP28 (100GE)",
	"QSFP-DD (100GE)",
	"QSFP56 (200GE)",
	"QSFP-DD (200GE)",
	"QSFP112 (400GE)",
	"QSFP-DD (400GE)",
	"OSFP (400GE)",
	"OSFP-RHS (400GE)",
	"CDFP (400GE)",
	"CPF8 (400GE)",
	"QSFP-DD (800GE)",
	"OSFP (800GE)",
	"1000BASE-KX (1GE)",
	"2.5GBASE-KX (2.5GE)",
	"5GBASE-KR (5GE)",
	"10GBASE-KR (10GE)",
	"10GBASE-KX4 (10GE)",
	"25GBASE-KR (25GE)",
	"40GBASE-KR4 (40GE)",
	"50GBASE-KR (50GE)",
	"100GBASE-KP4 (100GE)",
	"100GBASE-KR2 (100GE)",
	"100GBASE-KR4 (100GE)",
	"IEEE 802.11a",
	"IEEE 802.11b/g",
	"IEEE 802.11n",
	"IEEE 802.11ac",
	"IEEE 802.11ad",
	"IEEE 802.11ax",
	"IEEE 802.11ay",
	"IEEE 802.11be",
	"IEEE 802.15.1 (Bluetooth)",
	"IEEE 802.15.4 (LR-WPAN)",
	"Other (Wireless)",
	"GSM",
	"CDMA",
	"LTE",
	"4G",
	"5G",
	"OC-3/STM-1",
	"OC-12/STM-4",
	"OC-48/STM-16",
	"OC-192/STM-64",
	"OC-768/STM-256",
	"OC-1920/STM-640",
	"OC-3840/STM-1234",
	"SFP (1GFC)",
	"SFP (2GFC)",
	"SFP (4GFC)",
	"SFP+ (8GFC)",
	"SFP+ (16GFC)",
	"SFP28 (32GFC)",
	"SFP+ (32GFC)",
	"QSFP+ (64GFC)",
	"SFP-DD (64GFC)",
	"SFP+ (64GFC)",
	"QSFP28 (128GFC)",
	"SDR (2 Gbps)",
	"DDR (4 Gbps)",
	"QDR (8 Gbps)",
	"FDR10 (10 Gbps)",
	"FDR (13.5 Gbps)",
	"EDR (25 Gbps)",
	"HDR (50 Gbps)",
	"NDR (100 Gbps)",
	"XDR (250 Gbps)",
	"T1 (1.544 Mbps)",
	"E1 (2.048 Mbps)",
	"T3 (45 Mbps)",
	"E3 (34 Mbps)",
	"xDSL",
	"DOCSIS",
	"BPON (622 Mbps / 155 Mbps)",
	"EPON (1 Gbps)",
	"10G-EPON (10 Gbps)",
	"GPON (2.5 Gbps / 1.25 Gbps)",
	"XG-PON (10 Gbps / 2.5 Gbps)",
	"XGS-PON (10 Gbps)",
	"NG-PON2 (TWDM-PON) (4x10 Gbps)",
	"25G-PON (25 Gbps)",
	"50G-PON (50 Gbps)",
	"Cisco StackWise",
	"Cisco StackWise Plus",
	"Cisco FlexStack",
	"Cisco FlexStack Plus",
	"Cisco StackWise-80",
	"Cisco StackWise-160",
	"Cisco StackWise-320",
	"Cisco StackWise-480",
	"Cisco StackWise-1T",
	"Juniper VCP",
	"Extreme SummitStack",
	"Extreme SummitStack-128",
	"Extreme SummitStack-256",
	"Extreme SummitStack-512",
	"Other",
}

func (v *InterfaceTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceTypeLabel(value)
	for _, existing := range AllowedInterfaceTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceTypeLabel", value)
}

// NewInterfaceTypeLabelFromValue returns a pointer to a valid InterfaceTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceTypeLabelFromValue(v string) (*InterfaceTypeLabel, error) {
	ev := InterfaceTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceTypeLabel: valid values are %v", v, AllowedInterfaceTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceTypeLabel) IsValid() bool {
	for _, existing := range AllowedInterfaceTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_type_label value
func (v InterfaceTypeLabel) Ptr() *InterfaceTypeLabel {
	return &v
}

type NullableInterfaceTypeLabel struct {
	value *InterfaceTypeLabel
	isSet bool
}

func (v NullableInterfaceTypeLabel) Get() *InterfaceTypeLabel {
	return v.value
}

func (v *NullableInterfaceTypeLabel) Set(val *InterfaceTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTypeLabel(val *InterfaceTypeLabel) *NullableInterfaceTypeLabel {
	return &NullableInterfaceTypeLabel{value: val, isSet: true}
}

func (v NullableInterfaceTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

