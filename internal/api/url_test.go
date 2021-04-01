package api

import (
	"reflect"
	"testing"
	"time"
)

var a *Api

func init() {
	a = &Api{ServerURL: "https://test"}
}

func TestApi_UrlSetContents(t *testing.T) {
	AssertEqual(t, a.UrlSetContents("linus", "", "", 0), "https://test/linus")
	AssertEqual(t, a.UrlSetContents("linus", "deadbeef", "", 0), "https://test/linus/deadbeef")
	AssertEqual(t, a.UrlSetContents("linus", "", "secret", 0), "https://test/linus?secret=secret")
	AssertEqual(t, a.UrlSetContents("linus", "deadbeef", "secret", 0), "https://test/linus/deadbeef?secret=secret")
	AssertEqual(t, a.UrlSetContents("linus", "", "", 10*time.Minute), "https://test/linus?ttl=10m0s")
	AssertEqual(t, a.UrlSetContents("linus", "", "secret", 10*time.Minute), "https://test/linus?secret=secret&ttl=10m0s")
}

func TestApi_UrlGetContents(t *testing.T) {
	AssertEqual(t, a.UrlGetContents("linus", ""), "https://test/linus")
	AssertEqual(t, a.UrlGetContents("linus", "test"), "https://test/linus?secret=test")
}

func TestApi_UrlGetHash(t *testing.T) {
	AssertEqual(t, a.UrlGetHash("linus"), "https://test/hash/linus")
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
