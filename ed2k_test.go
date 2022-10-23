package ed2k

import (
	"testing"
)

func TestZeroOneChunk(t *testing.T) {
	zf := make([]byte, 9728000)
	h := New()
	h.setTest(t)
	n, err := h.Write(zf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 9728000 {
		t.Fatalf("Bad write size: %d", n)
	}

	blue, err := h.SumBlue()
	if err != nil {
		t.Fatal(err)
	}

	red, err := h.SumRed()
	if err != nil {
		t.Fatal(err)
	}

	if blue != "d7def262a127cd79096a108e7a9fc138" {
		t.Logf("blue failed: %s", blue)
		t.Fail()
	}

	if red != "fc21d9af828f92a8df64beac3357425d" {
		t.Logf("red failed: %s", red)
		t.Fail()
	}
}

func TestZeroTwoChunk(t *testing.T) {
	zf := make([]byte, 9728000)
	h := New()
	h.setTest(t)
	n, err := h.Write(zf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 9728000 {
		t.Fatalf("Bad 1st write size: %d", n)
	}

	n, err = h.Write(zf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 9728000 {
		t.Fatalf("Bad 2nd write size: %d", n)
	}

	blue, err := h.SumBlue()
	if err != nil {
		t.Fatal(err)
	}

	red, err := h.SumRed()
	if err != nil {
		t.Fatal(err)
	}

	if blue != "194ee9e4fa79b2ee9f8829284c466051" {
		t.Logf("blue failed: %s", blue)
		t.Fail()
	} else {
		t.Logf("blue passed: %s", blue)
	}

	if red != "114b21c63a74b6ca922291a11177dd5c" {
		t.Logf("red failed: %s", red)
		t.Fail()
	} else {
		t.Logf("red passed: %s", red)
	}
}
