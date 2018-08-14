package goluajit

// #cgo CFLAGS: -pagezero_size 10000 -image_base 100000000
// #cgo LDFLAGS: -lluajit -lm -ldl
import "C"
