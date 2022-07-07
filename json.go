package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type StringFloateMap map[string]float64

func (m *StringFloateMap) MarshalJson() ([]byte, error) {
	ss := map[string]string{}

	for key, val := range *m {
		str := fmt.Sprintf("%f", val)
		ss[key] = str
	}

	return json.Marshal(ss)
}

func (m *StringFloateMap) UnmarshalJson(data []byte) error {
	ss := map[string]string{}
	err := json.Unmarshal(data, &ss)

	if err != nil {
		return err
	}

	for key, val := range ss {
		result, resErr := strconv.ParseFloat(val, 64)

		if resErr != nil {
			return err
		}

		(*m)[key] = result
	}

	return nil
}
