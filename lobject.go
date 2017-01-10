package golua

type lu_byte int8

type GCObject struct {
	next   *GCObject
	tt     lu_byte
	marked lu_byte
}

type Value interface {
}

type TValue struct {
	value_ Value
	tt_ int
}