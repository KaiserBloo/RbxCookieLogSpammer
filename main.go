// I'm currently learning golang so this may be messy.
// UPDATE: yeah its fuckin messy

// Used to spam roblox cookie loggers with a dummy cookie.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/zenthangplus/goccm"
)

func main() {
	// Read proxy file
	data, err := ioutil.ReadFile("proxies.txt")
	if err != nil {
		fmt.Println("Error reading proxy file.")
		return
	}

	proxies := strings.Split(string(data), "\n")
	threads := 20
	c := goccm.New(threads)

	cookie := "_%7CWARNING:-DO-NOT-SHARE-THIS.--Sharing-this-will-allow-someone-to-log-in-as-you-and-to-steal-your-ROBUX-and-items.%7C_FDEDB31DCAF973782EB5DA575C8BC20AAC50AE159E8663A8E71D53ACC356D0BFEDE08BC4BBAA33948AE8EF3C896408646E5714BCBEF3DD9D727DA0F8C97642B49A3D24A33E547584A7FD8E06CB6FF6A03C37EDD409634A059B293C7A2349FB5A5E798CB9BE766698A1D05FC9A76DEF3D9EC62DCF5A17D4940A20C6E48B33EB25F5CB020991CFB692BE91DBD46ED534F6A7814CDC9F72E9857B7F39F23E480286EEEA3DA444557A29500BB2C523F27A3957125925E791A57ABB9021291C1B99E63E4CD659D66D2C9004470B4175D2524A76F965EEB7A4B5380BCAE70693BBDD12DD59E633E6C788B7D80B9AC8FA62B8E605CC7BAAF43A0021861C0E02994D85A19470491C09C5DC0B22B1F5EBCEE6228075C72DD17578E811947E5FD45308AA7F05053973875E05E7ED61FBFB08417B41555BC2790E5D78F7EE013A3372930E13DCE8FEA5"
	url := "https://assetcopy.xyz/AssetGen/send.php?t="
	method := "GET"

	c.Wait()

	for i := 0; i < len(proxies); i++ {
		go dothing(cookie, url, method, proxies[i], i)
	}

	c.WaitAllDone()
}

func dothing(cookie string, addy string, method string, proxy string, i int) {

	// Create a new client
	proxyUrl, err := url.Parse("http://" + proxy)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	req, err := http.NewRequest(method, addy+cookie, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Inf request loop
	for {
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		status := res.StatusCode
		fmt.Println("[" + strconv.Itoa(i) + "] " + proxy + " | " + strconv.Itoa(status))
	}
}
