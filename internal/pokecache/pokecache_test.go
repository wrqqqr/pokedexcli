package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)
	if cache.cache == nil {
		t.Error("cache is nil")
	}

}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val5"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)

		data, ok := cache.Get(cas.inputKey)

		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(data) != string(cas.inputVal) {
			t.Errorf("%s does not match %s", string(data), string(cas.inputVal))
			continue
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cache.Add("key1", []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")

	if ok {
		t.Errorf("%s found after interval, bad reap", "key1")
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cache.Add("key1", []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get("key1")

	if !ok {
		t.Errorf("%s should not have been reaped", "key1")
	}
}
