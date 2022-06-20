package main

import (
	"strings"

	aw "github.com/deanishe/awgo"
)

func SearchChains(searchStr string) {

    c := aw.NewCache(cache_dir)
    var chains []Chain

    if c.Exists(cache_file) {
        if err := c.LoadJSON(cache_file, &chains); err != nil {
            wf.FatalError(err)
        }
        for _, chain := range chains {
            if chain.Explorers == nil || len(chain.Explorers) == 0   {
                continue
            }
            icon := aw.Icon{Value:"./cache/icon/"+ chain.Chain + ".png", }
            if len(searchStr) == 0 {
                wf.NewItem(chain.Name).Subtitle(chain.Chain).Autocomplete(chain.Chain).Arg(chain.Explorers[0].Url).Icon(&icon).Valid(true)
            }else if (strings.Contains(chain.Name, searchStr) || strings.Contains(chain.Chain, searchStr)) {
                wf.NewItem(chain.Name).Subtitle(chain.Chain).Autocomplete(chain.Chain).Arg(chain.Explorers[0].Url).Icon(&icon).Valid(true)
            } 
        }
    }


	wf.WarnEmpty("No matching chains", "Try a different query?")
    wf.SendFeedback()
}