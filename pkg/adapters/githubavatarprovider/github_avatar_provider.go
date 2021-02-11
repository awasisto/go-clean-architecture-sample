package githubavatarprovider

import (
	"encoding/json"
	"fmt"
	"golang-clean-architecture-sample/pkg/core"
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

func (g *GithubAvatarProvider) GetAvatarUrlByEmail(email string) (string, error) {
	avatarUrl, cacheHit := g.avatarUrlCache[email]

	if !cacheHit {
		httpResponse, err := http.Get(fmt.Sprintf("https://api.github.com/search/users?q=%s%%20in:email", email))
		if err != nil {
			return "", core.ErrInternal
		}

		defer httpResponse.Body.Close()

		body, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return "", core.ErrInternal
		}

		var githubSearchResponse struct {
			Items []struct {
				AvatarUrl string `json:"avatar_url"`
			} `json:"items"`
		}

		err = json.Unmarshal(body, &githubSearchResponse)
		if err != nil {
			return "", core.ErrInternal
		}

		if len(githubSearchResponse.Items) != 1 {
			return "", nil
		}

		avatarUrl = githubSearchResponse.Items[0].AvatarUrl

		g.avatarUrlCache[email] = avatarUrl
	}

	return avatarUrl, nil
}
