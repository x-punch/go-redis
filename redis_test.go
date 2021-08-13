package redis_test

import (
	"context"
	"testing"
	"time"
)

func TestTTL(t *testing.T) {
	ctx, key, value, exp := context.TODO(), "ttl", "value", 5*time.Second
	if err := srv.Set(ctx, key, value, 0); err != nil {
		t.Fatal(err)
	}
	if ok, err := srv.Expire(ctx, key, exp); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fail()
	}
	if dur, err := srv.TTL(ctx, key); err != nil {
		t.Fatal(err)
	} else if dur > exp {
		t.Fail()
	}
	exp2 := time.Second
	if ok, err := srv.ExpireAt(ctx, key, time.Now().Add(exp2)); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fail()
	}
	if dur, err := srv.TTL(ctx, key); err != nil {
		t.Fatal(err)
	} else if dur > exp2 {
		t.Fail()
	}
}
