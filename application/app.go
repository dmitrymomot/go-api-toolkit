package application

import (
	"github.com/dmitrymomot/go-api-toolkit/auth"
	"github.com/dmitrymomot/go-api-toolkit/facebook"
	"github.com/dmitrymomot/go-api-toolkit/mailer"
	"github.com/dmitrymomot/go-api-toolkit/pusher"
	"github.com/dmitrymomot/go-api-toolkit/rbac"
	"github.com/dmitrymomot/go-api-toolkit/storage"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Apper application instance interface
type Apper interface {
	Auth() auth.JWT
	Config() Configer
	DB() *gorm.DB
	Fb() facebook.Facebooker
	Mail() mailer.Mailer
	Push() pusher.Senderer
	RBAC() rbac.RBACer
	Router() *gin.Engine
	Storage() storage.Clienter
	IsMasterPassword(p string) bool
}

// App is an application instance structure
type App struct {
	Authorization auth.JWT
	Configuration Configer
	Database      *gorm.DB
	Engine        *gin.Engine
	FbClient      facebook.Facebooker
	MailClient    mailer.Mailer
	PushClient    pusher.Senderer
	Rbac          rbac.RBACer
	StorageClient storage.Clienter
}

// Auth returns auth.JWT instance
func (a *App) Auth() auth.JWT {
	return a.Authorization
}

// Config returns application config
func (a *App) Config() Configer {
	return a.Configuration
}

// DB returns gorm db instance
func (a *App) DB() *gorm.DB {
	return a.Database
}

// Fb returns gorm facebook client
func (a *App) Fb() facebook.Facebooker {
	return a.FbClient
}

// Mail returns mail client instance
func (a *App) Mail() mailer.Mailer {
	return a.MailClient
}

// Push returns push client instance
func (a *App) Push() pusher.Senderer {
	return a.PushClient
}

// RBAC returns RBAC instance
func (a *App) RBAC() rbac.RBACer {
	return a.Rbac
}

// Router returns gin engine instance
func (a *App) Router() *gin.Engine {
	return a.Engine
}

// Storage returns storage client
func (a *App) Storage() storage.Clienter {
	return a.StorageClient
}

// IsMasterPassword checks master password
func (a *App) IsMasterPassword(p string) bool {
	return a.Config().MasterPassword(p)
}
