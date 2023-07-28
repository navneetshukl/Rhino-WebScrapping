package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factID, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}
		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}
		allFacts = append(allFacts, fact)
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	/*for _, data := range allFacts {
		fmt.Println("Id is : ", data.ID)
		fmt.Println("Description is : ", data.Description)

	}*/

	/*enc:=json.NewEncoder(os.Stdout)
	enc.SetIndent(""," ")
	enc.Encode(allFacts)*/

	writeJSON(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}

	_ = ioutil.WriteFile("rhinofacts.json", file, 0644)
}
