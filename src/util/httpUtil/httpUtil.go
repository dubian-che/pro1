package httpUtil

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpUtil_Req(url string) (ans string, code int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	code = resp.StatusCode
	ans = string(body)
	if resp.StatusCode != 200 {

		fmt.Println("http fail: ", url)
	}
	return
}
func Test() {
	HttpUtil_Req("http://my-json-server.typicode.com/dubian-che/pro1/posts/1")
}
