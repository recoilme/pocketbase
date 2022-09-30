package auth

import (
	"golang.org/x/oauth2"
)

var _ Provider = (*Telegram)(nil)

// NameTelegram is the unique name of the Telegram provider.
const NameTelegram string = "telegram"

// Telegram allows authentication via Telegram OAuth2.
type Telegram struct {
	*baseProvider
}

// NewFacebookProvider creates new Facebook provider instance with some defaults.
func NewTelegramProvider() *Telegram {
	return &Telegram{&baseProvider{
		scopes:     []string{"email"},
		authUrl:    "https://www.facebook.com/dialog/oauth",
		tokenUrl:   "https://graph.facebook.com/oauth/access_token",
		userApiUrl: "https://graph.facebook.com/me?fields=name,email,picture.type(large)",
	}}
}

// FetchAuthUser returns an AuthUser instance based on the Facebook's user api.
func (p *Telegram) FetchAuthUser(token *oauth2.Token) (*AuthUser, error) {
	// https://developers.facebook.com/docs/graph-api/reference/user/
	rawData := struct {
		Id      string
		Name    string
		Email   string
		Picture struct {
			Data struct{ Url string }
		}
	}{}

	if err := p.FetchRawUserData(token, &rawData); err != nil {
		return nil, err
	}

	user := &AuthUser{
		Id:        rawData.Id,
		Name:      rawData.Name,
		Email:     rawData.Email,
		AvatarUrl: rawData.Picture.Data.Url,
	}

	return user, nil
}
