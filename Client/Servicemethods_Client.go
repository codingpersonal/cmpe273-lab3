package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"io/ioutil"
)


//response msg struct
type KeyValuePair struct {
	Key int `json:"key"`
	Value string `json:"value"`
}

var cons ConsistentHashingMgr

// defines all the REST APIs handlers

func PutKeyToServer(key string, value string, server string) {
	client := &http.Client{}
	reqURL := "http://" + server + "/keys/" + key + "/" + value
	fmt.Println("URL formed:" + reqURL)
	req, err := http.NewRequest("PUT", reqURL, strings.NewReader(""))
	_, err = client.Do(req)

	if err != nil {
		fmt.Println("error in sending put req to Server ", err)
	}
}

func GetKeyFromServer(key string, server string) KeyValuePair {
	client := &http.Client{}
	reqURL := "http://" + server + "/keys/" + key
	fmt.Println("GET URL formed:" + reqURL)

	req, err := http.NewRequest("GET", reqURL, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error in sending req to google: ", err);	
	}
	
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		fmt.Println("error in sending GET req to Server ", err)
		return KeyValuePair{}
	}
	
	var resB KeyValuePair
	err = json.Unmarshal(body, &resB)
	fmt.Println("got the body: ", resB, err)
	return resB
}

func PutKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key_id"]
	value:=vars["value"]
	server, _ := cons.GetShardForKey(key)
	PutKeyToServer(key, value, server)
//	fmt.Println("putting the value in map");
	w.WriteHeader(200);
	json.NewEncoder(w).Encode("")
}

//Get request handling - to get one key
func GetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key_id"]
	server, err := cons.GetShardForKey(key)
	if err != "" {
		fmt.Println("some error in getting shard");
	}
	
	value := GetKeyFromServer(key, server)
	fmt.Println("getting the value in map for server: ", server);
	w.WriteHeader(200);
	json.NewEncoder(w).Encode(value)
}

