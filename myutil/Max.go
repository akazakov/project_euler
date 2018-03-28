package myutil

import (
  "io/ioutil"
  "strings"
)

func Max(a,b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func ReadFile(file string) ([]string, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	raw_strings := strings.Split(strings.TrimSpace(string(bs)), "\n")
	trimmed_strings := make([]string, len(raw_strings))
	for i, s := range raw_strings {
		trimmed_strings[i] = strings.TrimSpace(s)
	}
	return trimmed_strings, nil
}
