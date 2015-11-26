package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// defines all the REST APIs handlers

func PutKey_3000(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key_id"]
	key_int,_ := strconv.Atoi(key)
	
	value:=vars["value"]
	map_3000[key_int] = value
	fmt.Println("putting the value in map");
	w.WriteHeader(200);
}

//Get request handling - to get one key
func GetKey_3000(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key_id"]
	key_int,_ := strconv.Atoi(key)
	fmt.Println(key_int);
	
	var Error string
	
	value, ok := map_3000[key_int]
	
	var res KeyValuePair
	if ok {
		fmt.Println("Key:",key_int,"Value found in map: ", value)
		res.Key = key_int
		res.Value = value
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
		
	} else {
		fmt.Println("key not found")
		Error = "key not found"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Error)
		
		}
}

//Get request handling to get all the keys
func GetAllKeys_3000(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the keys stored in the map");
	len:= len(map_3000)

	var Error string

	if len == 0{
		fmt.Println("Map is empty");
		Error = "Map is empty"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Error)
		return
	}

	i := 0
	// set the size
	var AllKeys = make([]KeyValuePair, len)
	for key,val := range map_3000 {
        AllKeys[i].Key = key
        AllKeys[i].Value = val
        fmt.Println("Key:",AllKeys[i].Key,"Value:",AllKeys[i].Value);
        i++
    }
    w.WriteHeader(200)
	json.NewEncoder(w).Encode(AllKeys)
	return
}
