package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jzelinskie/geddit"
)

type Story struct {
	title string
	url   string
}

func init() {

	//loginToReddit()
}

func loginToReddit() {
	// Login to reddit
	var err error
	var session *geddit.LoginSession
	session, err = geddit.NewLoginSession(
		"saied63",
		"Rd7111443",
		"gedditAgent v1",
	)
	if err != nil {
		log.Println("loggin failed")
		os.Exit(1)
	} else {
		LOggedIn(session)
	}
}

type category struct {
	cat string
	ch  chan []Story
}

var AllCategory []category = []category{
	{cat: "it"}, {cat: "car"}, {cat: "flower"}, {cat: "iran"}, {cat: "usa"}, {cat: "hotel"}, {cat: "marga"}, {cat: "msa"}, {cat: "apple"}, {cat: "hot coin"}, {cat: "crypto"}, {cat: "calculator"},
}

func (cat *category) addChanells() {
	cat.ch = make(chan []Story, 10)
}

func LOggedIn(session *geddit.LoginSession) {

	for i := 0; i < len(AllCategory); i++ {
		AllCategory[i].addChanells()
	}
	for i := 0; i < 30; i++ {
		for i := 0; i < len(AllCategory); i++ {
			go setListingOption(AllCategory[i].cat, session, AllCategory[i].ch)
		}
	}

	for i := 0; i < len(AllCategory); i++ {
		select {
		case result := <-AllCategory[i].ch:
			fmt.Println(result)
		}

	}

	for i := 0; i < len(AllCategory); i++ {
		close(AllCategory[i].ch)
	}

}

func setListingOption(category string, session *geddit.LoginSession, ch chan<- []Story) {
	// Set listing options

	//defer close(ch)

	storeies := []Story{}
	subOpts := geddit.ListingOptions{
		Limit: 100,
	}
	submissions, err := session.SubredditSubmissions(category, geddit.NewSubmissions, subOpts)
	if err != nil {
		ch <- nil

	}

	for _, s := range submissions {
		newStory := Story{
			title: s.Title,
			url:   s.URL,
		}
		storeies = append(storeies, newStory)
	}

	ch <- storeies

}

func printToFile(ch <-chan []Story) {

	file, err := os.Create("reditInfo.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(AllCategory); i++ {
		stories := <-AllCategory[i].ch
		for _, story := range stories {
			fmt.Printf("TITLE : %s URL : %s \n\n", story.title, story.url)
			fmt.Fprintf(file, "Title : \n\n %s , url : \n\n %s \n\n", story.title, story.url)
		}
	}
}
