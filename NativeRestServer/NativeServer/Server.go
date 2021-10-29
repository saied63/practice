package NativeServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Form struct {
	Age       int
	Name      string
	City      string
	Country   string
	CONTINENT string
}

type rServer struct {
}

func (rs *rServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//fmt.Println("hhhhh")
}

type User func(res http.ResponseWriter, req *http.Request)

func (fn User) ServerHTTP(res http.ResponseWriter, req *http.Request) {
	fn(res, req)
}

func HandleInsideThisFunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// get url path
	path := req.URL.Path
	fmt.Println("url path is : ", path)

	// by using req.URL.Query() we can access url.values (url parameters)
	keys := req.URL.Query()
	// Get("") will return only first found key with given name
	bb1 := keys.Get("bb")
	fmt.Println("all bb in url : " + bb1)
	bb2 := keys.Get("bb")
	fmt.Println("all bb in url : " + bb2)
	// process all url parameters this way we can access parameters with same names
	if len(keys) > 0 {
		for _, value := range keys {
			fmt.Println("url param : ", value)
		}
	}
	// we should first call (ParseMultipartForm) for get form-data from request
	// this parse mode will work on Content-Type = multipart/form-data and dos'nt work on application/x-www-form-urlencoded
	parseFormError := req.ParseMultipartForm(0)
	if parseFormError == nil {
		fmt.Println("parseFormError : ", parseFormError)
		//fmt.Println("Age ", req.FormValue("Age"))
		for key, sa := range req.PostForm {
			fmt.Println(key)
			fmt.Println(sa)
		}
	}

	f := Form{Age: 2}

	node := json.NewDecoder(req.Body)

	err := node.Decode(&f)
	fmt.Println("parsed json : ", f)
	if f.Country == "" {
		err = errors.New(`{"error":"country is needed"}`)
		http.Error(res, err.Error(), 400)
		return
	}
	if err != nil {
		reposnseError, er := json.Marshal(err)
		if er != nil {
			res.WriteHeader(400)
			res.Write([]byte(`{"error" : "bad request"}`))
			return
		}
		fmt.Println(err)
		res.WriteHeader(400)
		res.Write(reposnseError)
		return
	}

	res.WriteHeader(200)
	res.Write([]byte(`{"msg":"simpleServer"}`))
}

func StartServerRest() {
	//rs := &rServer{}
	//http.Handle("/", rs)
	http.HandleFunc("/api/v1/user/saied", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("call inside User function Handler", req.Body)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write([]byte(`{"msg":"simpleServer"}`))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
