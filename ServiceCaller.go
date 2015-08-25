package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"log"
    "net/http"
	"strings"
    //"os"
)

const exampleTheme = `{
    "Items":
	[
		{
			"Activated":true,
			"AppUserID":104,
			"Avatar":"pic19",
			"DisplayName":"Nate Goins",
			"Email":"Nate@meetball.com",
			"FacebookID":"100004264873142",
			"Favorite":true,
			"FirstName":"Natron",
			"Handle":"NatronGoins",
			"LastName":"Goins",
			"PhoneNumber":"3129099646",
			"StatusID":1
		},
		{
			"Activated":true,
			"AppUserID":907,
			"Avatar":null,
			"DisplayName":"Derek Hearn",
			"Email":"HoopDoop@boop.com",
			"FacebookID":"565762384",
			"Favorite":false,
			"FirstName":"Derek",
			"Handle":"Vicda",
			"LastName":"Hearn2",
			"PhoneNumber":"7734390751",
			"StatusID":1
		}
	],
	"MBResult":
	{
		"DeveloperErrorMsg":"",
		"ErrorType":1,
		"FriendlyErrorMsg":"",
		"Success":true
	}
}`

type retVal struct {
	Items []gf_item 
	MBResult result //`json:"MBResult"`
}

type result struct {
	DeveloperErrorMsg string //`json:"DeveloperErrorMsg`
	FriendlyErrorMsg string //`json:"FriendlyErrorMsg`
	Success bool //`json:"Success"`
}

type gf_item struct {
	//Activated bool
	AppUserID int
	//Avatar string
	DisplayName string
	//Email string
	//FacebookID string
	//Favorite bool
	FirstName string
	Handle string
	LastName string
	//PhoneNumber string
	//StatusID int
}

func main() {
	url := "http://wsdev.meetball.com/6.0/Service.svc/json/Friends/{SessionID}"
	session := "b58aeb59-16e9-4d77-a95a-01fa81531b08"
	if isDev(url){
		session = "AE328CAC-1BC9-47A4-82EE-014F98831197"
	}
	
	url = strings.Replace(url, "{SessionID}", session, 1)	
	
	fmt.Printf(url + "\n")
	
	resp, err := http.Get(url)
	
	if err != nil {
		log.Fatal(err)
    }
	
    defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal(err)
    }
	
	var s retVal
	err = json.Unmarshal(body, &s)
	
	if err != nil {
		log.Fatal(err)
    }
	
	if s.MBResult.Success {
		fmt.Printf("Success!!!")
	} else {
		fmt.Println("Something went awry")
		fmt.Println(s.MBResult.DeveloperErrorMsg)
	}
	
	fmt.Println(s)
	
	fmt.Println(string(body))
	fmt.Println("", s.Items)
	
	for index, item := range s.Items {
		fmt.Println(index, item)
	}
	
}

func isDev(url string)bool{
	return strings.Contains(url,"wsdev")
}

