package magnates

import (
	"fmt"
	"testing"

	"github.com/harshrastogiexe/lib/go/common"
)

func TestMagnatesMarket_GetMarketInfo(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		m       MagnatesMarket
		args    args
		want    common.MarketInfo
		wantErr bool
	}{
		{name: "fetches the market price of power", args: args{id: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MagnatesMarket{}
			got, err := m.GetMarketInfo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MagnatesMarket.GetMarketInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId := fmt.Sprint(got[0].ResourceId); tt.args.id != gotId {
				t.Errorf("MagnatesMarket.GetMarketInfo() = %s, want %v", gotId, tt.args.id)
			}
		})
	}
}
