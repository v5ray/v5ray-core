package strmatcher_test

import (
	"reflect"
	"testing"

	"github.com/v4fly/v4ray-core/v0/common"
	. "github.com/v4fly/v4ray-core/v0/common/strmatcher"
)

// See https://github.com/v2fly/v2ray-core/issues/92#issuecomment-673238489
func TestLinearIndexMatcher(t *testing.T) {
	rules := []struct {
		Type   Type
		Domain string
	}{
		{
			Type:   Regex,
			Domain: "apis\\.us$",
		},
		{
			Type:   Substr,
			Domain: "apis",
		},
		{
			Type:   Domain,
			Domain: "googleapis.com",
		},
		{
			Type:   Domain,
			Domain: "com",
		},
		{
			Type:   Full,
			Domain: "www.baidu.com",
		},
		{
			Type:   Substr,
			Domain: "apis",
		},
		{
			Type:   Domain,
			Domain: "googleapis.com",
		},
		{
			Type:   Full,
			Domain: "fonts.googleapis.com",
		},
		{
			Type:   Full,
			Domain: "www.baidu.com",
		},
		{
			Type:   Domain,
			Domain: "example.com",
		},
	}
	cases := []struct {
		Input  string
		Output []uint32
	}{
		{
			Input:  "www.baidu.com",
			Output: []uint32{5, 9, 4},
		},
		{
			Input:  "fonts.googleapis.com",
			Output: []uint32{8, 3, 7, 4, 2, 6},
		},
		{
			Input:  "example.googleapis.com",
			Output: []uint32{3, 7, 4, 2, 6},
		},
		{
			Input:  "testapis.us",
			Output: []uint32{2, 6, 1},
		},
		{
			Input:  "example.com",
			Output: []uint32{10, 4},
		},
	}
	matcherGroup := NewLinearIndexMatcher()
	for _, rule := range rules {
		matcher, err := rule.Type.New(rule.Domain)
		common.Must(err)
		matcherGroup.Add(matcher)
	}
	matcherGroup.Build()
	for _, test := range cases {
		if m := matcherGroup.Match(test.Input); !reflect.DeepEqual(m, test.Output) {
			t.Error("unexpected output: ", m, " for test case ", test)
		}
	}
}
