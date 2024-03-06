package utils_test

import (
	"reflect"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
)

func TestTagsUtils(t *testing.T) {
	t.Parallel()

	t.Run("splits comma separated tags", func(t *testing.T) {
		t.Parallel()
		tags := "json,yaml"
		expected := []string{"json", "yaml"}
		actual, _ := utils.FormatTags(tags)
		if !reflect.DeepEqual(actual, expected) {
			t.Error("expected", expected, "got", actual)
		}
	})

	t.Run("defaults to json if no tags are provided", func(t *testing.T) {
		t.Parallel()
		tags := ""
		expected := []string{"json"}
		actual, _ := utils.FormatTags(tags)
		if !reflect.DeepEqual(actual, expected) {
			t.Error("expected", expected, "got", actual)
		}
	})

	t.Run("returns an error if an invalid tag is provided", func(t *testing.T) {
		t.Parallel()
		tags := "json,xml"
		expected := "invalid tag: xml\n"
		_, err := utils.FormatTags(tags)
		if err.Error() != expected {
			t.Error("expected", expected, "got", err)
		}
	})
}
