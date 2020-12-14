// +build go1.9

package ber

import "encoding/asn1"

// Compatibility vars for ber_asn1_test.go
var (
	NullRawValue = asn1.NullRawValue
	NullBytes    = asn1.NullBytes
)
