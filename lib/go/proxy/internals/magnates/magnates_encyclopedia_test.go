package magnates

import (
	"fmt"
	"testing"
)

func TestMagnatesEncyclopedia_GetResourceById(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{name: "returns power", args: "1", want: "1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := MagnatesEncyclopedia{}
			r, err := e.GetResourceById(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("MagnatesEncyclopedia.GetResourceById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if fmt.Sprint(r.Id) != tt.want {
				t.Errorf("MagnatesEncyclopedia.GetResourceById() = %v, want %v", r, tt.want)
			}
		})
	}
}
