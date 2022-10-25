package ed2k

import (
	"testing"
)

func TestZeroOneChunk(t *testing.T) {
	zf := make([]byte, 9728000)
	h := New()
	h.setTest(t)

	verifyWrite(t, h, 1, zf)

	verifyBlueRed(
		t, h,
		"d7def262a127cd79096a108e7a9fc138",
		"fc21d9af828f92a8df64beac3357425d",
	)
}

func TestZeroTwoChunk(t *testing.T) {
	zf := make([]byte, 9728000)
	h := New()
	h.setTest(t)

	verifyWrite(t, h, 1, zf)
	verifyWrite(t, h, 2, zf)

	verifyBlueRed(
		t, h,
		"194ee9e4fa79b2ee9f8829284c466051",
		"114b21c63a74b6ca922291a11177dd5c",
	)
}

func verifyWrite(t *testing.T, h *Ed2k, i int, b []byte) {
	t.Helper()

	n, err := h.Write(b)
	if err != nil {
		t.Logf("Write %v failed: %v", i, err)
		if n != len(b) {
			t.Logf("Bad size: expected %v, got %v", len(b), n)
		}
		t.FailNow()
	}
	if n != len(b) {
		t.Fatalf("Write %v: bad size: expected %d, got %d", i, len(b), n)
	}
}

func verifyBlueRed(t *testing.T, h *Ed2k, expectBlue, expectRed string) {
	t.Helper()

	gotBlue, errBlue := h.SumBlue()
	gotRed, errRed := h.SumRed()

	if errBlue != nil {
		t.Errorf("SumBlue error: %v", errBlue)
	}
	if errRed != nil {
		t.Errorf("SumRed error: %v", errRed)
	}

	if gotBlue != expectBlue {
		t.Errorf("blue failed: %s", gotBlue)
		t.Errorf("   expected: %s", expectBlue)
	} else {
		t.Logf("blue passed: %s", gotBlue)
	}

	if gotRed != expectRed {
		t.Errorf("red  failed: %s", gotRed)
		t.Errorf("   expected: %s", expectRed)
	} else {
		t.Logf("red  passed: %s", gotRed)
	}
}
