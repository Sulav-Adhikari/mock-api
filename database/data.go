package database

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jaswdr/faker"
)

type database struct {
	sync.Mutex
	Offers []Offer
}

var Database database
var Faker faker.Faker

func init() {
	Database.Offers = make([]Offer, 0, 1024)
	fmt.Println("database initialized with 1024 capacity")

	Faker = faker.New()
}

func getNextOfferId() int {
	Database.Lock()
	defer Database.Unlock()
	return len(Database.Offers) + 1
}

type JSONDate time.Time

func (t JSONDate) MarshalJSON() ([]byte, error) {
	date := fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02"))
	return []byte(date), nil
}

type Offer struct {
	OfferID          int       `json:"offer_id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Code             string    `json:"code"`
	Featured         string    `json:"featured"`
	Source           string    `json:"source"`
	URL              string    `json:"url"`
	AffiliateLink    string    `json:"affiliate_link"`
	ImageURL         string    `json:"image_url"`
	BrandLogo        string    `json:"brand_logo"`
	Type             string    `json:"type"`
	Store            string    `json:"store"`
	MerchantHomePage string    `json:"merchant_home_page"`
	Categories       []string  `json:"categories"`
	StartDate        JSONDate  `json:"start_date"`
	EndDate          JSONDate  `json:"end_date"`
	Status           string    `json:"status"`
	PrimaryLocation  []string  `json:"primary_location"`
	Rating           int       `json:"rating"`
	Label            string    `json:"label"`
	Language         string    `json:"language"`
	DeepLink         string    `json:"deeplink"`
	CashbackLink     string    `json:"cashback_link"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

func GetOffers() []Offer {
	return Database.Offers
}

func AddNewFakeOffer() {

	fmt.Println("adding a new fake offer")

	lorem := Faker.Lorem()

	offer := Offer{
		OfferID:         getNextOfferId(),
		Title:           lorem.Word(),
		Description:     lorem.Sentence(5),
		Code:            strings.ToUpper(Faker.RandomStringWithLength(5)),
		Featured:        Faker.RandomStringElement([]string{"Yes", "No"}),
		Source:          Faker.Company().Name(),
		URL:             Faker.Internet().URL(),
		AffiliateLink:   Faker.Internet().URL(),
		ImageURL:        Faker.Internet().URL(),
		BrandLogo:       "",
		Type:            Faker.RandomStringElement([]string{"Code", "Deals"}),
		Store:           Faker.Company().Name(),
		StartDate:       JSONDate(time.Now()),
		EndDate:         JSONDate(time.Now().Add(time.Duration(24 * time.Hour))),
		Status:          "new",
		PrimaryLocation: []string{Faker.Address().Country()},
		Rating:          Faker.IntBetween(0, 100),
		Label:           lorem.Sentence(10),
		Language:        "English",
		DeepLink:        Faker.Internet().URL(),
		CashbackLink:    Faker.Internet().URL(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	Database.Lock()
	defer Database.Unlock()
	Database.Offers = append(Database.Offers, offer)

}

func UpdateOffer() {

	fmt.Println("updating an existing fake offer")
	Database.Lock()
	defer Database.Unlock()

	itemIndex := Faker.IntBetween(0, len(Database.Offers) - 1)

	Database.Offers[itemIndex].Description = Faker.Lorem().Sentence(5)
	Database.Offers[itemIndex].Code = strings.ToUpper(Faker.RandomStringWithLength(5))
	Database.Offers[itemIndex].UpdatedAt = time.Now()
	Database.Offers[itemIndex].Status = "updated"

}

func SuspendOffer() {
	fmt.Println("suspending an existing fake offer")

	Database.Lock()
	defer Database.Unlock()

	itemIndex := Faker.IntBetween(0, len(Database.Offers) - 1)

	Database.Offers[itemIndex].Status = "suspended"
	Database.Offers[itemIndex].UpdatedAt = time.Now()

}

func FetchOfferAfter(d time.Time) []Offer {

	Database.Lock()
	defer Database.Unlock()

	validOffers := []Offer{}

	for _, v := range Database.Offers {
		if v.UpdatedAt.After(d) || v.CreatedAt.After(d) {
			validOffers = append(validOffers, v)
		}
	}

	return validOffers

}
