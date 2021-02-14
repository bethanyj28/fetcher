package pfclient

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var defaultBaseURL = "https://api.petfinder.com/v2"

// PFClient handles making requests to petfinder
type PFClient struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new PFClient with the provided credentials
func NewClient(apiKey, secret string) *PFClient {
	url := os.Getenv("PF_URL")
	if url == "" {
		url = defaultBaseURL
	}

	oauthConfig := &clientCredentials.Config{
		ClientID:     apiKey,
		ClientSecret: secret,
		TokenURL:     fmt.Sprintf("%s/oauth2/token", url),
	}

	return PFClient{httpClient: oauthConfig.Client(oauth.NoContext), baseURL: url}
}

// GetAnimalsByCriteria fetches animals meeting a certain criteria
func (p *PFClient) GetAnimalsByCriteria(criteria criteria) (*Animal, error) {
	return nil, nil
}

// GetAnimalTypes returns all the types of animals are available and the available properties
func (p *PFClient) GetAnimalTypes() ([]*Type, error) {
	return nil, nil
}
