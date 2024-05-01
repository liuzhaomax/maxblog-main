package business

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIsValidTimeFormat(t *testing.T) {
	cases := []struct {
		name string
		have string
		want bool
	}{
		{
			name: "成功-正确时间格式",
			have: "2024-04-17 12:28:48.026",
			want: true,
		},
		{
			name: "成功-没有横杠",
			have: "20240417 12:28:48.026",
			want: false,
		},
	}
	for _, tc := range cases {
		got := IsValidTimeFormat(tc.have)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, got, tc.want, fmt.Sprintf("\n*** Expected: \n %#v \n*** Got: \n %#v", tc.want, got))
		})
	}
}
