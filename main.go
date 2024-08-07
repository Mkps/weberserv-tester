package main

import (
	"fmt"
	//    "net/http/httptest"
	"io"
	"net/http"
	"strings"
)

func testGet(address string, x int, s string) string {
	resp, err := http.Get(address)
	if err != nil {
		return "error getting request"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "error opening body"
	}
	strBody := string(body)
	if resp.StatusCode != x {
		return fmt.Sprint("Error codes do not match: expected ", x, " got ", resp.StatusCode)
	}
	if resp.StatusCode >= 400 {
		return ""
	}
	if strBody != s {
		return fmt.Sprint("Bodies do not match: expected ", s, " got ", strBody)
	}
	return ""

}

func testCgi(address string, x int, s string) string {
	resp, err := http.Get(address)
	if err != nil {
		return "error getting request"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "error opening body"
	}
	strBody := string(body)
	if resp.StatusCode != x {
		return fmt.Sprint("Error codes do not match: expected ", x, " got ", resp.StatusCode)
	}
	if resp.StatusCode >= 400 {
		return ""
	}
	if !strings.Contains(strBody, s) {
		return fmt.Sprint("Could not find substring s", s, "in response body")
	}
	return ""

}
func logTest(test string, s string) {
	red := "\033[31m"
	green := "\033[32m"
	reset := "\033[0m"
	var tmp string
	if s == "" {
		test := fmt.Sprintf("Testing %s", test)
		result := fmt.Sprint(green, "[OK]", reset)
		tmp = fmt.Sprintf("%-50s%s", test, result)
	} else {
		test := fmt.Sprintf("Testing %s", test)
		result := fmt.Sprint(red, "[KO]", reset)
		tmp = fmt.Sprintf("%-50s%s\t%s", test, result, s)
	}
	fmt.Println(tmp)

}
func main() {
	address := "http://127.0.0.1:8080"
	logTest("autoindex", testGet(address+"/", 200, "42"))
	logTest("direct access", testGet(address+"/hello/hello.html", 200, "hello"))
	logTest("method not allowed on path", testGet(address+"/noget/", 405, ""))
	logTest("access non-existing ressource", testGet(address+"/doesntexist/", 404, ""))
	logTest("forbidden directory", testGet(address+"/notallowed/", 403, ""))
	logTest("python script without argument", testGet(address+"/hello/hello.py", 200, "hello\n"))
	logTest("python script with argument", testGet(address+"/hello/hello.py?user=test", 200, "hello test\n"))
	logTest("php script without argument", testCgi(address+"/cgi/phptest.php", 200, "GET"))
	logTest("php script with argument", testCgi(address+"/cgi/phptest.php?user=test", 200, "test"))
	logTest("uri test invalid character %", testCgi(address+"/cgi/phptest.php?user=test%", 400, ""))
    bigstring := fmt.Sprintf("%s%s%s",address,"/",strings.Repeat("a",5000))
	logTest("uri too long", testGet(bigstring, 414, ""))
}
