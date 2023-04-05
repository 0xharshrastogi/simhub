package profitcalculator

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/harshrastogiexe/lib/go/common"
	"github.com/harshrastogiexe/lib/go/proxy"
)

type profitCalculator struct {
	realmType     common.RealmType
	realm         *common.Realm
	cachedPrice   map[string]float64
	cachePriceMux sync.Mutex
	Admin         float64
}

// New returns a new ProfitCalculator for the given realm
func New(r common.RealmType) *profitCalculator {
	var realm *common.Realm
	switch r {
	case common.Entrepreneurs:
		realm = &proxy.Entrepreneurs
	case common.Magnates:
		realm = &proxy.Magnates
	}
	return &profitCalculator{
		realmType:   r,
		realm:       realm,
		cachedPrice: map[string]float64{},
		Admin:       40,
	}
}

func (c *profitCalculator) Calculate(building string) (any, error) {
	t := time.Now()
	cl := log.New(os.Stdout, fmt.Sprintf("Calculate(id: %s) ", building), log.Default().Flags())
	cl.Println("fetching properties of building id:", building)
	b, err := c.realm.Encyclopedia.GetBuildingById(building)
	if err != nil {
		return nil, err
	}
	cl.Println("successfully fetched info for building", b.Name)
	cl.Println("fetching information for all resources of building produces")
	rCh := c.getAllBuildingResources(b)
	wg := sync.WaitGroup{}
	for r := range rCh {
		cl.Println("building", b.Name, "produces:", r.Name)
		cl.Println("fetching production prices for", strings.ToLower(r.Name))
		wg.Add(1)
		go func(r *common.Resource) {
			defer wg.Done()
			p, err := c.realm.GetProductionCost(r, c.Admin, 0)
			if err == nil {
				cl.Println("production price for", strings.ToLower(r.Name), "is:", p.TotalCost())
			}
		}(r)
	}
	cl.Println("waiting for all resources price to be calculated")
	wg.Wait()
	cl.Println("stopping")
	cl.Println("time elapsed:", time.Since(t).Seconds())
	return nil, nil
}

// getAllBuildingResources returns channels for resource
func (c *profitCalculator) getAllBuildingResources(building *common.Building) <-chan *common.Resource {
	l := log.New(os.Stdout, fmt.Sprintf("getAllBuildingResources(id: %s) ", building.Name), log.Default().Flags())
	ch := make(chan *common.Resource)
	wg := sync.WaitGroup{}
	for _, r := range building.DoesProduce {
		l.Println("fetching information for resource", r.Name)
		wg.Add(1)
		go func(rb common.ResourceBase) {
			defer wg.Done()
			rInfo, err := c.realm.Encyclopedia.GetResourceById(fmt.Sprint(rb.Id))
			if err != nil {
				return
			}
			l.Println("sending", rInfo.Name, "info to channel")
			ch <- rInfo
		}(r)
	}
	go func() {
		l.Println("waiting for all resources to be fetched")
		wg.Wait()
		l.Println("closing channel for resource")
		close(ch)
	}()
	return ch
}
