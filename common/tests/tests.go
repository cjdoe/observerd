package tests

import "testing"

func InterruptIfError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
