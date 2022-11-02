package common

import "fmt"

func MapToString(m map[string]string) (res string) {

	for k, v := range m {
		res += fmt.Sprintf("%s=%s \n", k, v)
	}
	return
}
