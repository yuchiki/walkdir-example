package main

import (
	"testing"
)


func TestConvertRoot(t *testing.T) {
	type args struct {
		path     string
		srcRoot  string
		dstRoot  string
	}
	tests := []struct {
		name string
		args args
		expected string
	}{
		{
			name: "convert root of path from srcRoot to dstRoot",
			args: args{
				path: "a/b/c",
				srcRoot: "a",
				dstRoot: "x",
			},
			expected: "x/b/c",
		},
		{
			name: "does not change the path if the path has no srcRoot",
			args: args{
				path: "a/b/c",
				srcRoot: "x",
				dstRoot: "y",
			},
			expected: "a/b/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := convertRoot(tt.args.path, tt.args.srcRoot, tt.args.dstRoot)
			if actual != tt.expected {
				t.Errorf("actual: %s, expected: %s", actual, tt.expected)
			}
		})
	}
}

func TestConvertJSONToYAML(t *testing.T) {
	type args struct {
		jsonBytes []byte
	}
	tests := []struct {
		name string
		args args
		expected string
	}{
		{
			name: "converts a json object to a yaml object",
			args: args{
				jsonBytes: []byte(`{"a": 1, "b": 2}`),
			},
			expected: "a: 1\nb: 2\n",
		},
		{
			name: "converts a json array to a yaml array",
			args: args{
				jsonBytes: []byte(`[1, 2, 3]`),
			},
			expected: "- 1\n- 2\n- 3\n",
		},
		{
			name: "converts a json string to a yaml string",
			args: args{
				jsonBytes: []byte(`"hello"`),
			},
			expected: "hello\n",
		},
		{
			name: "converts a json number to a yaml number",
			args: args{
				jsonBytes: []byte(`3.14`),
			},
			expected: "3.14\n",
		},
		{
			name: "converts a json boolean to a yaml boolean",
			args: args{
				jsonBytes: []byte(`true`),
			},
			expected: "true\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := convertJSONToYAML(tt.args.jsonBytes)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if string(actual) != tt.expected {
				t.Errorf("actual: %s, expected: %s", actual, tt.expected)
			}
		})
	}
}

func TestConvertAllJsonFilesToYaml(t *testing.T) {
	// TODO: implement this test
}
