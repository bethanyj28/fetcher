package fetcher

import (
	"encoding/json"
	"reflect"

	"github.com/bethanyj28/fetcher/internal/pfclient"
	"github.com/pkg/errors"
)

// Fetcher has all the information needed to retrieve a pet from PetFinder
type Fetcher struct {
	client   *pfapi.Client
	criteria *pfclient.Criteria
	interval time.Duration
}

// NewFetcher creates a new instance of Fetcher
func NewFetcher(apiKey, secret string, criteria pfclient.Criteria, interval time.Duration) (*Fetcher, error) {
	pfClient, err := pfclient.NewClient(apiKey, secret)
	if err != nil {
		return nil, errors.New(err, "failed to create new PetFinder client")
	}

	return &Fetcher{client: pfClient, criteria: criteria, interval: interval}, nil
}

// AddCriteria adds criteria that will need to be searched
func (f *Fetcher) AddCriteria(criteria *pfclient.Criteria) error {
	currentCriteria, err := pfclient.mapToCriteria(f.criteria)
	if err != nil {
		return errors.Wrap(err, "failed to convert current criteria to map")
	}

	newCriteria, err := pfclient.mapToCriteria(criteria)
	if err != nil {
		return errors.Wrap(err, "failed to convert current criteria to map")
	}

	for k, v := range newCriteria {
		// make sure there is a place to put the value
		_, ok := currentCriteria[k]

		// determine whether we should append or overwrite
		vType := reflect.TypeOf(v)
		if vType.Kind() == reflect.Slice && ok { // if it's a slice and there are already entries, append
			currentCriteria[k] = append(currentCriteria[k], v)
			continue
		}

		currentCriteria[k] = v
	}

	updatedCriteria, err := pfclient.mapToCriteria(currentCriteria)
	if err != nil {
		return errors.Wrap(err, "failed to convert map to criteria")
	}

	f.criteria = updatedCriteria

	return nil
}

// UpdateInterval updates the time between checks for pets
func (f *Fetcher) UpdateInterval(interval time.Duration) {
	f.interval = interval
}

// Find determines if there are available pets with the specified criteria
func (f *Fetcher) Find() {}
