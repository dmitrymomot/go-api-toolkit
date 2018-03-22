package application

// BaseController structure
type BaseController struct {
	app Apper
}

// App returns application instance
func (c *BaseController) App() Apper {
	return c.app
}

// SetApp sets application instance
func (c *BaseController) SetApp(app Apper) {
	c.app = app
}
