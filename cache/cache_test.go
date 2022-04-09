package cache

import "testing"

func TestRedisCache_Has(t *testing.T) {
	if err := testRedisCache.Forget("test_has"); err != nil {
		t.Errorf("error while forgetting the key: %s", err)
	}

	found, err := testRedisCache.Has("test_has")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if found {
		t.Errorf("key %q is found when it should not be", "test_has")
	}

	if err := testRedisCache.Set("test_has", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	found, err = testRedisCache.Has("test_has")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if !found {
		t.Errorf("key %q is not found when it should be", "test_has")
	}
}

func TestRedisCache_Get(t *testing.T) {
	if err := testRedisCache.Set("test_get", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	val, err := testRedisCache.Get("test_get")
	if err != nil {
		t.Errorf("error while getting the key: %s", err)
	}

	if val != "test_val" {
		t.Errorf("wrong value, expected %s, got %s", "test_val", val)
	}
}

func TestRedisCache_Forget(t *testing.T) {
	if err := testRedisCache.Set("test_forget", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testRedisCache.Forget("test_forget"); err != nil {
		t.Errorf("error while forgetting the key: %s", err)
	}

	found, err := testRedisCache.Has("test_forget")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if found {
		t.Errorf("key %q is found when it should not be", "test_forget")
	}
}

func TestRedisCache_Empty(t *testing.T) {
	if err := testRedisCache.Set("test_empty", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testRedisCache.Empty(); err != nil {
		t.Errorf("error while emptying: %s", err)
	}

	found, err := testRedisCache.Has("test_empty")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if found {
		t.Errorf("key %q is found when it should not be", "test_empty")
	}
}

func TestRedisCache_EmptyByMatch(t *testing.T) {
	if err := testRedisCache.Set("test_empty_bm1", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testRedisCache.Set("test_empty_bm2", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testRedisCache.Set("test_keep", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testRedisCache.EmptyByMatch("test_empty_bm"); err != nil {
		t.Errorf("error while emptying by match: %s", err)
	}

	found, err := testRedisCache.Has("test_empty_bm1")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}
	if found {
		t.Errorf("key %q is found when it should not be", "test_empty_bm1")
	}

	found, err = testRedisCache.Has("test_empty_bm2")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}
	if found {
		t.Errorf("key %q is found when it should not be", "test_empty_bm2")
	}

	found, err = testRedisCache.Has("test_keep")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}
	if !found {
		t.Errorf("key %q is not found when it should be", "test_keep")
	}
}

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
