package facebook

import (
	fb "github.com/huandu/facebook"
)

// FbProfile model
type FbProfile struct {
	*data
	Picture *picture `json:"picture"`
}

// PictureURL returns profile picture URL
func (p *FbProfile) PictureURL() string {
	if p.Picture == nil {
		return ""
	}
	return p.Picture.URL
}

// Profile facebook profile structure
type data struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Address    string `json:"address,omitempty"`
	Birthday   string `json:"birthday,omitempty"`
	Gender     string `json:"gender,omitempty"`
	IsVerified bool   `json:"is_verified"`
}

// picture structure
type picture struct {
	Height       int
	IsSilhouette bool
	URL          string
	Width        string
}

func getProfile(accessToken string) (*data, error) {
	profile := &data{}
	res, err := fb.Get("/me", fb.Params{
		"access_token": accessToken,
		"fields": []string{
			"id",
			"first_name",
			"last_name",
			"name",
			"email",
			"address",
			"birthday",
			"gender",
			"is_verified",
		},
	})
	if err != nil {
		return nil, err
	}
	if err := res.Decode(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func getProfilePicture(accessToken string, height, width uint, imgType string, redirect bool) (*picture, error) {
	res, err := fb.Get("/me/picture", fb.MakeParams(map[string]interface{}{
		"access_token": accessToken,
		"height":       height,
		"width":        width,
		"type":         imgType,
		"redirect":     redirect,
		"fields":       []string{"url", "height", "width", "is_silhouette"},
	}))
	if err != nil {
		return nil, err
	}
	picture := &struct {
		Data *picture
	}{}
	if err := res.Decode(picture); err != nil {
		return nil, err
	}
	return picture.Data, nil
}
