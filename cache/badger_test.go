package cache

import "testing"

func TestBadgerCache_Has(t *testing.T) {
	if err := testBadgerCache.Forget("test_has"); err != nil {
		t.Errorf("error while forgetting the key: %s", err)
	}

	found, err := testBadgerCache.Has("test_has")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if found {
		t.Errorf("key %q is found when it should not be", "test_has")
	}

	if err := testBadgerCache.Set("test_has", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	found, err = testBadgerCache.Has("test_has")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if !found {
		t.Errorf("key %q is not found when it should be", "test_has")
	}
}

func TestBadgerCache_Get(t *testing.T) {
	if err := testBadgerCache.Set("test_get", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	val, err := testBadgerCache.Get("test_get")
	if err != nil {
		t.Errorf("error while getting the key: %s", err)
	}

	if val != "test_val" {
		t.Errorf("wrong value, expected %s, got %s", "test_val", val)
	}
}

func TestBadgerCache_EmptyCache_Forget(t *testing.T) {
	if err := testBadgerCache.Set("test_forget", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testBadgerCache.Forget("test_forget"); err != nil {
		t.Errorf("error while forgetting the key: %s", err)
	}

	found, err := testBadgerCache.Has("test_forget")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if found {
		t.Errorf("key %q is found when it should not be", "test_forget")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	if err := testBadgerCache.Set("test_empty", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testBadgerCache.Empty(); err != nil {
		t.Errorf("error while emptying: %s", err)
	}

	found, err := testBadgerCache.Has("test_empty")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}

	if found {
		t.Errorf("key %q is found when it should not be", "test_empty")
	}
}

func TestBadgerCache_EmptyByMatch(t *testing.T) {
	if err := testBadgerCache.Set("test_empty_bm1", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testBadgerCache.Set("test_empty_bm2", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testBadgerCache.Set("test_keep", "test_val"); err != nil {
		t.Errorf("error while setting the key: %s", err)
	}

	if err := testBadgerCache.EmptyByMatch("test_empty_bm"); err != nil {
		t.Errorf("error while emptying by match: %s", err)
	}

	found, err := testBadgerCache.Has("test_empty_bm1")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}
	if found {
		t.Errorf("key %q is found when it should not be", "test_empty_bm1")
	}

	found, err = testBadgerCache.Has("test_empty_bm2")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}
	if found {
		t.Errorf("key %q is found when it should not be", "test_empty_bm2")
	}

	found, err = testBadgerCache.Has("test_keep")
	if err != nil {
		t.Errorf("error while checking the key: %s", err)
	}
	if !found {
		t.Errorf("key %q is not found when it should be", "test_keep")
	}
}
