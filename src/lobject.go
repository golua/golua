package main

type lu_byte int8

type GCObject struct {
	next   *GCObject
	tt     lu_byte
	marked lu_byte
}
