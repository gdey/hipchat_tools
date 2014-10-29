package capabilities

import "net/url"

type Vendor struct {
	URL  *url.URL `json:"url"`
	Name string   `json:"name"`
}

type Links struct {
	Self     *url.URL `json:"self"`
	Homepage *url.URL `json:"homepage"`
}

type OAuth2Provider struct {
	TokenURL         *url.URL `json:"tokenUrl"`
	AuthorizationUrl string   `json:"authorizationUrl"`
}
type OpenIdProvider struct {
	URL       *url.URL `json:"url"`
	Name      string   `json:"name"`
	LogoutUrl string   `json:"logotUrl"`
}
type WebHook struct {
	URL     *url.URL `json:"url"`
	Pattern string   `json:"pattern"`
	Event   string   `json:"event"`
	Name    string   `json:"name"`
}
type Scope string
type ScopeDefinination map[string]string
type HipchatAPIProvider struct {
	URL             *url.URL                    `json:"url"`
	AvailableScopes map[Scope]ScopeDefinination `json:"availableScopes"`
}
type Configurable struct {
	URL                     *url.URL `json:"url"`
	AllowAccessToRoomAdmins bool     `json:"allowAccessToRoomAdmins"`
}
type OAuth2Consumer struct {
	RedirectionUrls *url.URL `json:"redirectionUrls"`
}

type HipchatApiConsumer struct {
	Scopes   []string `json:"scopes"`
	FromName string   `json:"fromName"`
}

type Installable struct {
	InstalledUrl   *url.URL `json:"installedUrl"`
	UninstalledUrl *url.URL `json:"uninstalledUrl"`
	AllowRoom      bool     `json:"allowRoom"`
	AllowGlobal    bool     `json:"allowGlobal"`
	CallbackUrl    *url.URL `json:"callbackUrl"`
}
type Capabilities struct {
	OAuth2Provider     *OAuth2Provider     `json:"oauth2Provider"`
	OpenIdProvider     *OpenIdProvider     `json:"openIdProvider"`
	WebHooks           []*WebHook          `json:"webhook"`
	HipchatAPIProvider *HipchatAPIProvider `json:"hipechatAPIProvider"`
	Configurable       *Configurable       `json"configurable"`
	OAuth2Consumer     *OAuth2Consumer     `json:"oauth2Consumer"`
	HipchatApiConsumer *HipchatApiConsumer `json:"hipchatApiConsumer"`
	Installable        *Installable        `json:"installable"`
}
