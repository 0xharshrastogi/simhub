package magnates

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harshrastogiexe/lib/go/common"
)

type MagnatesEncyclopedia struct{}

// GetResourceById fetches resource information from api
func (MagnatesEncyclopedia) GetResourceById(id string) (*common.Resource, error) {
	u := fmt.Sprintf("https://www.simcompanies.com/api/v4/en/0/encyclopedia/resources/1/%s", id)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	r := &common.Resource{}
	if err := d.Decode(&r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetBuildingById fetches building information from api
func (MagnatesEncyclopedia) GetBuildingById(id string) (*common.Building, error) {
	u := fmt.Sprintf("https://www.simcompanies.com/api/v3/0/encyclopedia/buildings/%s", id)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	d, r := json.NewDecoder(res.Body), &common.Building{}
	if err := d.Decode(r); err != nil {
		return nil, err
	}
	return r, nil
}
