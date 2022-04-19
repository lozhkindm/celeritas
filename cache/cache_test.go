package cache

import "testing"

func TestEncodeDecode(t *testing.T) {
	entry := Entry{}
	entry["encode"] = "decode"

	bytes, err := encode(entry)
	if err != nil {
		t.Errorf("error while encoding: %s", err)
	}

	decoded, err := decode(bytes)
	if err != nil {
		t.Errorf("error while decoding: %s", err)
	}

	v, ok := decoded["encode"]

	if !ok {
		t.Errorf("entry does not have an expected key %q", "encode")
	}

	if v != "decode" {
		t.Errorf("entry does not have an expected value %q in key %q", "decode", "encode")
	}
}
