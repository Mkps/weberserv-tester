package main

import (
	"fmt"
	//    "net/http/httptest"
	"io"
	"net/http"
)

func testGet(address string, x int, s string) bool {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Println("error getting request")
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error opening body")
	}
	strBody := string(body)
	if resp.StatusCode != x {
        fmt.Println("Error codes do not match: expected ", x, " got ", resp.StatusCode)
		return false
	}
	if resp.StatusCode >= 400 {
		return true
	}
	if strBody != s {
        fmt.Println("Bodies do not match: expected ", s, " got ", strBody)
		return false
	}
	return true

}

func logTest(b bool) {
	var s string
	if b {
		s = "OK"
	} else {
		s = "KO"
	}
	fmt.Println("Test is ", s)
}
func main() {
	address := "http://127.0.0.1:8080"
	logTest(testGet(address+"/", 200, "42"))
	logTest(testGet(address+"/hello/hello.html", 200, "hello"))
	logTest(testGet(address+"/noget/", 405, ""))
	logTest(testGet(address+"/doesntexist/", 404, ""))
	logTest(testGet(address+"/notallowed/", 403, ""))
	logTest(testGet(address+"/hello/hello.py", 200, "hello\n"))
	logTest(testGet(address+"/hello/hello.py?user=test", 200, "hello test\n"))
}
