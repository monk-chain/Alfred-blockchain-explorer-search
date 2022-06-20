package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	aw "github.com/deanishe/awgo"
)

func UpdateChains() {


	wf.NewItem("Update Channels").Valid(true)

	c := aw.NewCache(cache_dir)
	resp, err := http.Get("https://chainid.network/chains.json")
	if err != nil {
		wf.Warn("Error", "Error chains Api")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wf.Warn("Error", "Error Body pase")
	}

	var data []Chain

	if err := json.Unmarshal(body, &data); err != nil {
		wf.Warn("Error", "Error json Unmarshal")
	}

	var icons []Icon
	os.Mkdir(cache_dir + "/" + icon_dir, 0755)


	var wg sync.WaitGroup
	wg.Add(1)

	go func (){
		for key, chain := range data {


			if len(chain.Icon) == 0  {
				continue
			}

			resp, err := http.Get("https://raw.githubusercontent.com/ethereum-lists/chains/master/_data/icons/" + chain.Icon + ".json")
			if err != nil {
				wf.FatalError(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				wf.Warn("Error", "Error Body pase")
			}
			if err := json.Unmarshal(body, &icons); err != nil {
				wf.Warn("Error", "Error json Unmarshal")
			}

			cid := strings.Replace(icons[0].Url, "ipfs://", "", 1)
			data[key].Cid = cid

			resp, err = http.Get("https://github.com/ethereum-lists/chains/raw/master/_data/iconsDownload/"+ cid)
			if err != nil {
				fmt.Println(err)
			}

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			img, _, err := image.Decode(bytes.NewReader(body))
			if err != nil {
				continue
			}

			file, err := os.Create(cache_dir + "/" + icon_dir + "/" + chain.Chain+".png")
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()


			err = png.Encode(file, img)
			if err != nil {
				continue
			}
			fmt.Printf("Type: %[1]T , Value: %[1]v\n", cid)
		}
		wg.Done()
	}()

	wg.Wait()

	c.StoreJSON(cache_file, data)
	wf.SendFeedback()
}