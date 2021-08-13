package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/x-punch/go-redis"
)

func TestGet(t *testing.T) {
	ctx, key, value := context.TODO(), "noexist", "value"
	if _, err := srv.GetBytes(ctx, key); err != redis.ErrKeyNotExist {
		t.Fail()
	}
	if err := srv.Set(ctx, key, value, 5*time.Second); err != nil {
		t.Fatal(err)
	}
	defer func() {
		srv.Del(ctx, key)
	}()
	if v, err := srv.GetBytes(ctx, key); err != nil {
		t.Fatal(err)
	} else if string(v) != value {
		t.Fail()
	}
	if v, err := srv.GetString(ctx, key); err != nil {
		t.Fatal(err)
	} else if v != value {
		t.Fail()
	}
	srv.Del(ctx, key)
	if _, err := srv.GetString(ctx, key); err != redis.ErrKeyNotExist {
		t.Fail()
	}
}

func TestSetNX(t *testing.T) {
	ctx, key, value := context.TODO(), "key_nx", "value1"
	if exists, err := srv.Exist(ctx, key); err != nil {
		t.Fatal(err)
	} else if exists {
		t.Fail()
	}
	if ok, err := srv.SetNX(ctx, key, value, 5*time.Second); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fail()
	}
	defer func() {
		srv.Del(ctx, key)
	}()
	if count, err := srv.Exists(ctx, key); err != nil {
		t.Fatal(err)
	} else if count != 1 {
		t.Fail()
	}
	if ok, err := srv.SetNX(ctx, key, value, 5*time.Second); err != nil {
		t.Fatal(err)
	} else if ok {
		t.Fail()
	}
}

func TestSetXX(t *testing.T) {
	ctx, key, value := context.TODO(), "key_xx", "value1"
	if exists, err := srv.Exist(ctx, key); err != nil {
		t.Fatal(err)
	} else if exists {
		t.Fail()
	}
	if ok, err := srv.SetXX(ctx, key, value, 5*time.Second); err != nil {
		t.Fatal(err)
	} else if ok {
		t.Fail()
	}
	if err := srv.Set(ctx, key, value, 5*time.Second); err != nil {
		t.Fatal(err)
	}
	defer func() {
		srv.Del(ctx, key)
	}()
	if count, err := srv.Exists(ctx, key); err != nil {
		t.Fatal(err)
	} else if count != 1 {
		t.Fail()
	}
	if ok, err := srv.SetXX(ctx, key, value, 5*time.Second); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fail()
	}
}
