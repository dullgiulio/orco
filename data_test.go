package main

import (
    "testing"
)

func TestRowDiff(t *testing.T) {
    r0 := makeRow("a", "b", "c")
    r1 := makeRow("a")

    if r0.diff(r1) != nil {
        t.Error("expected nil on uncomparable rows")
    }

    if r0.diff(r0) != nil {
        t.Error("expected nil on comparison with itself")
    }

    r2 := makeRow("a", "c", "b")
    d0 := r0.diff(r2)
    if d0 == nil {
        t.Fatal("expected non-nil diff between different rows")
    }
    d1 := makeRowDiff()
    d1[1] = "c"
    d1[2] = "b"
    for k, v := range d1 {
        if d0[k] != v {
            t.Errorf("expected value %s: %s has %s", k, v, d0[k])
        }
    }
}
