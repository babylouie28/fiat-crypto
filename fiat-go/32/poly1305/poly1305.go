// Code generated by Fiat Cryptography. DO NOT EDIT.
//
// Autogenerated: 'src/ExtractionOCaml/unsaturated_solinas' --lang Go --relax-primitive-carry-to-bitwidth 32,64 --cmovznz-by-mul --internal-static --package-case flatcase --public-function-case UpperCamelCase --private-function-case camelCase --public-type-case UpperCamelCase --private-type-case camelCase --no-prefix-fiat --doc-newline-in-typedef-bounds --doc-prepend-header 'Code generated by Fiat Cryptography. DO NOT EDIT.' --doc-text-before-function-name '' --doc-text-before-type-name '' --package-name poly1305 '' 32 '(auto)' '2^130 - 5' carry_mul carry_square carry add sub opp selectznz to_bytes from_bytes relax carry_add carry_sub carry_opp
//
// curve description (via package name): poly1305
//
// machine_wordsize = 32 (from "32")
//
// requested operations: carry_mul, carry_square, carry, add, sub, opp, selectznz, to_bytes, from_bytes, relax, carry_add, carry_sub, carry_opp
//
// n = 5 (from "(auto)")
//
// s-c = 2^130 - [(1, 5)] (from "2^130 - 5")
//
// tight_bounds_multiplier = 1 (from "")
//
//
//
// Computed values:
//
//   carry_chain = [0, 1, 2, 3, 4, 0, 1]
//
//   eval z = z[0] + (z[1] << 26) + (z[2] << 52) + (z[3] << 78) + (z[4] << 104)
//
//   bytes_eval z = z[0] + (z[1] << 8) + (z[2] << 16) + (z[3] << 24) + (z[4] << 32) + (z[5] << 40) + (z[6] << 48) + (z[7] << 56) + (z[8] << 64) + (z[9] << 72) + (z[10] << 80) + (z[11] << 88) + (z[12] << 96) + (z[13] << 104) + (z[14] << 112) + (z[15] << 120) + (z[16] << 128)
//
//   balance = [0x7fffff6, 0x7fffffe, 0x7fffffe, 0x7fffffe, 0x7fffffe]
package poly1305

type uint1 uint64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927
type int1 int64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927

// LooseFieldElement is a field element with loose bounds.
//
// Bounds:
//
//   [[0x0 ~> 0xc000000], [0x0 ~> 0xc000000], [0x0 ~> 0xc000000], [0x0 ~> 0xc000000], [0x0 ~> 0xc000000]]
type LooseFieldElement [5]uint32

// TightFieldElement is a field element with tight bounds.
//
// Bounds:
//
//   [[0x0 ~> 0x4000000], [0x0 ~> 0x4000000], [0x0 ~> 0x4000000], [0x0 ~> 0x4000000], [0x0 ~> 0x4000000]]
type TightFieldElement [5]uint32

// addcarryxU26 is an addition with carry.
//
// Postconditions:
//   out1 = (arg1 + arg2 + arg3) mod 2^26
//   out2 = ⌊(arg1 + arg2 + arg3) / 2^26⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x3ffffff]
//   arg3: [0x0 ~> 0x3ffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x3ffffff]
//   out2: [0x0 ~> 0x1]
func addcarryxU26(out1 *uint32, out2 *uint1, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := ((uint32(arg1) + arg2) + arg3)
	x2 := (x1 & 0x3ffffff)
	x3 := uint1((x1 >> 26))
	*out1 = x2
	*out2 = x3
}

// subborrowxU26 is a subtraction with borrow.
//
// Postconditions:
//   out1 = (-arg1 + arg2 + -arg3) mod 2^26
//   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^26⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x3ffffff]
//   arg3: [0x0 ~> 0x3ffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x3ffffff]
//   out2: [0x0 ~> 0x1]
func subborrowxU26(out1 *uint32, out2 *uint1, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := ((int32(arg2) - int32(arg1)) - int32(arg3))
	x2 := int1((x1 >> 26))
	x3 := (uint32(x1) & 0x3ffffff)
	*out1 = x3
	*out2 = (0x0 - uint1(x2))
}

// cmovznzU32 is a single-word conditional move.
//
// Postconditions:
//   out1 = (if arg1 = 0 then arg2 else arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0xffffffff]
//   arg3: [0x0 ~> 0xffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0xffffffff]
func cmovznzU32(out1 *uint32, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := (uint32(arg1) * 0xffffffff)
	x2 := ((x1 & arg3) | ((^x1) & arg2))
	*out1 = x2
}

// CarryMul multiplies two field elements and reduces the result.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 * eval arg2) mod m
//
func CarryMul(out1 *TightFieldElement, arg1 *LooseFieldElement, arg2 *LooseFieldElement) {
	x1 := (uint64(arg1[4]) * uint64((arg2[4] * 0x5)))
	x2 := (uint64(arg1[4]) * uint64((arg2[3] * 0x5)))
	x3 := (uint64(arg1[4]) * uint64((arg2[2] * 0x5)))
	x4 := (uint64(arg1[4]) * uint64((arg2[1] * 0x5)))
	x5 := (uint64(arg1[3]) * uint64((arg2[4] * 0x5)))
	x6 := (uint64(arg1[3]) * uint64((arg2[3] * 0x5)))
	x7 := (uint64(arg1[3]) * uint64((arg2[2] * 0x5)))
	x8 := (uint64(arg1[2]) * uint64((arg2[4] * 0x5)))
	x9 := (uint64(arg1[2]) * uint64((arg2[3] * 0x5)))
	x10 := (uint64(arg1[1]) * uint64((arg2[4] * 0x5)))
	x11 := (uint64(arg1[4]) * uint64(arg2[0]))
	x12 := (uint64(arg1[3]) * uint64(arg2[1]))
	x13 := (uint64(arg1[3]) * uint64(arg2[0]))
	x14 := (uint64(arg1[2]) * uint64(arg2[2]))
	x15 := (uint64(arg1[2]) * uint64(arg2[1]))
	x16 := (uint64(arg1[2]) * uint64(arg2[0]))
	x17 := (uint64(arg1[1]) * uint64(arg2[3]))
	x18 := (uint64(arg1[1]) * uint64(arg2[2]))
	x19 := (uint64(arg1[1]) * uint64(arg2[1]))
	x20 := (uint64(arg1[1]) * uint64(arg2[0]))
	x21 := (uint64(arg1[0]) * uint64(arg2[4]))
	x22 := (uint64(arg1[0]) * uint64(arg2[3]))
	x23 := (uint64(arg1[0]) * uint64(arg2[2]))
	x24 := (uint64(arg1[0]) * uint64(arg2[1]))
	x25 := (uint64(arg1[0]) * uint64(arg2[0]))
	x26 := (x25 + (x10 + (x9 + (x7 + x4))))
	x27 := (x26 >> 26)
	x28 := (uint32(x26) & 0x3ffffff)
	x29 := (x21 + (x17 + (x14 + (x12 + x11))))
	x30 := (x22 + (x18 + (x15 + (x13 + x1))))
	x31 := (x23 + (x19 + (x16 + (x5 + x2))))
	x32 := (x24 + (x20 + (x8 + (x6 + x3))))
	x33 := (x27 + x32)
	x34 := (x33 >> 26)
	x35 := (uint32(x33) & 0x3ffffff)
	x36 := (x34 + x31)
	x37 := (x36 >> 26)
	x38 := (uint32(x36) & 0x3ffffff)
	x39 := (x37 + x30)
	x40 := (x39 >> 26)
	x41 := (uint32(x39) & 0x3ffffff)
	x42 := (x40 + x29)
	x43 := uint32((x42 >> 26))
	x44 := (uint32(x42) & 0x3ffffff)
	x45 := (uint64(x43) * uint64(0x5))
	x46 := (uint64(x28) + x45)
	x47 := uint32((x46 >> 26))
	x48 := (uint32(x46) & 0x3ffffff)
	x49 := (x47 + x35)
	x50 := uint1((x49 >> 26))
	x51 := (x49 & 0x3ffffff)
	x52 := (uint32(x50) + x38)
	out1[0] = x48
	out1[1] = x51
	out1[2] = x52
	out1[3] = x41
	out1[4] = x44
}

// CarrySquare squares a field element and reduces the result.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 * eval arg1) mod m
//
func CarrySquare(out1 *TightFieldElement, arg1 *LooseFieldElement) {
	x1 := (arg1[4] * 0x5)
	x2 := (x1 * 0x2)
	x3 := (arg1[4] * 0x2)
	x4 := (arg1[3] * 0x5)
	x5 := (x4 * 0x2)
	x6 := (arg1[3] * 0x2)
	x7 := (arg1[2] * 0x2)
	x8 := (arg1[1] * 0x2)
	x9 := (uint64(arg1[4]) * uint64(x1))
	x10 := (uint64(arg1[3]) * uint64(x2))
	x11 := (uint64(arg1[3]) * uint64(x4))
	x12 := (uint64(arg1[2]) * uint64(x2))
	x13 := (uint64(arg1[2]) * uint64(x5))
	x14 := (uint64(arg1[2]) * uint64(arg1[2]))
	x15 := (uint64(arg1[1]) * uint64(x2))
	x16 := (uint64(arg1[1]) * uint64(x6))
	x17 := (uint64(arg1[1]) * uint64(x7))
	x18 := (uint64(arg1[1]) * uint64(arg1[1]))
	x19 := (uint64(arg1[0]) * uint64(x3))
	x20 := (uint64(arg1[0]) * uint64(x6))
	x21 := (uint64(arg1[0]) * uint64(x7))
	x22 := (uint64(arg1[0]) * uint64(x8))
	x23 := (uint64(arg1[0]) * uint64(arg1[0]))
	x24 := (x23 + (x15 + x13))
	x25 := (x24 >> 26)
	x26 := (uint32(x24) & 0x3ffffff)
	x27 := (x19 + (x16 + x14))
	x28 := (x20 + (x17 + x9))
	x29 := (x21 + (x18 + x10))
	x30 := (x22 + (x12 + x11))
	x31 := (x25 + x30)
	x32 := (x31 >> 26)
	x33 := (uint32(x31) & 0x3ffffff)
	x34 := (x32 + x29)
	x35 := (x34 >> 26)
	x36 := (uint32(x34) & 0x3ffffff)
	x37 := (x35 + x28)
	x38 := (x37 >> 26)
	x39 := (uint32(x37) & 0x3ffffff)
	x40 := (x38 + x27)
	x41 := uint32((x40 >> 26))
	x42 := (uint32(x40) & 0x3ffffff)
	x43 := (uint64(x41) * uint64(0x5))
	x44 := (uint64(x26) + x43)
	x45 := uint32((x44 >> 26))
	x46 := (uint32(x44) & 0x3ffffff)
	x47 := (x45 + x33)
	x48 := uint1((x47 >> 26))
	x49 := (x47 & 0x3ffffff)
	x50 := (uint32(x48) + x36)
	out1[0] = x46
	out1[1] = x49
	out1[2] = x50
	out1[3] = x39
	out1[4] = x42
}

// Carry reduces a field element.
//
// Postconditions:
//   eval out1 mod m = eval arg1 mod m
//
func Carry(out1 *TightFieldElement, arg1 *LooseFieldElement) {
	x1 := arg1[0]
	x2 := ((x1 >> 26) + arg1[1])
	x3 := ((x2 >> 26) + arg1[2])
	x4 := ((x3 >> 26) + arg1[3])
	x5 := ((x4 >> 26) + arg1[4])
	x6 := ((x1 & 0x3ffffff) + ((x5 >> 26) * 0x5))
	x7 := (uint32(uint1((x6 >> 26))) + (x2 & 0x3ffffff))
	x8 := (x6 & 0x3ffffff)
	x9 := (x7 & 0x3ffffff)
	x10 := (uint32(uint1((x7 >> 26))) + (x3 & 0x3ffffff))
	x11 := (x4 & 0x3ffffff)
	x12 := (x5 & 0x3ffffff)
	out1[0] = x8
	out1[1] = x9
	out1[2] = x10
	out1[3] = x11
	out1[4] = x12
}

// Add adds two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 + eval arg2) mod m
//
func Add(out1 *LooseFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := (arg1[0] + arg2[0])
	x2 := (arg1[1] + arg2[1])
	x3 := (arg1[2] + arg2[2])
	x4 := (arg1[3] + arg2[3])
	x5 := (arg1[4] + arg2[4])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
}

// Sub subtracts two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 - eval arg2) mod m
//
func Sub(out1 *LooseFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := ((0x7fffff6 + arg1[0]) - arg2[0])
	x2 := ((0x7fffffe + arg1[1]) - arg2[1])
	x3 := ((0x7fffffe + arg1[2]) - arg2[2])
	x4 := ((0x7fffffe + arg1[3]) - arg2[3])
	x5 := ((0x7fffffe + arg1[4]) - arg2[4])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
}

// Opp negates a field element.
//
// Postconditions:
//   eval out1 mod m = -eval arg1 mod m
//
func Opp(out1 *LooseFieldElement, arg1 *TightFieldElement) {
	x1 := (0x7fffff6 - arg1[0])
	x2 := (0x7fffffe - arg1[1])
	x3 := (0x7fffffe - arg1[2])
	x4 := (0x7fffffe - arg1[3])
	x5 := (0x7fffffe - arg1[4])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
}

// Selectznz is a multi-limb conditional select.
//
// Postconditions:
//   out1 = (if arg1 = 0 then arg2 else arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
//   arg3: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
func Selectznz(out1 *[5]uint32, arg1 uint1, arg2 *[5]uint32, arg3 *[5]uint32) {
	var x1 uint32
	cmovznzU32(&x1, arg1, arg2[0], arg3[0])
	var x2 uint32
	cmovznzU32(&x2, arg1, arg2[1], arg3[1])
	var x3 uint32
	cmovznzU32(&x3, arg1, arg2[2], arg3[2])
	var x4 uint32
	cmovznzU32(&x4, arg1, arg2[3], arg3[3])
	var x5 uint32
	cmovznzU32(&x5, arg1, arg2[4], arg3[4])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
}

// ToBytes serializes a field element to bytes in little-endian order.
//
// Postconditions:
//   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..16]
//
// Output Bounds:
//   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x3]]
func ToBytes(out1 *[17]uint8, arg1 *TightFieldElement) {
	var x1 uint32
	var x2 uint1
	subborrowxU26(&x1, &x2, 0x0, arg1[0], 0x3fffffb)
	var x3 uint32
	var x4 uint1
	subborrowxU26(&x3, &x4, x2, arg1[1], 0x3ffffff)
	var x5 uint32
	var x6 uint1
	subborrowxU26(&x5, &x6, x4, arg1[2], 0x3ffffff)
	var x7 uint32
	var x8 uint1
	subborrowxU26(&x7, &x8, x6, arg1[3], 0x3ffffff)
	var x9 uint32
	var x10 uint1
	subborrowxU26(&x9, &x10, x8, arg1[4], 0x3ffffff)
	var x11 uint32
	cmovznzU32(&x11, x10, uint32(0x0), 0xffffffff)
	var x12 uint32
	var x13 uint1
	addcarryxU26(&x12, &x13, 0x0, x1, (x11 & 0x3fffffb))
	var x14 uint32
	var x15 uint1
	addcarryxU26(&x14, &x15, x13, x3, (x11 & 0x3ffffff))
	var x16 uint32
	var x17 uint1
	addcarryxU26(&x16, &x17, x15, x5, (x11 & 0x3ffffff))
	var x18 uint32
	var x19 uint1
	addcarryxU26(&x18, &x19, x17, x7, (x11 & 0x3ffffff))
	var x20 uint32
	var x21 uint1
	addcarryxU26(&x20, &x21, x19, x9, (x11 & 0x3ffffff))
	x22 := (x18 << 6)
	x23 := (x16 << 4)
	x24 := (x14 << 2)
	x25 := (uint8(x12) & 0xff)
	x26 := (x12 >> 8)
	x27 := (uint8(x26) & 0xff)
	x28 := (x26 >> 8)
	x29 := (uint8(x28) & 0xff)
	x30 := uint8((x28 >> 8))
	x31 := (x24 + uint32(x30))
	x32 := (uint8(x31) & 0xff)
	x33 := (x31 >> 8)
	x34 := (uint8(x33) & 0xff)
	x35 := (x33 >> 8)
	x36 := (uint8(x35) & 0xff)
	x37 := uint8((x35 >> 8))
	x38 := (x23 + uint32(x37))
	x39 := (uint8(x38) & 0xff)
	x40 := (x38 >> 8)
	x41 := (uint8(x40) & 0xff)
	x42 := (x40 >> 8)
	x43 := (uint8(x42) & 0xff)
	x44 := uint8((x42 >> 8))
	x45 := (x22 + uint32(x44))
	x46 := (uint8(x45) & 0xff)
	x47 := (x45 >> 8)
	x48 := (uint8(x47) & 0xff)
	x49 := (x47 >> 8)
	x50 := (uint8(x49) & 0xff)
	x51 := uint8((x49 >> 8))
	x52 := (uint8(x20) & 0xff)
	x53 := (x20 >> 8)
	x54 := (uint8(x53) & 0xff)
	x55 := (x53 >> 8)
	x56 := (uint8(x55) & 0xff)
	x57 := uint8((x55 >> 8))
	out1[0] = x25
	out1[1] = x27
	out1[2] = x29
	out1[3] = x32
	out1[4] = x34
	out1[5] = x36
	out1[6] = x39
	out1[7] = x41
	out1[8] = x43
	out1[9] = x46
	out1[10] = x48
	out1[11] = x50
	out1[12] = x51
	out1[13] = x52
	out1[14] = x54
	out1[15] = x56
	out1[16] = x57
}

// FromBytes deserializes a field element from bytes in little-endian order.
//
// Postconditions:
//   eval out1 mod m = bytes_eval arg1 mod m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x3]]
func FromBytes(out1 *TightFieldElement, arg1 *[17]uint8) {
	x1 := (uint32(arg1[16]) << 24)
	x2 := (uint32(arg1[15]) << 16)
	x3 := (uint32(arg1[14]) << 8)
	x4 := arg1[13]
	x5 := (uint32(arg1[12]) << 18)
	x6 := (uint32(arg1[11]) << 10)
	x7 := (uint32(arg1[10]) << 2)
	x8 := (uint32(arg1[9]) << 20)
	x9 := (uint32(arg1[8]) << 12)
	x10 := (uint32(arg1[7]) << 4)
	x11 := (uint32(arg1[6]) << 22)
	x12 := (uint32(arg1[5]) << 14)
	x13 := (uint32(arg1[4]) << 6)
	x14 := (uint32(arg1[3]) << 24)
	x15 := (uint32(arg1[2]) << 16)
	x16 := (uint32(arg1[1]) << 8)
	x17 := arg1[0]
	x18 := (x16 + uint32(x17))
	x19 := (x15 + x18)
	x20 := (x14 + x19)
	x21 := (x20 & 0x3ffffff)
	x22 := uint8((x20 >> 26))
	x23 := (x13 + uint32(x22))
	x24 := (x12 + x23)
	x25 := (x11 + x24)
	x26 := (x25 & 0x3ffffff)
	x27 := uint8((x25 >> 26))
	x28 := (x10 + uint32(x27))
	x29 := (x9 + x28)
	x30 := (x8 + x29)
	x31 := (x30 & 0x3ffffff)
	x32 := uint8((x30 >> 26))
	x33 := (x7 + uint32(x32))
	x34 := (x6 + x33)
	x35 := (x5 + x34)
	x36 := (x3 + uint32(x4))
	x37 := (x2 + x36)
	x38 := (x1 + x37)
	out1[0] = x21
	out1[1] = x26
	out1[2] = x31
	out1[3] = x35
	out1[4] = x38
}

// Relax is the identity function converting from tight field elements to loose field elements.
//
// Postconditions:
//   out1 = arg1
//
func Relax(out1 *LooseFieldElement, arg1 *TightFieldElement) {
	x1 := arg1[0]
	x2 := arg1[1]
	x3 := arg1[2]
	x4 := arg1[3]
	x5 := arg1[4]
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
}

// CarryAdd adds two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 + eval arg2) mod m
//
func CarryAdd(out1 *TightFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := (arg1[0] + arg2[0])
	x2 := ((x1 >> 26) + (arg1[1] + arg2[1]))
	x3 := ((x2 >> 26) + (arg1[2] + arg2[2]))
	x4 := ((x3 >> 26) + (arg1[3] + arg2[3]))
	x5 := ((x4 >> 26) + (arg1[4] + arg2[4]))
	x6 := ((x1 & 0x3ffffff) + ((x5 >> 26) * 0x5))
	x7 := (uint32(uint1((x6 >> 26))) + (x2 & 0x3ffffff))
	x8 := (x6 & 0x3ffffff)
	x9 := (x7 & 0x3ffffff)
	x10 := (uint32(uint1((x7 >> 26))) + (x3 & 0x3ffffff))
	x11 := (x4 & 0x3ffffff)
	x12 := (x5 & 0x3ffffff)
	out1[0] = x8
	out1[1] = x9
	out1[2] = x10
	out1[3] = x11
	out1[4] = x12
}

// CarrySub subtracts two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 - eval arg2) mod m
//
func CarrySub(out1 *TightFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := ((0x7fffff6 + arg1[0]) - arg2[0])
	x2 := ((x1 >> 26) + ((0x7fffffe + arg1[1]) - arg2[1]))
	x3 := ((x2 >> 26) + ((0x7fffffe + arg1[2]) - arg2[2]))
	x4 := ((x3 >> 26) + ((0x7fffffe + arg1[3]) - arg2[3]))
	x5 := ((x4 >> 26) + ((0x7fffffe + arg1[4]) - arg2[4]))
	x6 := ((x1 & 0x3ffffff) + ((x5 >> 26) * 0x5))
	x7 := (uint32(uint1((x6 >> 26))) + (x2 & 0x3ffffff))
	x8 := (x6 & 0x3ffffff)
	x9 := (x7 & 0x3ffffff)
	x10 := (uint32(uint1((x7 >> 26))) + (x3 & 0x3ffffff))
	x11 := (x4 & 0x3ffffff)
	x12 := (x5 & 0x3ffffff)
	out1[0] = x8
	out1[1] = x9
	out1[2] = x10
	out1[3] = x11
	out1[4] = x12
}

// CarryOpp negates a field element.
//
// Postconditions:
//   eval out1 mod m = -eval arg1 mod m
//
func CarryOpp(out1 *TightFieldElement, arg1 *TightFieldElement) {
	x1 := (0x7fffff6 - arg1[0])
	x2 := (uint32(uint1((x1 >> 26))) + (0x7fffffe - arg1[1]))
	x3 := (uint32(uint1((x2 >> 26))) + (0x7fffffe - arg1[2]))
	x4 := (uint32(uint1((x3 >> 26))) + (0x7fffffe - arg1[3]))
	x5 := (uint32(uint1((x4 >> 26))) + (0x7fffffe - arg1[4]))
	x6 := ((x1 & 0x3ffffff) + (uint32(uint1((x5 >> 26))) * 0x5))
	x7 := (uint32(uint1((x6 >> 26))) + (x2 & 0x3ffffff))
	x8 := (x6 & 0x3ffffff)
	x9 := (x7 & 0x3ffffff)
	x10 := (uint32(uint1((x7 >> 26))) + (x3 & 0x3ffffff))
	x11 := (x4 & 0x3ffffff)
	x12 := (x5 & 0x3ffffff)
	out1[0] = x8
	out1[1] = x9
	out1[2] = x10
	out1[3] = x11
	out1[4] = x12
}
