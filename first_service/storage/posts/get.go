package posts

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/go-resty/resty/v2"
)

var results []Meta = []Meta{}

func first25Pages(wg *sync.WaitGroup) {
	defer wg.Done()

	// for i := 0; i < 25; i++ {
	for i := 0; i < 1; i++ {

		strToInt := strconv.Itoa(i)
		URL := `https://gorest.co.in/public/v1/posts?page=` + strToInt

		client := resty.New()

		resp, err := client.R().SetHeader("Content-Type", "application/json").Get(URL)
		if err != nil {
			log.Println("Error while getting posts: ", err.Error())
		}

		datas := Meta{}

		err = json.Unmarshal(resp.Body(), &datas)
		if err != nil {
			log.Println("Error while unmarshalling posts: ", err.Error())
		}

		results = append(results, datas)
	}
}

func second25Pages(wg *sync.WaitGroup) {
	defer wg.Done()

	// for i := 25; i < 50; i++ {
	for i := 1; i < 2; i++ {

		strToInt := strconv.Itoa(i)
		URL := `https://gorest.co.in/public/v1/posts?page=` + strToInt

		client := resty.New()

		resp, err := client.R().SetHeader("Content-Type", "application/json").Get(URL)
		if err != nil {
			log.Println("Error while getting posts: ", err.Error())
		}

		datas := Meta{}

		err = json.Unmarshal(resp.Body(), &datas)
		if err != nil {
			log.Println("Error while unmarshalling posts: ", err.Error())
		}

		results = append(results, datas)
	}
}

func GetPosts() []Meta {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)

	go first25Pages(waitGroup)
	go second25Pages(waitGroup)

	waitGroup.Wait()

	return results
}
