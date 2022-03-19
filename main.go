package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	for p := 1; p < 10; p++ {
		u := fmt.Sprintf("https://api.github.com/repos/CosmWasm/cw-plus/releases?page=%d", p)
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Accept", "application/vnd.github.v3+json")
		client := &http.Client{}
		resp, err := client.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		rm := make([]map[string]interface{}, 0, 30)
		err = json.Unmarshal(body, &rm)
		if err != nil {
			panic(err)
		}
		for ri := range rm {
			fmt.Println("====================================================")
			v := rm[ri]["name"].(string)
			fmt.Println("v =", v)
			fc := false
			assets := rm[ri]["assets"].([]interface{})
			for ai := range assets {
				asset := assets[ai].(map[string]interface{})
				if asset["name"].(string) == "cw3_fixed_multisig.wasm" {
					fc = true
					fmt.Println("file size = ", asset["size"])
					break
				}
			}
			if !fc {
				fmt.Println("NOT FOUND CONTRACT", v)
				continue
			}
			csu := fmt.Sprintf("https://github.com/CosmWasm/cw-plus/releases/download/%s/checksums.txt", v)
			resp, err := http.Get(csu)
			if err != nil {
				panic(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			s := strings.Split(string(body), "\n")
			for si := range s {
				sl := strings.Split(s[si], " ")
				if len(sl) < 3 {
					continue
				}
				if sl[2] == "cw3_fixed_multisig.wasm" {
					fmt.Println("hash =", sl[0])
				}
			}

			csu = fmt.Sprintf("https://github.com/CosmWasm/cw-plus/releases/download/%s/checksums_intermediate.txt", v)
			resp, err = http.Get(csu)
			if err != nil {
				panic(err)
			}
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			s = strings.Split(string(body), "\n")
			for si := range s {
				sl := strings.Split(s[si], " ")
				if len(sl) < 3 {
					continue
				}
				if sl[2] == "target/wasm32-unknown-unknown/release/cw3_fixed_multisig.wasm" {
					fmt.Println("hash intermediate =", sl[0])
				}
			}
			fmt.Println("filename = \"cw3_fixed_multisig.wasm\"")
		}
	}
}
