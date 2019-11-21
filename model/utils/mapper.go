package commutils

import (
	"fmt"
	"github.com/petersunbag/coven"
	"log"
	"reflect"
	"sync"
)

var (
	lock sync.Mutex
	cMap = make(map[string]*coven.Converter)
)

// 映射(单);
func Map(src, dst interface{}) error {
	key := fmt.Sprintf("%v_%v", reflect.TypeOf(src).String(), reflect.TypeOf(dst).String())
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Map:%v:%v", key, err)
		}
	}()
	var err error
	c := cMap[key]
	if c == nil {
		lock.Lock()
		defer lock.Unlock()
		c, err = coven.NewConverter(dst, src)
		if err != nil {
			log.Printf("Map:%v:%v", key, err)
			return err
		}
		cMap[key] = c
	}
	err = c.Convert(dst, src)
	if err != nil {
		log.Printf("Map:%v:%v", key, err)
		return err
	}
	return nil
}

// 映射(多);
func Maps(src, dst interface{}) (err error) {
	return Map(src, dst)
}
