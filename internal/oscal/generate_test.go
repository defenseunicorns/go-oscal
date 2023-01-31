package oscal

import (
	"testing"
)

// TestFmtFieldName tests that we handle formatting a json string to a go struct correctly
func TestFmtFieldName(t *testing.T) {
	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		{in: "foo_id", out: "FooID"},
		{in: "fooId", out: "FooID"},
		{in: "foo_url", out: "FooURL"},
		{in: "foobar", out: "Foobar"},
		{in: "url_sample", out: "URLSample"},
		{in: "_id", out: "ID"},
		{in: "__id", out: "ID"},
	}

	for _, testCase := range testCases {
		actual := FmtFieldName(testCase.in)
		expected := testCase.out
		if expected != actual {
			t.Errorf("error FmtFieldName(): expected: %s | got: %s", expected, actual)
		}
	}
}
