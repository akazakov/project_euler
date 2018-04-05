package myutil

import (
  "io/ioutil"
  "strings"
)

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

func ReadCSVFile(file string) ([]string, error) {
  if strs, err := ReadFile(file); err == nil {
    result := make([]string, 0)
    for _, s := range strs {
      result = append(result, strings.Split(s, ",")...)
    }
    return result, nil
  } else {
    return nil, err
  }
}
