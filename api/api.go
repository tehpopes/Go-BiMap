package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bimap"
)

type BiMapRequest struct {
	BiMapPairs []BiMapPair  `json:"bimap"`
	Type string   `json:"type"`
	NewKey *int    `json:"key"`
	NewValue *int  `json:"value"`
}

type BiMapPair struct {
	Key int    `json:"k"`
	Value int  `json:"v"`
} 

func logOnError(err error, msg string) {
	if err != nil{
		log.Println(msg,": ", err)
	}
}

func main() {
	http.HandleFunc("/", handle_bimap_function_request)
	fmt.Println("Waiting for json POST request on port 3000")
	log.Fatal(http.ListenAndServe(":3000",nil))
}

func handle_bimap_function_request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { 
		var request BiMapRequest
		data, err := ioutil.ReadAll(r.Body) 
		logOnError(err, "Failed reading message body")
		
		err = json.Unmarshal(data, &request)
		logOnError(err, "Failed unmarshal")
		
		bm := createBiMap(&request.BiMapPairs)
		if bm != nil {
			if performUserTask(bm,&request) == 0 {
				fmt.Fprintf(w, convertBiMapToJSON(bm))
			} 
		} else {
			log.Println("Invalid BiMap structure. BiMap must contain unique keys and values.") 
		}
		
	} // Ignore all other request types
}

func createBiMap(bm_pairs *[]BiMapPair) *bimap.BiMap {
	bm := bimap.NewBiMap()
	var keys, vals map[int]bool
	for _, bm_pair := range *bm_pairs {
		if _, ok := keys[bm_pair.Key]; ok {
			return nil
		}
		if _, ok := vals[bm_pair.Value]; ok {
			return nil
		}
		bm.Insert(bm_pair.Key, bm_pair.Value)
	}
	return bm
}

func performUserTask(bm *bimap.BiMap, bmr *BiMapRequest) int {
	switch t := bmr.Type; t {
	case "Insert":
		return handleInsert(bm,bmr)
	case "RemoveByKey":
		return handleRemoveByKey(bm,bmr)
	case "RemoveByValue":
		return handleRemoveByValue(bm,bmr)
	default:
		log.Println("Invalid Request Type. Type can only be one of the following: \"Insert\", \"FindByKey\", or \"FindByValue\"")
		return -1
	}
}

func handleInsert(bm *bimap.BiMap, bmr *BiMapRequest) int {
	if bmr.NewKey == nil || bmr.NewValue == nil {
		log.Println("Error. User request does not contain both key and value.")
		return -1
	}
	bm.Insert(*bmr.NewKey, *bmr.NewValue)
	return 0
}

func handleRemoveByKey(bm *bimap.BiMap, bmr *BiMapRequest) int {
	if bmr.NewKey == nil {
		log.Println("Error. User request does not contain key.")
		return -1
	}
	bm.RemoveByKey(*bmr.NewKey)
	return 0
}

func handleRemoveByValue(bm *bimap.BiMap, bmr *BiMapRequest) int {
	if bmr.NewValue == nil {
		log.Println("Error. User request does not contain value.")
		return -1
	}
	bm.RemoveByValue(*bmr.NewValue)
	return 0
}

func convertBiMapToJSON(bm *bimap.BiMap) string {
	newPairs := []BiMapPair{}
	for k,v := range bm.Map() {
		var bmp BiMapPair
		bmp.Key = k.(int)
		bmp.Value = v.(int)

		newPairs = append(newPairs, bmp)
	}
	jsonString, err := json.Marshal(newPairs)
	logOnError(err, "Marshal to JSON failed")
	return string(jsonString)
}