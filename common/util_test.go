package common

import (
	"fmt"
	"testing"
)

func TestConvertToCaseAlpha(t *testing.T) {
	fmt.Println(ConvertToCaseAlpha("doc_name_test"))
}

func TestConvertHumpNameToSnakeCase(t *testing.T) {
	fmt.Println(ConvertHumpNameToSnakeCase("docNameTest"))
}
