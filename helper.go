package main

import (
	proto "gitlab.com/and07/rss2website/proto"
)

func getRSSList() proto.RssList {
	//https://gist.github.com/stungeye/fe88fc810651174d0d180a95d79a8d97
	/*
		j := `[
			{
				"url": "http://www.foxnews.com",
				"rss": "http://feeds.foxnews.com/foxnews/latest"
			}
		]`

		var list RssList
		if err := json.Unmarshal([]byte(j), &list); err != nil {
			log.Fatal(err)
		}
		return list
	*/

	l := make(map[string]string)
	l["foxnews"] = "http://feeds.foxnews.com/foxnews/latest"
	l["cnn"] = "http://rss.cnn.com/rss/money_topstories.rss"
	l["forbes"] = "https://www.forbes.com/real-time/feed2/"
	return proto.RssList{List: l}
}
