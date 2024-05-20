package tags_test

import (
	"reflect"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/pkg/tags"
)

func TestTagsUtils(t *testing.T) {
	t.Parallel()

	t.Run("splits comma separated inputTags", func(t *testing.T) {
		t.Parallel()
		inputTags := "json,yaml"
		expected := []string{"json", "yaml"}
		actual, _ := tags.FormatTags(inputTags)
		if !reflect.DeepEqual(actual, expected) {
			t.Error("expected", expected, "got", actual)
		}
	})

	t.Run("defaults to json if no inputTags are provided", func(t *testing.T) {
		t.Parallel()
		inputTags := ""
		expected := []string{"json"}
		actual, _ := tags.FormatTags(inputTags)
		if !reflect.DeepEqual(actual, expected) {
			t.Error("expected", expected, "got", actual)
		}
	})

	t.Run("returns an error if an invalid tag is provided", func(t *testing.T) {
		t.Parallel()
		inputTags := "json,xml"
		expected := "invalid tag: xml\n"
		_, err := tags.FormatTags(inputTags)
		if err.Error() != expected {
			t.Error("expected", expected, "got", err)
		}
	})
}
