package facebook

// Facebook default params
const (
	PictureDefaultHeight uint   = 500
	PictureDefaultWidth  uint   = 500
	PictureSquareType    string = "square"
	PictureDefaultType   string = PictureSquareType
)

// Configer interface
type Configer interface {
	ApplicationID() string
	AppSecretKey() string
	PictureHeight() uint
	PictureWidth() uint
	PictureType() string
}

// Config struct
type Config struct {
	AppID     string
	SecretKey string
	Picture   *struct {
		Height uint   `default:"500"`
		Width  uint   `default:"500"`
		Type   string `default:"square"`
	}
}

// ApplicationID return facebook application id
func (c *Config) ApplicationID() string {
	return c.AppID
}

// AppSecretKey return facebook application secret key
func (c *Config) AppSecretKey() string {
	return c.SecretKey
}

// PictureHeight return facebook profile picture height
func (c *Config) PictureHeight() uint {
	if c.Picture == nil {
		return PictureDefaultHeight
	}
	return c.Picture.Height
}

// PictureWidth return facebook profile picture width
func (c *Config) PictureWidth() uint {
	if c.Picture == nil {
		return PictureDefaultWidth
	}
	if c.PictureType() == PictureSquareType {
		c.Picture.Width = c.PictureHeight()
	}
	return c.Picture.Width
}

// PictureType return facebook profile picture height
func (c *Config) PictureType() string {
	if c.Picture == nil || c.Picture.Type == "" {
		return PictureDefaultType
	}
	return c.Picture.Type
}
