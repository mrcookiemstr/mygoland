package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
)

func Test_MarshalJson(t *testing.T) {
	m := make(StringFloateMap)
	m["first"] = float64(math.NaN())
	m["second"] = float64(math.Pi)
	m["third"] = float64(math.Inf(1))
	m["fourth"] = float64(math.Inf(-1))

	marshalJson, marshalErr := m.MarshalJson()

	if marshalErr != nil {
		t.Errorf("Test_MarshalJson error %v", marshalErr)
	}

	resultStr := string(marshalJson[:])

	if strings.Contains(resultStr, "NaN") == false {
		t.Errorf("Test_MarshalJson has no NaN")
	}

	piStr := fmt.Sprintf("%f", math.Pi)

	if strings.Contains(resultStr, piStr) == false {
		t.Errorf("Test_MarshalJson has no NaN")
	}

	if strings.Contains(resultStr, "+Inf") == false {
		t.Errorf("Test_MarshalJson has no +Inf")
	}

	if strings.Contains(resultStr, "-Inf") == false {
		t.Errorf("Test_MarshalJson has no -Inf")
	}
}

func Test_UnmarshalJson(t *testing.T) {
	m := make(StringFloateMap)
	m["first"] = float64(math.NaN())
	m["second"] = float64(math.Pi)
	m["third"] = float64(math.Inf(1))
	m["fourth"] = float64(math.Inf(-1))

	marshalJson, marshalErr := m.MarshalJson()

	if marshalErr != nil {
		t.Errorf("Test_UnmarshalJson error %v", marshalErr)
	}

	r := make(StringFloateMap)
	r.UnmarshalJson(marshalJson)

	if reflect.DeepEqual(m, r) {
		t.Errorf("Test_UnmarshalJson input and result maps are not equal")
	}
}
