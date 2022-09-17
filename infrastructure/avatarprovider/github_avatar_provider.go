package avatarprovider

import (
	"encoding/json"
	"fmt"
	"go-clean-architecture-sample/application/common/errors"
	"io/ioutil"
	"net/http"
)

type GithubAvatarProvider struct {
	avatarUrlCache map[string]string
}

func NewGithubAvatarProvider() *GithubAvatarProvider {
	return &GithubAvatarProvider{
		avatarUrlCache: make(map[string]string),
	}
}

func (p *GithubAvatarProvider) GetAvatarUrlByEmail(email string) (string, error) {
	avatarUrl, cacheHit := p.avatarUrlCache[email]

	if !cacheHit {
		httpResponse, err := http.Get(fmt.Sprintf("https://api.github.com/search/users?q=%s%%20in:email", email))
		if err != nil {
			return "", errors.ErrInternal
		}

		defer httpResponse.Body.Close()

		body, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return "", errors.ErrInternal
		}

		var githubSearchResponse struct {
			Items []struct {
				AvatarUrl string `json:"avatar_url"`
			} `json:"items"`
		}

		err = json.Unmarshal(body, &githubSearchResponse)
		if err != nil {
			return "", errors.ErrInternal
		}

		if len(githubSearchResponse.Items) != 1 {
			return "", nil
		}

		avatarUrl = githubSearchResponse.Items[0].AvatarUrl

		p.avatarUrlCache[email] = avatarUrl
	}

	return avatarUrl, nil
}
