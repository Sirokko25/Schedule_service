package helpers

import (
	"fmt"
	"strings"
)

func CreateResponceString(idSlice []int64) string {
	baseStr := "IDs of current schedules:"
	var sb strings.Builder
	for i, v := range idSlice {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	resultString := fmt.Sprintf("%s %s", baseStr, sb.String())
	return resultString
}
