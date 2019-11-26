package main

import "fmt"
import "sync"
import "time"
import "math/rand"

// import "net/http"
import "github.com/go-resty/resty/v2"

var url string = "https://api.cloudflareclient.com/v0a745/reg"
var myIDPc string = "84a66e5f-7136-4211-bef0-c8ef711d27ef"
var myIDAndroid string = "2e4f5722-0362-450f-ad3b-30eefa35c249"

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		// if i % 10 == 0 {
		//   time.Sleep(1 * time.Second)
		// }
		time.Sleep(500 * time.Millisecond)
		go func(i int) {
			fmt.Println(i)
			callAPI(myIDPc)
			callAPI(myIDAndroid)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func genString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func callAPI(id string) {
	installID := genString(10)
	body := make(map[string]interface{})
	body["key"] = genString(41) + "="
	body["install_id"] = installID
	body["fcm_token"] = installID + ":APA91b" + genString(133)
	body["referrer"] = id
	// body["warp_enabled"] = false
	// body["tos"] = time.Now().Format("2006-01-02T15:04:05.999-07:00")
	// body["type"] = "Android"
	// body["locale"] = "zh-CN"

	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=UTF-8"
	header["Host"] = "api.cloudflareclient.com"
	header["Connection"] = "Keep-Alive"
	header["Accept-Encoding"] = "gzip"
	header["User-Agent"] = "okhttp/3.12.1"

	res, err := resty.New().NewRequest().SetBody(body).SetHeaders(header).Post(url)
	if err != nil {
		println(err)
	}
	if res.IsError() {
		fmt.Printf("%+v", res)
	}
	// fmt.Printf("%+v \n", res.Request.Body)
	// fmt.Println("=====================================")
	// fmt.Printf("%+v \n", res.Request.Header)
	// fmt.Println("=====================================")
	// fmt.Printf("%+v \n", res)
}
