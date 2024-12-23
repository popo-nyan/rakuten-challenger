package rc

import (
	"encoding/binary"
	"fmt"
)

func getUTF8Bytes(input string) []byte {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		charCode := input[i]
		if charCode > 127 {
			return []byte(input)
		}
		result[i] = charCode
	}
	return result
}

func x64Add(m, n []uint32) {
	m0, m1 := m[0]>>16, m[0]&0xffff
	m2, m3 := m[1]>>16, m[1]&0xffff

	n0, n1 := n[0]>>16, n[0]&0xffff
	n2, n3 := n[1]>>16, n[1]&0xffff

	o0, o1, o2, o3 := uint32(0), uint32(0), uint32(0), uint32(0)
	o3 += m3 + n3
	o2 += o3 >> 16
	o3 &= 0xffff
	o2 += m2 + n2
	o1 += o2 >> 16
	o2 &= 0xffff
	o1 += m1 + n1
	o0 += o1 >> 16
	o1 &= 0xffff
	o0 += m0 + n0
	o0 &= 0xffff

	m[0] = (o0 << 16) | o1
	m[1] = (o2 << 16) | o3
}

func x64Multiply(m, n []uint32) {
	m0, m1 := m[0]>>16, m[0]&0xffff
	m2, m3 := m[1]>>16, m[1]&0xffff

	n0, n1 := n[0]>>16, n[0]&0xffff
	n2, n3 := n[1]>>16, n[1]&0xffff

	o0, o1, o2, o3 := uint32(0), uint32(0), uint32(0), uint32(0)

	o3 += m3 * n3
	o2 += o3 >> 16
	o3 &= 0xffff
	o2 += m2 * n3
	o1 += o2 >> 16
	o2 &= 0xffff
	o2 += m3 * n2
	o1 += o2 >> 16
	o2 &= 0xffff
	o1 += m1 * n3
	o0 += o1 >> 16
	o1 &= 0xffff
	o1 += m2 * n2
	o0 += o1 >> 16
	o1 &= 0xffff
	o1 += m3 * n1
	o0 += o1 >> 16
	o1 &= 0xffff
	o0 += m0*n3 + m1*n2 + m2*n1 + m3*n0
	o0 &= 0xffff

	m[0] = (o0 << 16) | o1
	m[1] = (o2 << 16) | o3
}

func x64Rotl(m []uint32, bits uint32) {
	m0 := m[0]
	bits %= 64
	if bits == 32 {
		m[0], m[1] = m[1], m0
	} else if bits < 32 {
		m[0] = (m0 << bits) | (m[1] >> (32 - bits))
		m[1] = (m[1] << bits) | (m0 >> (32 - bits))
	} else {
		bits -= 32
		m[0] = (m[1] << bits) | (m0 >> (32 - bits))
		m[1] = (m0 << bits) | (m[1] >> (32 - bits))
	}
}

func x64LeftShift(m []uint32, bits uint32) {
	bits %= 64
	if bits == 0 {
		return
	} else if bits < 32 {
		m[0] = m[1] >> (32 - bits)
		m[1] = m[1] << bits
	} else {
		m[0] = m[1] << (bits - 32)
		m[1] = 0
	}
}

func x64Xor(m, n []uint32) {
	m[0] ^= n[0]
	m[1] ^= n[1]
}

var F1 = []uint32{0xff51afd7, 0xed558ccd}
var F2 = []uint32{0xc4ceb9fe, 0x1a85ec53}

func x64Fmix(h []uint32) {
	shifted := []uint32{0, h[0] >> 1}
	x64Xor(h, shifted)
	x64Multiply(h, F1)
	shifted[1] = h[0] >> 1
	x64Xor(h, shifted)
	x64Multiply(h, F2)
	shifted[1] = h[0] >> 1
	x64Xor(h, shifted)
}

var C1 = []uint32{0x87c37b91, 0x114253d5}
var C2 = []uint32{0x4cf5ad43, 0x2745937f}
var M = []uint32{0, 5}
var N1 = []uint32{0, 0x52dce729}
var N2 = []uint32{0, 0x38495ab5}

func x64hash128(input string, seed uint32) string {
	key := getUTF8Bytes(input)
	length := []uint32{0, uint32(len(key))}
	remainder := length[1] % 16
	bytes := length[1] - remainder
	h1 := []uint32{0, seed}
	h2 := []uint32{0, seed}
	k1 := []uint32{0, 0}
	k2 := []uint32{0, 0}

	for i := 0; i < int(bytes); i += 16 {
		k1[0] = binary.LittleEndian.Uint32(key[i+4 : i+8])
		k1[1] = binary.LittleEndian.Uint32(key[i : i+4])
		k2[0] = binary.LittleEndian.Uint32(key[i+12 : i+16])
		k2[1] = binary.LittleEndian.Uint32(key[i+8 : i+12])

		x64Multiply(k1, C1)
		x64Rotl(k1, 31)
		x64Multiply(k1, C2)
		x64Xor(h1, k1)
		x64Rotl(h1, 27)
		x64Add(h1, h2)
		x64Multiply(h1, M)
		x64Add(h1, N1)
		x64Multiply(k2, C2)
		x64Rotl(k2, 33)
		x64Multiply(k2, C1)
		x64Xor(h2, k2)
		x64Rotl(h2, 31)
		x64Add(h2, h1)
		x64Multiply(h2, M)
		x64Add(h2, N2)
	}

	k1 = []uint32{0, 0}
	k2 = []uint32{0, 0}
	val := []uint32{0, 0}

	switch remainder {
	case 15:
		val[1] = uint32(key[bytes+14])
		x64LeftShift(val, 48)
		x64Xor(k2, val)
	case 14:
		val[1] = uint32(key[bytes+13])
		x64LeftShift(val, 40)
		x64Xor(k2, val)
	case 13:
		val[1] = uint32(key[bytes+12])
		x64LeftShift(val, 32)
		x64Xor(k2, val)
	case 12:
		val[1] = uint32(key[bytes+11])
		x64LeftShift(val, 24)
		x64Xor(k2, val)
	case 11:
		val[1] = uint32(key[bytes+10])
		x64LeftShift(val, 16)
		x64Xor(k2, val)
	case 10:
		val[1] = uint32(key[bytes+9])
		x64LeftShift(val, 8)
		x64Xor(k2, val)
	case 9:
		val[1] = uint32(key[bytes+8])
		x64Xor(k2, val)
		x64Multiply(k2, C2)
		x64Rotl(k2, 33)
		x64Multiply(k2, C1)
		x64Xor(h2, k2)
	case 8:
		val[1] = uint32(key[bytes+7])
		x64LeftShift(val, 56)
		x64Xor(k1, val)
	case 7:
		val[1] = uint32(key[bytes+6])
		x64LeftShift(val, 48)
		x64Xor(k1, val)
	case 6:
		val[1] = uint32(key[bytes+5])
		x64LeftShift(val, 40)
		x64Xor(k1, val)
	case 5:
		val[1] = uint32(key[bytes+4])
		x64LeftShift(val, 32)
		x64Xor(k1, val)
	case 4:
		val[1] = uint32(key[bytes+3])
		x64LeftShift(val, 24)
		x64Xor(k1, val)
	case 3:
		val[1] = uint32(key[bytes+2])
		x64LeftShift(val, 16)
		x64Xor(k1, val)
	case 2:
		val[1] = uint32(key[bytes+1])
		x64LeftShift(val, 8)
		x64Xor(k1, val)
	case 1:
		val[1] = uint32(key[bytes])
		x64Xor(k1, val)
		x64Multiply(k1, C1)
		x64Rotl(k1, 31)
		x64Multiply(k1, C2)
		x64Xor(h1, k1)
	}

	x64Xor(h1, length)
	x64Xor(h2, length)
	x64Add(h1, h2)
	x64Add(h2, h1)
	x64Fmix(h1)
	x64Fmix(h2)
	x64Add(h1, h2)
	x64Add(h2, h1)

	return fmt.Sprintf("%08x%08x%08x%08x", h1[0], h1[1], h2[0], h2[1])
}
