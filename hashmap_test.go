package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/x-punch/go-redis"
)

var srv redis.Service

func init() {
	srv = redis.NewService(redis.NewConfig())
}

func TestHGet(t *testing.T) {
	ctx, key, field, value := context.TODO(), "hget", "field", "value"
	if err := srv.HSet(ctx, key, map[string]interface{}{field: value}); err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	defer func() {
		srv.Del(ctx, key)
	}()
	if v, err := srv.HGet(ctx, key, field); err != nil {
		t.Fatal(err)
	} else if v != value {
		t.Fail()
	}
}

func TestHDel(t *testing.T) {
	ctx, key, field, value := context.TODO(), "hdel", "field", "value"
	if err := srv.HSet(ctx, key, field, value); err != nil {
		t.Fatal(err)
	}
	defer func() {
		srv.Del(ctx, key)
	}()
	if exists, err := srv.HExists(ctx, key, field); err != nil {
		t.Fatal(err)
	} else if !exists {
		t.Fail()
	}
	if err := srv.HDel(ctx, key, field); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.HGet(ctx, key, field); err != redis.ErrKeyNotExist {
		t.Fail()
	}
}

func TestHGetAll(t *testing.T) {
	ctx, key, fields, values := context.TODO(), "hdel", []string{"field1", "field2", "field3"}, []string{"value1", "value2", "value3"}
	for i := 0; i < len(fields); i++ {
		if err := srv.HSet(ctx, key, fields[i], values[i]); err != nil {
			t.Fatal(err)
		}
	}
	defer func() {
		srv.Del(ctx, key)
	}()
	if result, err := srv.HKeys(ctx, key); err != nil {
		t.Fatal(err)
	} else {
		if len(result) != len(fields) {
			t.Fail()
		}
	}
	if n, err := srv.HLen(ctx, key); err != nil {
		t.Fatal(err)
	} else if n != int64(len(fields)) {
		t.Fail()
	}
	if all, err := srv.HGetAll(ctx, key); err != nil {
		t.Fatal(err)
	} else {
		if len(all) != len(fields) {
			t.Fail()
		}
		for i := 0; i < len(fields); i++ {
			if v, ok := all[fields[i]]; !ok {
				t.Fail()
			} else if v != values[i] {
				t.Fail()
			}
		}
	}
	if all, err := srv.HMGet(ctx, key, fields...); err != nil {
		t.Fatal(err)
	} else {
		if len(all) != len(fields) {
			t.Fail()
		}
		for i := 0; i < len(fields); i++ {
			if all[i] != values[i] {
				t.Fail()
			}
		}
	}
}
