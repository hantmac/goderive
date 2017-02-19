//DO NOT EDIT generated by goderive

package test

import (
	"bytes"
)

func deriveEqualSliceOfint(this, that []int) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfintToint(this, that map[int]int) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualPtrToint(this, that *int) bool {
	return ((this == nil && that == nil) || (this != nil && that != nil && *this == *that))
}

func deriveEqualPtrToSliceOfint(this, that *[]int) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) && deriveEqualSliceOfint(*this, *that)
}

func deriveEqualPtrToArray10Ofint(this, that *[10]int) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) && deriveEqualArray10Ofint(*this, *that)
}

func deriveEqualPtrToMapOfintToint(this, that *map[int]int) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) && deriveEqualMapOfintToint(*this, *that)
}

func deriveEqualPtrToBuiltInTypes(this, that *BuiltInTypes) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Bool == that.Bool &&
		this.Byte == that.Byte &&
		this.Complex128 == that.Complex128 &&
		this.Complex64 == that.Complex64 &&
		this.Float64 == that.Float64 &&
		this.Float32 == that.Float32 &&
		this.Int == that.Int &&
		this.Int16 == that.Int16 &&
		this.Int32 == that.Int32 &&
		this.Int64 == that.Int64 &&
		this.Int8 == that.Int8 &&
		this.Rune == that.Rune &&
		this.String == that.String &&
		this.Uint == that.Uint &&
		this.Uint16 == that.Uint16 &&
		this.Uint32 == that.Uint32 &&
		this.Uint64 == that.Uint64 &&
		this.Uint8 == that.Uint8 &&
		this.UintPtr == that.UintPtr
}

func deriveEqualPtrToPtrToBuiltInTypes(this, that *PtrToBuiltInTypes) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		((this.Bool == nil && that.Bool == nil) || (this.Bool != nil && that.Bool != nil && *this.Bool == *that.Bool)) &&
		((this.Byte == nil && that.Byte == nil) || (this.Byte != nil && that.Byte != nil && *this.Byte == *that.Byte)) &&
		((this.Complex128 == nil && that.Complex128 == nil) || (this.Complex128 != nil && that.Complex128 != nil && *this.Complex128 == *that.Complex128)) &&
		((this.Complex64 == nil && that.Complex64 == nil) || (this.Complex64 != nil && that.Complex64 != nil && *this.Complex64 == *that.Complex64)) &&
		((this.Float64 == nil && that.Float64 == nil) || (this.Float64 != nil && that.Float64 != nil && *this.Float64 == *that.Float64)) &&
		((this.Float32 == nil && that.Float32 == nil) || (this.Float32 != nil && that.Float32 != nil && *this.Float32 == *that.Float32)) &&
		((this.Int == nil && that.Int == nil) || (this.Int != nil && that.Int != nil && *this.Int == *that.Int)) &&
		((this.Int16 == nil && that.Int16 == nil) || (this.Int16 != nil && that.Int16 != nil && *this.Int16 == *that.Int16)) &&
		((this.Int32 == nil && that.Int32 == nil) || (this.Int32 != nil && that.Int32 != nil && *this.Int32 == *that.Int32)) &&
		((this.Int64 == nil && that.Int64 == nil) || (this.Int64 != nil && that.Int64 != nil && *this.Int64 == *that.Int64)) &&
		((this.Int8 == nil && that.Int8 == nil) || (this.Int8 != nil && that.Int8 != nil && *this.Int8 == *that.Int8)) &&
		((this.Rune == nil && that.Rune == nil) || (this.Rune != nil && that.Rune != nil && *this.Rune == *that.Rune)) &&
		((this.String == nil && that.String == nil) || (this.String != nil && that.String != nil && *this.String == *that.String)) &&
		((this.Uint == nil && that.Uint == nil) || (this.Uint != nil && that.Uint != nil && *this.Uint == *that.Uint)) &&
		((this.Uint16 == nil && that.Uint16 == nil) || (this.Uint16 != nil && that.Uint16 != nil && *this.Uint16 == *that.Uint16)) &&
		((this.Uint32 == nil && that.Uint32 == nil) || (this.Uint32 != nil && that.Uint32 != nil && *this.Uint32 == *that.Uint32)) &&
		((this.Uint64 == nil && that.Uint64 == nil) || (this.Uint64 != nil && that.Uint64 != nil && *this.Uint64 == *that.Uint64)) &&
		((this.Uint8 == nil && that.Uint8 == nil) || (this.Uint8 != nil && that.Uint8 != nil && *this.Uint8 == *that.Uint8)) &&
		((this.UintPtr == nil && that.UintPtr == nil) || (this.UintPtr != nil && that.UintPtr != nil && *this.UintPtr == *that.UintPtr))
}

func deriveEqualPtrToSliceOfBuiltInTypes(this, that *SliceOfBuiltInTypes) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		deriveEqualSliceOfbool(this.Bool, that.Bool) &&
		bytes.Equal(this.Byte, that.Byte) &&
		deriveEqualSliceOfcomplex128(this.Complex128, that.Complex128) &&
		deriveEqualSliceOfcomplex64(this.Complex64, that.Complex64) &&
		deriveEqualSliceOffloat64(this.Float64, that.Float64) &&
		deriveEqualSliceOffloat32(this.Float32, that.Float32) &&
		deriveEqualSliceOfint(this.Int, that.Int) &&
		deriveEqualSliceOfint16(this.Int16, that.Int16) &&
		deriveEqualSliceOfint32(this.Int32, that.Int32) &&
		deriveEqualSliceOfint64(this.Int64, that.Int64) &&
		deriveEqualSliceOfint8(this.Int8, that.Int8) &&
		deriveEqualSliceOfrune(this.Rune, that.Rune) &&
		deriveEqualSliceOfstring(this.String, that.String) &&
		deriveEqualSliceOfuint(this.Uint, that.Uint) &&
		deriveEqualSliceOfuint16(this.Uint16, that.Uint16) &&
		deriveEqualSliceOfuint32(this.Uint32, that.Uint32) &&
		deriveEqualSliceOfuint64(this.Uint64, that.Uint64) &&
		bytes.Equal(this.Uint8, that.Uint8) &&
		deriveEqualSliceOfuintptr(this.UintPtr, that.UintPtr)
}

func deriveEqualPtrToSliceOfPtrToBuiltInTypes(this, that *SliceOfPtrToBuiltInTypes) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		deriveEqualSliceOfPtrTobool(this.Bool, that.Bool) &&
		deriveEqualSliceOfPtrTobyte(this.Byte, that.Byte) &&
		deriveEqualSliceOfPtrTocomplex128(this.Complex128, that.Complex128) &&
		deriveEqualSliceOfPtrTocomplex64(this.Complex64, that.Complex64) &&
		deriveEqualSliceOfPtrTofloat64(this.Float64, that.Float64) &&
		deriveEqualSliceOfPtrTofloat32(this.Float32, that.Float32) &&
		deriveEqualSliceOfPtrToint(this.Int, that.Int) &&
		deriveEqualSliceOfPtrToint16(this.Int16, that.Int16) &&
		deriveEqualSliceOfPtrToint32(this.Int32, that.Int32) &&
		deriveEqualSliceOfPtrToint64(this.Int64, that.Int64) &&
		deriveEqualSliceOfPtrToint8(this.Int8, that.Int8) &&
		deriveEqualSliceOfPtrTorune(this.Rune, that.Rune) &&
		deriveEqualSliceOfPtrTostring(this.String, that.String) &&
		deriveEqualSliceOfPtrTouint(this.Uint, that.Uint) &&
		deriveEqualSliceOfPtrTouint16(this.Uint16, that.Uint16) &&
		deriveEqualSliceOfPtrTouint32(this.Uint32, that.Uint32) &&
		deriveEqualSliceOfPtrTouint64(this.Uint64, that.Uint64) &&
		deriveEqualSliceOfPtrTouint8(this.Uint8, that.Uint8) &&
		deriveEqualSliceOfPtrTouintptr(this.UintPtr, that.UintPtr)
}

func deriveEqualPtrToArrayOfBuiltInTypes(this, that *ArrayOfBuiltInTypes) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Bool == that.Bool &&
		this.Byte == that.Byte &&
		this.Complex128 == that.Complex128 &&
		this.Complex64 == that.Complex64 &&
		this.Float64 == that.Float64 &&
		this.Float32 == that.Float32 &&
		this.Int == that.Int &&
		this.Int16 == that.Int16 &&
		this.Int32 == that.Int32 &&
		this.Int64 == that.Int64 &&
		this.Int8 == that.Int8 &&
		this.Rune == that.Rune &&
		this.String == that.String &&
		this.Uint == that.Uint &&
		this.Uint16 == that.Uint16 &&
		this.Uint32 == that.Uint32 &&
		this.Uint64 == that.Uint64 &&
		this.Uint8 == that.Uint8 &&
		this.UintPtr == that.UintPtr &&
		this.AnotherBoolOfDifferentSize == that.AnotherBoolOfDifferentSize
}

func deriveEqualPtrToArrayOfPtrToBuiltInTypes(this, that *ArrayOfPtrToBuiltInTypes) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		deriveEqualArray1OfPtrTobool(this.Bool, that.Bool) &&
		deriveEqualArray2OfPtrTobyte(this.Byte, that.Byte) &&
		deriveEqualArray3OfPtrTocomplex128(this.Complex128, that.Complex128) &&
		deriveEqualArray4OfPtrTocomplex64(this.Complex64, that.Complex64) &&
		deriveEqualArray5OfPtrTofloat64(this.Float64, that.Float64) &&
		deriveEqualArray6OfPtrTofloat32(this.Float32, that.Float32) &&
		deriveEqualArray7OfPtrToint(this.Int, that.Int) &&
		deriveEqualArray8OfPtrToint16(this.Int16, that.Int16) &&
		deriveEqualArray9OfPtrToint32(this.Int32, that.Int32) &&
		deriveEqualArray10OfPtrToint64(this.Int64, that.Int64) &&
		deriveEqualArray11OfPtrToint8(this.Int8, that.Int8) &&
		deriveEqualArray12OfPtrTorune(this.Rune, that.Rune) &&
		deriveEqualArray13OfPtrTostring(this.String, that.String) &&
		deriveEqualArray14OfPtrTouint(this.Uint, that.Uint) &&
		deriveEqualArray15OfPtrTouint16(this.Uint16, that.Uint16) &&
		deriveEqualArray16OfPtrTouint32(this.Uint32, that.Uint32) &&
		deriveEqualArray17OfPtrTouint64(this.Uint64, that.Uint64) &&
		deriveEqualArray18OfPtrTouint8(this.Uint8, that.Uint8) &&
		deriveEqualArray19OfPtrTouintptr(this.UintPtr, that.UintPtr) &&
		deriveEqualArray10OfPtrTobool(this.AnotherBoolOfDifferentSize, that.AnotherBoolOfDifferentSize)
}

func deriveEqualPtrToMapsOfBuiltInTypes(this, that *MapsOfBuiltInTypes) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		deriveEqualMapOfboolTostring(this.BoolToString, that.BoolToString) &&
		deriveEqualMapOfstringTobool(this.StringToBool, that.StringToBool) &&
		deriveEqualMapOfcomplex128Tocomplex64(this.Complex128ToComplex64, that.Complex128ToComplex64) &&
		deriveEqualMapOffloat64Touint32(this.Float64ToUint32, that.Float64ToUint32) &&
		deriveEqualMapOfuint16Touint8(this.Uint16ToUint8, that.Uint16ToUint8)
}

func deriveEqualPtrToSliceToSlice(this, that *SliceToSlice) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		deriveEqualSliceOfSliceOfint(this.Ints, that.Ints) &&
		deriveEqualSliceOfSliceOfstring(this.Strings, that.Strings) &&
		deriveEqualSliceOfSliceOfPtrToint(this.IntPtrs, that.IntPtrs)
}

func deriveEqualPtrToPtrTo(this, that *PtrTo) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		((this.Basic == nil && that.Basic == nil) || (this.Basic != nil && that.Basic != nil && *this.Basic == *that.Basic)) &&
		((this.Slice == nil && that.Slice == nil) || (this.Slice != nil && that.Slice != nil && deriveEqualSliceOfint(*this.Slice, *that.Slice))) &&
		((this.Array == nil && that.Array == nil) || (this.Array != nil && that.Array != nil && *this.Array == *that.Array)) &&
		((this.Map == nil && that.Map == nil) || (this.Map != nil && that.Map != nil && deriveEqualMapOfintToint(*this.Map, *that.Map)))
}

func deriveEqualPtrToName(this, that *Name) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Name == that.Name
}

func deriveEqualPtrToStructs(this, that *Structs) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Struct == that.Struct &&
		this.PtrToStruct.Equal(that.PtrToStruct) &&
		deriveEqualSliceOfName(this.SliceOfStructs, that.SliceOfStructs) &&
		deriveEqualSliceOfPtrToName(this.SliceToPtrOfStruct, that.SliceToPtrOfStruct)
}

func deriveEqualPtrToMapWithStructs(this, that *MapWithStructs) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		deriveEqualMapOfNameTostring(this.NameToString, that.NameToString) &&
		deriveEqualMapOfstringToName(this.StringToName, that.StringToName) &&
		deriveEqualMapOfstringToPtrToName(this.StringToPtrToName, that.StringToPtrToName) &&
		deriveEqualMapOfstringToSliceOfName(this.StringToSliceOfName, that.StringToSliceOfName) &&
		deriveEqualMapOfstringToSliceOfPtrToName(this.StringToSliceOfPtrToName, that.StringToSliceOfPtrToName)
}

func deriveEqualPtrToRecursiveType(this, that *RecursiveType) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		bytes.Equal(this.Bytes, that.Bytes) &&
		deriveEqualMapOfintToRecursiveType(this.N, that.N)
}

func deriveEqualPtrToEmbeddedStruct1(this, that *EmbeddedStruct1) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Name == that.Name &&
		this.Structs.Equal(that.Structs)
}

func deriveEqualPtrToEmbeddedStruct2(this, that *EmbeddedStruct2) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Structs.Equal(&that.Structs) &&
		this.Name.Equal(that.Name)
}

func deriveEqualPtrToUnnamedStruct(this, that *UnnamedStruct) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Unnamed == that.Unnamed
}

func deriveEqualArray10Ofint(this, that [10]int) bool {
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfbool(this, that []bool) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfcomplex128(this, that []complex128) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfcomplex64(this, that []complex64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOffloat64(this, that []float64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOffloat32(this, that []float32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfint16(this, that []int16) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfint32(this, that []int32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfint64(this, that []int64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfint8(this, that []int8) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfrune(this, that []rune) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfstring(this, that []string) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfuint(this, that []uint) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfuint16(this, that []uint16) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfuint32(this, that []uint32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfuint64(this, that []uint64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfuintptr(this, that []uintptr) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTobool(this, that []*bool) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTobyte(this, that []*byte) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTocomplex128(this, that []*complex128) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTocomplex64(this, that []*complex64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTofloat64(this, that []*float64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTofloat32(this, that []*float32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrToint(this, that []*int) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrToint16(this, that []*int16) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrToint32(this, that []*int32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrToint64(this, that []*int64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrToint8(this, that []*int8) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTorune(this, that []*rune) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTostring(this, that []*string) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTouint(this, that []*uint) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTouint16(this, that []*uint16) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTouint32(this, that []*uint32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTouint64(this, that []*uint64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTouint8(this, that []*uint8) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrTouintptr(this, that []*uintptr) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray1OfPtrTobool(this, that [1]*bool) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray2OfPtrTobyte(this, that [2]*byte) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray3OfPtrTocomplex128(this, that [3]*complex128) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray4OfPtrTocomplex64(this, that [4]*complex64) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray5OfPtrTofloat64(this, that [5]*float64) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray6OfPtrTofloat32(this, that [6]*float32) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray7OfPtrToint(this, that [7]*int) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray8OfPtrToint16(this, that [8]*int16) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray9OfPtrToint32(this, that [9]*int32) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray10OfPtrToint64(this, that [10]*int64) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray11OfPtrToint8(this, that [11]*int8) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray12OfPtrTorune(this, that [12]*rune) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray13OfPtrTostring(this, that [13]*string) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray14OfPtrTouint(this, that [14]*uint) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray15OfPtrTouint16(this, that [15]*uint16) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray16OfPtrTouint32(this, that [16]*uint32) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray17OfPtrTouint64(this, that [17]*uint64) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray18OfPtrTouint8(this, that [18]*uint8) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray19OfPtrTouintptr(this, that [19]*uintptr) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualArray10OfPtrTobool(this, that [10]*bool) bool {
	for i := 0; i < len(this); i++ {
		if !((this[i] == nil && that[i] == nil) || (this[i] != nil && that[i] != nil && *this[i] == *that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfboolTostring(this, that map[bool]string) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfstringTobool(this, that map[string]bool) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfcomplex128Tocomplex64(this, that map[complex128]complex64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualMapOffloat64Touint32(this, that map[float64]uint32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfuint16Touint8(this, that map[uint16]uint8) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfSliceOfint(this, that [][]int) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(deriveEqualSliceOfint(this[i], that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfSliceOfstring(this, that [][]string) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(deriveEqualSliceOfstring(this[i], that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfSliceOfPtrToint(this, that [][]*int) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(deriveEqualSliceOfPtrToint(this[i], that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfName(this, that []Name) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true
}

func deriveEqualSliceOfPtrToName(this, that []*Name) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i].Equal(that[i])) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfNameTostring(this, that map[Name]string) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfstringToName(this, that map[string]Name) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfstringToPtrToName(this, that map[string]*Name) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v.Equal(thatv)) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfstringToSliceOfName(this, that map[string][]Name) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(deriveEqualSliceOfName(v, thatv)) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfstringToSliceOfPtrToName(this, that map[string][]*Name) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(deriveEqualSliceOfPtrToName(v, thatv)) {
			return false
		}
	}
	return true
}

func deriveEqualMapOfintToRecursiveType(this, that map[int]RecursiveType) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v.Equal(&thatv)) {
			return false
		}
	}
	return true
}
