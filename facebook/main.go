package facebook

// Facebooker facebook struct interface
type Facebooker interface {
	Profile(accessToken string) (*FbProfile, error)
	ProfileWithImage(accessToken string) (*FbProfile, error)
}

// Facebook structure
type Facebook struct {
	config Configer
}

// Profile returns facebook profile
func (f *Facebook) Profile(accessToken string) (*FbProfile, error) {
	p, err := getProfile(accessToken)
	if err != nil {
		return nil, err
	}
	return &FbProfile{p, nil}, nil
}

// ProfileWithImage returns facebook profile with picture
func (f *Facebook) ProfileWithImage(accessToken string) (*FbProfile, error) {
	p, err := getProfile(accessToken)
	if err != nil {
		return nil, err
	}
	result := &FbProfile{p, nil}
	userpic, err := getProfilePicture(accessToken, f.config.PictureHeight(), f.config.PictureWidth(), f.config.PictureType(), false)
	if err != nil {
		return nil, err
	}
	result.Picture = userpic
	return result, nil
}

// Setup returns configured facade of package
func Setup(c Configer) Facebooker {
	return &Facebook{c}
}
