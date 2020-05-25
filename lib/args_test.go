package lib

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func checker2(t *testing.T, want1 interface{}, want2 interface{}) func(got1 interface{}, got2 interface{}) {
	return func(got1 interface{}, got2 interface{}) {
		if want1 != got1 || want2 != got2 {
			t.Errorf("got: %#v %#v, want %#v %#v", got1, got2, want1, want2)
		}
	}
}

func TestStripSigil(t *testing.T) {
	checker2(t, "%%", "test")(StripSigil("%%test"))
	checker2(t, "%", "test")(StripSigil("%test"))
	checker2(t, "/", "test")(StripSigil("/test"))
	checker2(t, "%%", "")(StripSigil("%%"))
	checker2(t, "", "test")(StripSigil("test"))
	checker2(t, "", "")(StripSigil(""))
}

func TestParseCommonArgs(t *testing.T) {
	proj := "proj"
	col := "col"
	query := "query"
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want CommonArgs
	}{
		{name: "", args: args{[]string{"%col", "/query", "rest"}}, want: CommonArgs{
			Column: &col,
			Search: &query,
			Args:   []string{"rest"},
		}},
		{name: "", args: args{[]string{"%%proj", "/query", "--", "rest", "%col"}}, want: CommonArgs{
			Project: &proj,
			Search:  &query,
			Args:    []string{"rest", "%col"},
		}},
		{name: "", args: args{[]string{"%%proj", "/query", "rest", "%col"}}, want: CommonArgs{
			Project: &proj,
			Column:  &col,
			Search:  &query,
			Args:    []string{"rest"},
		}},
		{name: "", args: args{[]string{"/query", "rest", "--"}}, want: CommonArgs{
			Search: &query,
			Args:   []string{"rest"},
		}},
		{name: "", args: args{[]string{"^%col", "/query", "rest", "--"}}, want: CommonArgs{
			FromColumn: &col,
			Search:     &query,
			Args:       []string{"rest"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseCommonArgs(tt.args.args)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}
