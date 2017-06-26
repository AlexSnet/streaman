package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/grafov/m3u8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	uri, uri_e := url.Parse("https://bitdash-a.akamaihd.net/content/MI201109210084_1/m3u8s/f08e80da-bf1d-4e3d-8899-f0f6155f6efa.m3u8")
	if uri_e != nil {
		panic(uri_e)
	}

	resp, err := http.Get(uri.String())
	check(err)
	defer resp.Body.Close()

	p, listType, err := m3u8.DecodeFrom(resp.Body, true)
	if err != nil {
		panic(err)
	}
	switch listType {
	case m3u8.MEDIA:
		mediapl := p.(*m3u8.MediaPlaylist)
		fmt.Println(mediapl)
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		for _, v := range masterpl.Variants {
			u, e := url.Parse(v.URI)
			if e != nil {

			} else {
				nu := uri.ResolveReference(u)
				fmt.Println(nu.String())
			}
		}
	}
}
