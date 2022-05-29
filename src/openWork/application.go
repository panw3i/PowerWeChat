package officialAccount

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	media2 "github.com/ArtisanCloud/PowerWeChat/v2/src/basicService/media"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/auth"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/base"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/user"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/user/tag"
)

type OfficialAccount struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken

	Config *kernel.Config

	User *user.Client
	Tag  *tag.Client

	Media *media2.Client
}

type UserConfig struct {
	AppID  string
	Secret string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface

	HttpDebug bool
	Debug     bool
}

type Log struct {
	Level string
	File  string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewOfficialAccount(config *UserConfig) (*OfficialAccount, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: userConfig,
		DefaultConfig: &object.HashMap{
			"http": &object.HashMap{
				"base_uri": "https://api.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &OfficialAccount{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Auth --------------
	app.AccessToken = auth.RegisterProvider(app)
	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- register User --------------
	app.User, app.Tag = user.RegisterProvider(app)

	//-------------- media --------------
	app.Media = media2.RegisterProvider(app)

	return app, err
}

func (app *OfficialAccount) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *OfficialAccount) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *OfficialAccount) GetConfig() *kernel.Config {
	return app.Config
}

func (app *OfficialAccount) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "Config":
		return app.Config

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{

		"app_id": userConfig.AppID,
		"secret": userConfig.Secret,

		"response_type": userConfig.ResponseType,
		"log": &object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"cache":      userConfig.Cache,
		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,
	}

	return config, nil

}
