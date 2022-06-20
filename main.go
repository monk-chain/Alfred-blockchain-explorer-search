package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
)

var (
    wf         *aw.Workflow
    cache_dir  = "./cache"
    icon_dir  = "./icon"
    cache_file = "cache.json"
)

type Chain struct {
    Name   string `json:"name"`
    Chain   string `json:"chain"`
    ChainId   int `json:"chainId"`
    ShortName   string `json:"shortName"`
    Icon   string `json:"icon"`
    Cid string
	Explorers []Explorers `json:"explorers"`
}

type Explorers struct {
    Url string `json:"url"`
}

type Icon struct {
    Url string `json:url`
	Width int  `json:width`
	Height int `json:height` 
	Format string `json:format`
}

func init() {
    wf = aw.New()
}

func run() {

    update := flag.Bool("update", false, "Update Channels")
	searchStr := flag.String("search", "", "Search String")
    flag.Parse()

    if *update {
        UpdateChains()
    } else{
        SearchChains(*searchStr)
    }

}

func main() {
    wf.Run(run)
}