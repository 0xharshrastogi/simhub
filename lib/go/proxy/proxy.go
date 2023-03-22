package proxy

import (
	"encoding/json"
	"net/http"

	"github.com/harshrastogiexe/lib/go/common"
	"github.com/harshrastogiexe/lib/go/proxy/internals/magnates"
)

var (
	Magnates common.Realm = common.Realm{
		Encyclopedia: magnates.MagnatesEncyclopedia{},
		Market:       magnates.MagnatesMarket{},
	}
	Entrepreneurs common.Realm = common.Realm{}
)

// GetRealmInfo returns information all realms in the game
func GetRealmInfo() (map[common.RealmType]*common.RealmInfo, error) {
	res, err := http.Get("https://www.simcompanies.com/api/realms/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	realm := map[common.RealmType]*common.RealmInfo{}
	if err := json.NewDecoder(res.Body).Decode(&realm); err != nil {
		return nil, err
	}
	return realm, nil
}
