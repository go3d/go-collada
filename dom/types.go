package cdom

import (
	"github.com/go-utils/unum"
)

//	Contains two bool values.
type Bool2 [2]bool

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Bool2) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains three bool values.
type Bool3 [3]bool

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Bool3) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains four bool values.
type Bool4 [4]bool

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Bool4) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains two float64 values.
type Float2 [2]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float2) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains four float64 values.
type Float2x2 [4]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional or two-dimensional indices.
func (me *Float2x2) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains six float64 values.
type Float2x3 [6]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float2x3) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains eight float64 values.
type Float2x4 [8]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float2x4) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains three float64 values.
type Float3 [3]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float3) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains six float64 values.
type Float3x2 [6]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float3x2) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains nine float64 values.
type Float3x3 [9]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float3x3) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains twelve float64 values.
type Float3x4 [12]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float3x4) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains four float64 values.
type Float4 [4]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float4) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains eight float64 values.
type Float4x2 [8]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float4x2) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains twelve float64 values.
type Float4x3 [12]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float4x3) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains sixteen float64 values.
type Float4x4 [16]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float4x4) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains seven float64 values.
type Float7 [7]float64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Float7) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains two int64 values.
type Int2 [2]int64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Int2) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains four int64 values.
type Int2x2 [4]int64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Int2x2) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains three int64 values.
type Int3 [3]int64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Int3) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains nine int64 values.
type Int3x3 [9]int64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Int3x3) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Contains four int64 values.
type Int4 [4]int64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Int4) AccessIndex(i, _ int) interface{} {
	return &me[i]
}

//	Contains sixteen int64 values.
type Int4x4 [16]int64

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *Int4x4) AccessIndex(i, j int) interface{} {
	if j >= 0 {
		i = index2D(i, j, len(me))
	}
	return &me[i]
}

//	Provides a bool value.
type ParamOrBool struct {
	//	The value provided if Param is empty.
	B bool

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}

//	Provides a float64 value.
type ParamOrFloat struct {
	//	The value provided if Param is empty.
	F float64

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}

//	Provides two float64 values.
type ParamOrFloat2 struct {
	//	The values provided if Param is empty.
	F Float2

	//	If set, refers to a previously defined parameter providing the values.
	Param RefParam
}

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *ParamOrFloat2) AccessIndex(i, _ int) interface{} {
	return &me.F[i]
}

//	Provides a int64 value.
type ParamOrInt struct {
	//	The value provided if Param is empty.
	I int64

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}

//	Provides a scoped float64 value.
type ParamOrSidFloat struct {
	//	The value provided if Param is empty.
	F SidFloat

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}

//	Provides a RefSid value.
type ParamOrRefSid struct {
	//	The value provided if Param is empty.
	Sr RefSid

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}

//	Provides a uint64 value.
type ParamOrUint struct {
	//	The value provided if Param is empty.
	U uint64

	//	If set, refers to a previously defined parameter providing the value.
	Param RefParam
}

//	A bool value that has a scoped identifier.
type SidBool struct {
	//	Sid
	HasSid

	//	The value.
	B bool
}

//	A float64 value that has a scoped identifier.
type SidFloat struct {
	//	Sid
	HasSid

	//	The value.
	F float64
}

//	Three float64 values that have a scoped identifier.
type SidFloat3 struct {
	//	Sid
	HasSid

	//	The values.
	F Float3
}

//	RefSidIndexer implementation.
//	Supports one-dimensional indices.
func (me *SidFloat3) AccessIndex(i, _ int) interface{} {
	return &me.F[i]
}

//	A string value that has a scoped identifier.
type SidString struct {
	//	Sid
	HasSid

	//	The value.
	S string
}

//	A 3D vector that has a scoped identifier.
type SidVec3 struct {
	//	Sid
	HasSid

	//	X, Y, Z
	unum.Vec3
}

//	RefSidFielder implementation.
//	Supported field names: "X", "Y", "Z".
func (me *SidVec3) AccessField(fn string) interface{} {
	switch fn {
	case "X":
		return &me.X
	case "Y":
		return &me.Y
	case "Z":
		return &me.Z
	}
	return nil
}

//	RefSidIndexer implementation.
//	Supports one-dimensional indices 0 through 2.
func (me *SidVec3) AccessIndex(i, _ int) interface{} {
	switch i {
	case 0:
		return &me.X
	case 1:
		return &me.Y
	case 2:
		return &me.Z
	}
	return nil
}
