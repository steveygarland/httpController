package main

import (
		"encoding/json"
		"fmt"
		"io"
		"io/ioutil"
		"net/http"

	)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome!")
}




func AppControl(w http.ResponseWriter, r *http.Request) {
        var  appToStart App
        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
                if err != nil   {
                                panic(err) 
                                }
                if err := r.Body.Close(); err != nil { 
                                panic(err) 
                                }
        if err := json.Unmarshal(body, &appToStart); err != nil { 
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(422) //unprocessable entity
                if err := json.NewEncoder(w).Encode(err); err != nil {
                        panic(err)
                }
        }

        a := RepoAppControl(appToStart)
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusCreated)
        if err:=json.NewEncoder(w).Encode(a); err != nil {
                panic(err)
        }
}



