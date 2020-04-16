// +build amd64

package ber

import "encoding/asn1"

func init() {
	objectIdentifierTestData = append(objectIdentifierTestData, objectIdentifierTest{
		[]byte{
			0x2b, 0x06, 0x01, 0x04, 0x01, 0x09, 0x0a, 0x81,
			0x0a, 0x01, 0x04, 0x01, 0x02, 0x01, 0x8c, 0x80,
			0x80, 0x80, 0x01},
		true,
		[]int{1, 3, 6, 1, 4, 1, 9, 10, 138, 1, 4, 1, 2, 1, 3221225473},
	})

	marshalTests = append(marshalTests, marshalTest{
		asn1.ObjectIdentifier([]int{1, 3, 6, 1, 4, 1, 9, 10, 138, 1, 4, 1, 2, 1, 3221225473}),
		"06132b06010401090a810a01040102018c80808001",
	})
}
