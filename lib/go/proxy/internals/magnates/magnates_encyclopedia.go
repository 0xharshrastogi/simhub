package magnates

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/harshrastogiexe/lib/go/common"
)

const tOTAL_RESOURCES int = 140

type MagnatesEncyclopedia struct{}

func CreateResourcesIdIterator() func() (int, bool) {
	const MAX_RESOURCE_ID = 145
	current := 1
	return func() (int, bool) {
		switch current {
		case 36:
			current = 40
		case 144:
			current = 145
		}
		id, done := current, current == MAX_RESOURCE_ID
		current++
		return id, done
	}
}

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

func (e MagnatesEncyclopedia) GetResourcesList() ([]*common.Resource, error) {
	const RESOURCE_BUFFER_SIZE int = 32
	var (
		resource = make([]*common.Resource, tOTAL_RESOURCES)
		isErr    bool
	)
	rChan := make(chan struct {
		index    int
		resource *common.Resource
	}, RESOURCE_BUFFER_SIZE)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// routine to put resource at their respective index
	go func() {
		defer wg.Done()
		for v := range rChan {
			resource[v.index] = v.resource
		}
	}()
	iterator := CreateResourcesIdIterator()
	fetchWg := &sync.WaitGroup{}
	for i := 0; !isErr; i++ {
		fetchWg.Add(1)
		v, done := iterator()
		// routine to fetch resource
		go func(id string, index int) {
			defer fetchWg.Done()
			r, err := e.GetResourceById(id)
			if err != nil {
				isErr = true
				close(rChan)
				return
			}
			rChan <- struct {
				index    int
				resource *common.Resource
			}{
				index:    index,
				resource: r,
			}
		}(fmt.Sprint(v), i)
		if done {
			break
		}
	}
	// routine to close resource channel if all routines fetching resource are done
	go func() {
		if isErr {
			return
		}
		fetchWg.Wait()
		close(rChan)
	}()
	wg.Wait()
	return resource, nil
}
