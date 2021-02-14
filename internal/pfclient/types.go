package pfclient

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Criteria is the search criteria for a pet
type Criteria struct {
	Type         []string `json:"type,omitempty"`
	Breed        []string `json:"breed,omitempty"`
	Color        []string `json:"color,omitempty"`
	Age          []string `json:"age,omitempty"`
	Size         []string `json:"size,omitempty"`
	Gender       []string `json:"gender,omitempty"`
	Coat         []string `json:"coat,omitempty"`
	GoodWith     []string `json:"good_with,omitempty"` // i.e. children, cats, dogs
	HouseTrained bool     `json:"house_trained,omitempty"`
	Declawed     bool     `json:"declawed,omitempty"`
	Location     string   `json:"location"`
	Distance     int32    `json:"distance,omitempty"`
}

func criteriaToMap(c *Criteria) (map[string]interface{}, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal struct")
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal bytes to map")
	}

	return m, nil
}

func mapToCriteria(m map[string]interface{}) (*Criteria, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal map")
	}

	var c *Criteria
	if err := json.Unmarshal(b, c); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal bytes to criteria")
	}

	return c, nil
}

// pfGetAnimalResponse is a response from the petfinder api
type pfGetAnimalResponse struct {
	animals    []Animal   `json:"animals"`
	pagination Pagination `json:"pagination"`
}

// Animal is petfinder's definition of an animal
type Animal struct {
	ID             int64        `json:"id"`
	OrganizationID int64        `json:"organization_id"`
	URL            string       `json:"url"`
	Type           string       `json:"type"`
	Species        string       `json:"species"`
	Breeds         AnimalBreeds `json:"breeds"`
	Colors         Colors       `json:"colors"`
	Age            string       `json:"age"`
	Gender         string       `json:"gender"`
	Size           string       `json:"size"`
	Coat           string       `json:"coat"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Photos         []Photo      `json:"photos"`
	Status         string       `json:"status"`
	Attributes     Attributes   `json:"attributes"`
	Environment    Environment  `json:"environment"`
	Tags           []string     `json:"tags"`
	Contact        Contact      `json:"contact"`
	PublishedAt    time.Time    `json:"published_at"`
	Distance       float64      `json:"distance"`
}

// AnimalBreeds represents the breeds of a specific animal
type AnimalBreeds struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
	Mixed     bool   `json:"mixed"`
	Unknown   bool   `json:"unknown"`
}

// Colors represent the colors of an animal
type Colors struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
	Tertiary  string `json:"tertiary"`
}

// Photo is the various sizes of a single photo
type Photo struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
	Full   string `json:"full"`
}

// Attributes are details on the animal
type Attributes struct {
	SpayedNeutered bool `json:"spayed_neutered"`
	HouseTrained   bool `json:"house_trained"`
	Declawed       bool `json:"declawed"`
	SpecialNeeds   bool `json:"special_needs"`
	ShotsCurrent   bool `json:"shots_current"`
}

// Environment is the ideal environment for the animal
type Environment struct {
	Children bool `json:"children"`
	Dogs     bool `json:"dogs"`
	Cats     bool `json:"cats"`
}

// Contact is the contact info to ask about the animal
type Contact struct {
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
}

// Address represents the address for an animal
type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Postcode string `json:"postcode"`
	Country  string `json:"country"`
}

// Pagination is the paginated results for a page
type Pagination struct {
	CountPerPage int64 `json:"count_per_page"`
	TotalCount   int64 `json:"total_count"`
	CurrentPage  int64 `json:"current_page"`
	TotalPages   int64 `json:"total_pages"`
}

// pfBreedsResponse is the petfinder get breeds response
type pfBreedsResponse struct {
	Breeds []Breed `json:"breed"`
}

// Breed is the breed of animal available on petfinder (not for a specific animal)
type Breed struct {
	Name string `json:"name"`
}

// pfTypesResponse is various possible attributes for all animal types
type pfTypesResponse struct {
	Types []Type `json:"types"`
}

// pfTypeResponse is the possible attributes for a single animal type
type pfTypeResponse struct {
	Type
}

// Type represents the type of animal and various attributes
type Type struct {
	Name    string   `json:"name"`
	Coats   []string `json:"coats"`
	Colors  []string `json:"colors"`
	Genders []string `json:"genders"`
}
