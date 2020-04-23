package main

import (
	"encoding/json"
	"encoding/xml"
	"exporter/exporter"
	"io/ioutil"
	"strconv"
)

func main() {
	fb := new(exporter.Facebook)
	twtr := new(exporter.Twitter)
	ldin := new(exporter.LinkedIn)

	feedf := make(map[string]string)
	for i := range fb.Feed() {
		feedf[strconv.Itoa(i+1)] = fb.Feed()[i]
	}
	f, _ := json.MarshalIndent(feedf, "", " ")
	_ = ioutil.WriteFile("Facebook.json", f, 0644)
	fcb, _ := xml.MarshalIndent(fb.Feed(), "", " ")
	_ = ioutil.WriteFile("Facebook.xml", fcb, 0644)

	feedt := make(map[string]string)
	for i := range twtr.Feed() {
		feedt[strconv.Itoa(i+1)] = twtr.Feed()[i]
	}
	t, _ := json.MarshalIndent(feedt, "", " ")
	_ = ioutil.WriteFile("witter.json", t, 0644)
	tw, _ := xml.MarshalIndent(twtr, "", " ")
	_ = ioutil.WriteFile("Twitter.xml", tw, 0644)

	feedl := make(map[string]string)
	for i := range ldin.Feed() {
		feedl[strconv.Itoa(i+1)] = ldin.Feed()[i]
	}
	in, _ := json.MarshalIndent(feedl, "", " ")
	_ = ioutil.WriteFile("LinkedIn.json", in, 0644)
	lin, _ := xml.MarshalIndent(ldin, "", " ")
	_ = ioutil.WriteFile("LinkedIn.xml", lin, 0644)

}