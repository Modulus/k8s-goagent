package envutils

import (
	"os"
	"testing"
)

func TestExtractPodFilter_LengthOfList_ShouldBeTwo(t *testing.T) {
	expected := "app=snadder,owner=nissen"
	os.Setenv("POD_FILTER", expected)
	result := ExtractPodFilter()
	if len(result) != 24 {
		t.Errorf("Length was incorrect, got: %v, want: %v.", result, 2)
	}

	if result != expected {
		t.Errorf("Result should be '%s', but was %s", expected, result)
	}
}
