package main

import (
	"fmt"
	"sync"

	"github.com/harshrastogiexe/lib/go/common"
)

func ProfitCalculator(realm *common.Realm, building string) (any, error) {
	b, err := realm.Encyclopedia.GetBuildingById(building)
	if err != nil {
		return nil, err
	}
	ch := make(chan *common.Resource)
	wg := &sync.WaitGroup{}
	for _, item := range b.DoesProduce {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			r, err := realm.Encyclopedia.GetResourceById(fmt.Sprintf("%d", id))
			if err != nil {
				return
			}
			ch <- r

		}(item.DBLetter)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for r := range ch {
		fmt.Println(r)
	}
	return nil, nil
}
