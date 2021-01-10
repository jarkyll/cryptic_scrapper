package service

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/kamva/mgm/v3"
	"golang-boilerplate/models"
	"regexp"
	"strconv"
	"strings"
)

const domainName = "https://www.mtggoldfish.com"

func Test() {
	c := colly.NewCollector()


	c.OnHTML(".archetype-tile-container", func(decksContainer *colly.HTMLElement) {
		// If attribute class is this long string return from callback
		// As this a is irrelevant


		decksContainer.ForEach(".archetype-tile", func(index int, deck *colly.HTMLElement) {
			deckCollector := colly.NewCollector()

			url := deck.ChildAttr(".archetype-tile-description-wrapper .archetype-tile-title .deck-price-online a", "href")
			//text := deck.ChildText(".archetype-tile-description-wrapper .archetype-tile-title .deck-price-online a")

			re := regexp.MustCompile(`(.*)#`)
			results := re.FindSubmatch([]byte(url))
			baseUrl := string(results[1])


			baseUrl = domainName + baseUrl + "/decks"

			deckCollector.OnHTML("table tbody", func(deckList *colly.HTMLElement) {
				deckList.ForEach("tr td:nth-child(2) a", func(index int, deck *colly.HTMLElement) {

					urlTest := deck.Attr("href")

					dc := colly.NewCollector()

					deckModel := &models.Deck{}
					//deckModel2 := make(map[string][]models.Card)


					dc.OnHTML(".deck-table-buttons-container .tab-content .active tbody", func(deckPage *colly.HTMLElement) {

						//var listType string;
						var currentTypeList []models.Card
						var cardType string
						deckPage.ForEach("tr", func(index int, row *colly.HTMLElement) {
							card := &models.Card{}

							// for handling the title type
							row.ForEach("th", func(index int, cardInfo *colly.HTMLElement) {

								if cardType != "" {
									deckModel.SetField(cardType, currentTypeList)
									currentTypeList = nil

								}

								trimmedType := strings.TrimSpace(row.Text)
								cardType = strings.Split(trimmedType, "\n")[0]
							})


							// for handling the cards and count
							row.ForEach("td", func(index int, cardInfo *colly.HTMLElement) {
								switch index {
								case 0:
									number := strings.TrimSpace(cardInfo.Text)
									card.Count, _ = strconv.Atoi(number)
								case 1:
									card.Name = strings.TrimSpace(cardInfo.Text)
									currentTypeList = append(currentTypeList, *card)
								default:
									return
								}
								if index > 1 {
									return
								}

							})



						})

						fmt.Println("or did i finish here?")
						mgm.Coll(deckModel).Create(deckModel)
					})


					dc.Visit(domainName + urlTest)

				})
			})
			deckCollector.Visit(baseUrl)

		})
	})

	c.Visit(domainName + "/metagame/commander")
}
