package init

import (
	"oauth_lambda/model"
	"reflect"
	"testing"
)

func TestConf(t *testing.T) {
	tests := []struct {
		name string
		want model.Config
	}{
		{
			"1",
			model.Config{HmacSecret: "1885df74d00dbbe19274c6d955feeb5b"},
		},
		{
			"2",
			model.Config{HmacSecret: "1885df74d00dbbe19274c6d955feeb5b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Conf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conf() = %v, want %v", got, tt.want)
			}
		})
	}
}
