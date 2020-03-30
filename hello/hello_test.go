package hello_test

import (
	"cuishuaijie/go-examples/hello"
	"testing"
)

func TestHello(t *testing.T )  {
	ret := hello.Hello()
	if ret != "hello"{
		t.Errorf("test error,expect :%s,really is %s\n","hello",ret)
	}
}
