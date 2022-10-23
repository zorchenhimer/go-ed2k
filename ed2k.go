package ed2k

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/crypto/md4"
)

type Ed2k struct {
	hashes []byte // md4 hashes of chunks
	buff *bytes.Buffer
	t *testing.T
}

func New() *Ed2k {
	return &Ed2k{
		hashes: []byte{},
		buff: &bytes.Buffer{},
	}
}

func (h *Ed2k) setTest(t *testing.T) {
	h.t = t
}

func (h *Ed2k) Write(p []byte) (int, error) {
	n, err := h.buff.Write(p)
	if err != nil {
		return n, err
	}

	for h.buff.Len() >= h.BlockSize() {
		c := make([]byte, h.BlockSize())
		_, err = h.buff.Read(c)
		if err != nil {
			return 0, err
		}

		cmd4 := md4.New()
		_, err = cmd4.Write(c)
		if err != nil {
			return 0, err
		}

		h.hashes = append(h.hashes, cmd4.Sum([]byte{})...)
		//h.hashes = cmd4.Sum(h.hashes)
	}


	if h.buff.Len() > 0 {
		overflow := h.buff.Bytes()
		//h.buff = bytes.NewBuffer(h.buff.Bytes())
		h.buff.Reset()
		_, err := h.buff.Write(overflow)
		if err != nil {
			return 0, err
		}
	} else {
		h.buff.Reset()
	}

	return n, nil
}

func (h *Ed2k) Sum(b []byte) []byte {
	leftover, hashes, err := h.currentHash()
	if err != nil {
		panic(err)
	}

	if !leftover && len(hashes) == h.Size() {
		//return fmt.Sprintf("%x", hashes), nil
		return append(b, hashes...)
	}

	hsh := md4.New()
	_, err = hsh.Write(hashes)
	if err != nil {
		panic(err)
	}

	return hsh.Sum(b)
}

func (h *Ed2k) Reset() {
	h.buff.Reset()
	h.hashes = []byte{}
}

func (h *Ed2k) Size() int {
	return 16
}

func (h *Ed2k) BlockSize() int {
	return 9728000
}

func (h *Ed2k) currentHash() (bool, []byte, error) {
	if h.buff.Len() != 0 {
		b := h.buff.Bytes()
		cmd4 := md4.New()
		_, err := cmd4.Write(b)
		if err != nil {
			return true, nil, err
		}

		return true, append(h.hashes, cmd4.Sum([]byte{})...), nil
	}
	return false, h.hashes, nil
}

func (h *Ed2k) SumBlue() (string, error) {
	leftover, hashes, err := h.currentHash()
	if err != nil {
		return "", err
	}

	if !leftover && len(hashes) == h.Size() {
		return fmt.Sprintf("%x", hashes), nil
	}

	hsh := md4.New()
	if h.t != nil {
		h.t.Logf("bluehashes: %X", hashes)
	}
	_, err = hsh.Write(hashes)
	if err != nil {
		return "", err
	}

	bhash := hsh.Sum([]byte{})
	return fmt.Sprintf("%x", bhash), nil
}

// The "bugged" version of the hash.  See https://wiki.anidb.net/Ed2k-hash#How_is_an_ed2k_hash_calculated_exactly? for more info.
func (h *Ed2k) SumRed() (string, error) {
	leftover, hashes, err := h.currentHash()
	if err != nil {
		return "", err
	}

	hsh := md4.New()
	if !leftover {
		lsh := md4.New()
		_, err = lsh.Write([]byte{})
		if err != nil {
			return "", err
		}
		hashes = append(hashes, lsh.Sum([]byte{})...)
	}

	if h.t != nil {
		h.t.Logf("red hashes: %X", hashes)
	}

	_, err = hsh.Write(hashes)
	if err != nil {
		return "", err
	}

	bhash := hsh.Sum([]byte{})
	return fmt.Sprintf("%x", bhash), nil
}
