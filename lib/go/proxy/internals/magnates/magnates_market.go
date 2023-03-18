package magnates

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harshrastogiexe/lib/go/common"
)

type MagnatesMarket struct{}

// GetMarketInfo gets information for a Magnates Market
func (MagnatesMarket) GetMarketInfo(id string) (common.MarketInfo, error) {
	res, err := http.Get(fmt.Sprintf("https://www.simcompanies.com/api/v3/market/0/%s", id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	m := common.MarketInfo{}
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}
