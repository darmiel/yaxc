package common

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

var valid = []string{
	"hello",                  // lower case
	"Hello",                  // Captial case
	"HELLO",                  // UPPER CASE
	strings.Repeat("A", 3),   // Min Length
	strings.Repeat("A", 128), // Max Length
	"Hel:lo!",                // Special Chars
	".-_+*!$%~@",             // More Special Chars
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:.-_+*!$%~@", // All valid chars
	"0123456789", // numbers
}

var invalid = []string{
	strings.Repeat("A", 2),   // < Min Length
	strings.Repeat("A", 129), // > Max Length
	strings.Repeat("#", 6),   // Special chars
}

func BenchmarkValidateAnywherePath(b *testing.B) {
	log.Println("# Iter:", b.N, "*", len(valid)+len(invalid))
	for n := 0; n < b.N; n++ {
		for _, v := range valid {
			ValidateAnywherePath(v)
		}
		for _, v := range invalid {
			ValidateAnywherePath(v)
		}
	}
}

func BenchmarkValidateAnywherePathRegex(b *testing.B) {
	log.Println("# Iter:", b.N, "*", len(valid)+len(invalid))
	for n := 0; n < b.N; n++ {
		for _, v := range valid {
			ValidateAnywherePathRegex(v)
		}
		for _, v := range invalid {
			ValidateAnywherePathRegex(v)
		}
	}
}

func TestValidateAnywherePath(t *testing.T) {

	// check valid paths
	for _, v := range valid {
		AssertTrue(t, ValidateAnywherePath(v))
		AssertTrue(t, ValidateAnywherePathRegex(v))
	}

	// check invalid paths
	for _, i := range invalid {
		AssertFalse(t, ValidateAnywherePath(i))
		AssertFalse(t, ValidateAnywherePathRegex(i))
	}
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
func AssertTrue(t *testing.T, a interface{}) {
	AssertEqual(t, a, true)
}
func AssertFalse(t *testing.T, a interface{}) {
	AssertEqual(t, a, false)
}
