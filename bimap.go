/*
* BiMap (Bidirectional Map) Functionality:
*  - maintains unique keys and values
*  - if same key value pair inserted, returns nil
*  - if same key but different value inserted,
*		new value replaces old one
*  - findByKey returns value or nil
*  - findByValue returns value or nil
*  - removeByKey and removeByValue are self explanatory,
*       same return values as their find counterparts
*/

package bimap

import (
	"fmt"
	//"log"
	//"encoding/json"
)

type BiMap struct {
	frontMap map[interface{}]interface{}
	backMap map[interface{}]interface{}
}

func NewBiMap() *BiMap {
	var bm BiMap
	bm.frontMap = make(map[interface{}]interface{})
	bm.backMap = make(map[interface{}]interface{})
	return &bm
}

func (bm *BiMap) FindByKey(key interface{}) interface{} {
	if val, ok := bm.frontMap[key]; ok {
		return val
	}
	return nil
}

func (bm *BiMap) FindByValue(value interface{}) interface{} {
	if key, ok := bm.backMap[value]; ok {
		return key
	}
	return nil
}

func (bm *BiMap) Insert(key interface{}, value interface{}) interface{} {
	if val, ok := bm.frontMap[key]; ok {
		if val != value {
			delete(bm.backMap,val)
		} else {
			return nil 
		}
	}
	if k, ok := bm.backMap[value]; ok {
		if k != key {
			delete(bm.frontMap, k)
		} else {
			return nil
		}
	}

	bm.frontMap[key] = value
	bm.backMap[value] = key
	return value
} 

func (bm *BiMap) RemoveByKey(key interface{}) interface{} {
	val := bm.FindByKey(key)
	if val != nil {
		delete(bm.frontMap,key)
		delete(bm.backMap,val)	
	}
	return val
}

func (bm *BiMap) RemoveByValue(value interface{}) interface{} {
	key := bm.FindByValue(value)
	if key != nil {
		delete(bm.frontMap,key)
		delete(bm.backMap,value)	
	}
	return key
}

func (bm *BiMap) Map() map[interface{}]interface{} {
	newMap := make(map[interface{}]interface{})
	for k,v := range bm.frontMap {
		newMap[k] = v
	}
	return newMap	
}

func (bm *BiMap) Keys() []interface{} {
	keys := make([]interface{}, 0, len(bm.frontMap))
	for k := range bm.frontMap {
		keys = append(keys,k)
	}
	return keys
}

func (bm *BiMap) Values() []interface{} {
	vals := make([]interface{}, 0, len(bm.backMap))
	for v := range bm.backMap {
		vals = append(vals,v)
	}
	return vals
}

func (bm *BiMap) Print() {
	for k, v := range bm.frontMap { 
		fmt.Println("Key: ", k, " Value: ", v)
	}
}
/*
func (bm *BiMap) JSON() string {
	jsonString, err := json.Marshal(bm.frontMap)
	if err != nil {
		log.Println(err)
	}
	return string(jsonString)
}*/