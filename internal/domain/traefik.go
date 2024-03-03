package domain

import (
	"fmt"
	"github.com/traefik/traefik/v3/pkg/types"
	"math"
	"strings"
)

// Configuration is the static configuration.
type Configuration struct {
	EntryPoints EntryPoints        `description:"Entry points definition." json:"entryPoints,omitempty" toml:"entryPoints,omitempty" yaml:"entryPoints,omitempty" export:"true"`
	HTTP        *HTTPConfiguration `description:"http" json:"http,omitempty" toml:"http,omitempty" yaml:"http,omitempty" export:"true"`
	Providers   *Providers         `description:"Providers configuration." json:"providers,omitempty" toml:"providers,omitempty" yaml:"providers,omitempty" export:"true"`
	API         *API               `description:"Enable api/dashboard." json:"api,omitempty" toml:"api,omitempty" yaml:"api,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	Log         *types.TraefikLog  `description:"Traefik log settings." json:"log,omitempty" toml:"log,omitempty" yaml:"log,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
}

func GetDefaultConfiguration() *Configuration {
	return &Configuration{
		Log: &types.TraefikLog{
			Level: "INFO",
		},
		API: &API{
			Insecure:  true,
			Dashboard: true,
		},
		EntryPoints: EntryPoints{
			"web": &EntryPoint{
				Address: "80",
			},
		},
		Providers: &Providers{
			File: &Provider{
				Filename: "./traefik.yml",
			},
		},
		HTTP: &HTTPConfiguration{
			Services: make(map[string]*Service, 10),
			Routers:  make(map[string]*Router, 10),
		},
	}
}

// HTTPConfiguration contains all the HTTP configuration parameters.
type HTTPConfiguration struct {
	Routers  map[string]*Router  `json:"routers,omitempty" toml:"routers,omitempty" yaml:"routers,omitempty" export:"true"`
	Services map[string]*Service `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
}

// Service holds a service configuration (can only be of one type at the same time).
type Service struct {
	LoadBalancer *ServersLoadBalancer `json:"loadBalancer,omitempty" toml:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty" export:"true"`
}

type ServersLoadBalancer struct {
	Servers []Server `json:"servers,omitempty" toml:"servers,omitempty" yaml:"servers,omitempty" label-slice-as-struct:"server" export:"true"`
}

type Server struct {
	URL    string `json:"url,omitempty" toml:"url,omitempty" yaml:"url,omitempty" label:"-"`
	Weight *int   `json:"weight,omitempty" toml:"weight,omitempty" yaml:"weight,omitempty" label:"weight"`
}

// Router holds the router configuration.
type Router struct {
	EntryPoints []string `json:"entryPoints,omitempty" toml:"entryPoints,omitempty" yaml:"entryPoints,omitempty" export:"true"`
	Service     string   `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	Rule        string   `json:"rule,omitempty" toml:"rule,omitempty" yaml:"rule,omitempty"`
	Priority    int      `json:"priority,omitempty" toml:"priority,omitempty,omitzero" yaml:"priority,omitempty" export:"true"`
}

// Providers contains providers configuration.
type Providers struct {
	File *Provider `description:"Enable File backend with default settings." json:"file,omitempty" toml:"file,omitempty" yaml:"file,omitempty" export:"true"`
}

type Provider struct {
	Directory                 string `description:"Load dynamic configuration from one or more .yml or .toml files in a directory." json:"directory,omitempty" toml:"directory,omitempty" yaml:"directory,omitempty" export:"true"`
	Watch                     bool   `description:"Watch provider." json:"watch,omitempty" toml:"watch,omitempty" yaml:"watch,omitempty" export:"true"`
	Filename                  string `description:"Load dynamic configuration from a file." json:"filename,omitempty" toml:"filename,omitempty" yaml:"filename,omitempty" export:"true"`
	DebugLogGeneratedTemplate bool   `description:"Enable debug logging of generated configuration template." json:"debugLogGeneratedTemplate,omitempty" toml:"debugLogGeneratedTemplate,omitempty" yaml:"debugLogGeneratedTemplate,omitempty" export:"true"`
}

// Core configures Traefik core behavior.
type Core struct {
	DefaultRuleSyntax string `description:"Defines the rule parser default syntax (v2 or v3)" json:"defaultRuleSyntax,omitempty" toml:"defaultRuleSyntax,omitempty" yaml:"defaultRuleSyntax,omitempty"`
}

// SetDefaults sets the default values.
func (c *Core) SetDefaults() {
	c.DefaultRuleSyntax = "v3"
}

// API holds the API configuration.
type API struct {
	Insecure           bool `description:"Activate API directly on the entryPoint named traefik." json:"insecure,omitempty" toml:"insecure,omitempty" yaml:"insecure,omitempty" export:"true"`
	Dashboard          bool `description:"Activate dashboard." json:"dashboard,omitempty" toml:"dashboard,omitempty" yaml:"dashboard,omitempty" export:"true"`
	Debug              bool `description:"Enable additional endpoints for debugging and profiling." json:"debug,omitempty" toml:"debug,omitempty" yaml:"debug,omitempty" export:"true"`
	DisableDashboardAd bool `description:"Disable ad in the dashboard." json:"disableDashboardAd,omitempty" toml:"disableDashboardAd,omitempty" yaml:"disableDashboardAd,omitempty" export:"true"`
}

// SetDefaults sets the default values.
func (a *API) SetDefaults() {
	a.Dashboard = true
}

// EntryPoint holds the entry point configuration.
type EntryPoint struct {
	Address          string            `description:"Entry point address." json:"address,omitempty" toml:"address,omitempty" yaml:"address,omitempty"`
	ReusePort        bool              `description:"Enables EntryPoints from the same or different processes listening on the same TCP/UDP port." json:"reusePort,omitempty" toml:"reusePort,omitempty" yaml:"reusePort,omitempty"`
	AsDefault        bool              `description:"Adds this EntryPoint to the list of default EntryPoints to be used on routers that don't have any Entrypoint defined." json:"asDefault,omitempty" toml:"asDefault,omitempty" yaml:"asDefault,omitempty"`
	ProxyProtocol    *ProxyProtocol    `description:"Proxy-Protocol configuration." json:"proxyProtocol,omitempty" toml:"proxyProtocol,omitempty" yaml:"proxyProtocol,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	ForwardedHeaders *ForwardedHeaders `description:"Trust client forwarding headers." json:"forwardedHeaders,omitempty" toml:"forwardedHeaders,omitempty" yaml:"forwardedHeaders,omitempty" export:"true"`
	HTTP             HTTPConfig        `description:"HTTP configuration." json:"http,omitempty" toml:"http,omitempty" yaml:"http,omitempty" export:"true"`
	HTTP2            *HTTP2Config      `description:"HTTP/2 configuration." json:"http2,omitempty" toml:"http2,omitempty" yaml:"http2,omitempty" export:"true"`
	HTTP3            *HTTP3Config      `description:"HTTP/3 configuration." json:"http3,omitempty" toml:"http3,omitempty" yaml:"http3,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
}

// GetAddress strips any potential protocol part of the address field of the
// entry point, in order to return the actual address.
func (ep EntryPoint) GetAddress() string {
	splitN := strings.SplitN(ep.Address, "/", 2)
	return splitN[0]
}

// GetProtocol returns the protocol part of the address field of the entry point.
// If none is specified, it defaults to "tcp".
func (ep EntryPoint) GetProtocol() (string, error) {
	splitN := strings.SplitN(ep.Address, "/", 2)
	if len(splitN) < 2 {
		return "tcp", nil
	}

	protocol := strings.ToLower(splitN[1])
	if protocol == "tcp" || protocol == "udp" {
		return protocol, nil
	}

	return "", fmt.Errorf("invalid protocol: %s", splitN[1])
}

// SetDefaults sets the default values.
func (ep *EntryPoint) SetDefaults() {
	ep.ForwardedHeaders = &ForwardedHeaders{}
	ep.HTTP2 = &HTTP2Config{}
	ep.HTTP2.SetDefaults()
}

// HTTPConfig is the HTTP configuration of an entry point.
type HTTPConfig struct {
	Redirections          *Redirections `description:"Set of redirection" json:"redirections,omitempty" toml:"redirections,omitempty" yaml:"redirections,omitempty" export:"true"`
	Middlewares           []string      `description:"Default middlewares for the routers linked to the entry point." json:"middlewares,omitempty" toml:"middlewares,omitempty" yaml:"middlewares,omitempty" export:"true"`
	TLS                   *TLSConfig    `description:"Default TLS configuration for the routers linked to the entry point." json:"tls,omitempty" toml:"tls,omitempty" yaml:"tls,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	EncodeQuerySemicolons bool          `description:"Defines whether request query semicolons should be URLEncoded." json:"encodeQuerySemicolons,omitempty" toml:"encodeQuerySemicolons,omitempty" yaml:"encodeQuerySemicolons,omitempty"`
}

// HTTP2Config is the HTTP2 configuration of an entry point.
type HTTP2Config struct {
	MaxConcurrentStreams int32 `description:"Specifies the number of concurrent streams per connection that each client is allowed to initiate." json:"maxConcurrentStreams,omitempty" toml:"maxConcurrentStreams,omitempty" yaml:"maxConcurrentStreams,omitempty" export:"true"`
}

// SetDefaults sets the default values.
func (c *HTTP2Config) SetDefaults() {
	c.MaxConcurrentStreams = 250 // https://cs.opensource.google/go/x/net/+/cd36cc07:http2/server.go;l=58
}

// HTTP3Config is the HTTP3 configuration of an entry point.
type HTTP3Config struct {
	AdvertisedPort int `description:"UDP port to advertise, on which HTTP/3 is available." json:"advertisedPort,omitempty" toml:"advertisedPort,omitempty" yaml:"advertisedPort,omitempty" export:"true"`
}

// Redirections is a set of redirection for an entry point.
type Redirections struct {
	EntryPoint *RedirectEntryPoint `description:"Set of redirection for an entry point." json:"entryPoint,omitempty" toml:"entryPoint,omitempty" yaml:"entryPoint,omitempty" export:"true"`
}

// RedirectEntryPoint is the definition of an entry point redirection.
type RedirectEntryPoint struct {
	To        string `description:"Targeted entry point of the redirection." json:"to,omitempty" toml:"to,omitempty" yaml:"to,omitempty" export:"true"`
	Scheme    string `description:"Scheme used for the redirection." json:"scheme,omitempty" toml:"scheme,omitempty" yaml:"scheme,omitempty" export:"true"`
	Permanent bool   `description:"Applies a permanent redirection." json:"permanent,omitempty" toml:"permanent,omitempty" yaml:"permanent,omitempty" export:"true"`
	Priority  int    `description:"Priority of the generated router." json:"priority,omitempty" toml:"priority,omitempty" yaml:"priority,omitempty" export:"true"`
}

// SetDefaults sets the default values.
func (r *RedirectEntryPoint) SetDefaults() {
	r.Scheme = "https"
	r.Permanent = true
	r.Priority = math.MaxInt32 - 1
}

// TLSConfig is the default TLS configuration for all the routers associated to the concerned entry point.
type TLSConfig struct {
	Options      string         `description:"Default TLS options for the routers linked to the entry point." json:"options,omitempty" toml:"options,omitempty" yaml:"options,omitempty" export:"true"`
	CertResolver string         `description:"Default certificate resolver for the routers linked to the entry point." json:"certResolver,omitempty" toml:"certResolver,omitempty" yaml:"certResolver,omitempty" export:"true"`
	Domains      []types.Domain `description:"Default TLS domains for the routers linked to the entry point." json:"domains,omitempty" toml:"domains,omitempty" yaml:"domains,omitempty" export:"true"`
}

// ForwardedHeaders Trust client forwarding headers.
type ForwardedHeaders struct {
	Insecure   bool     `description:"Trust all forwarded headers." json:"insecure,omitempty" toml:"insecure,omitempty" yaml:"insecure,omitempty" export:"true"`
	TrustedIPs []string `description:"Trust only forwarded headers from selected IPs." json:"trustedIPs,omitempty" toml:"trustedIPs,omitempty" yaml:"trustedIPs,omitempty"`
}

// ProxyProtocol contains Proxy-Protocol configuration.
type ProxyProtocol struct {
	Insecure   bool     `description:"Trust all." json:"insecure,omitempty" toml:"insecure,omitempty" yaml:"insecure,omitempty" export:"true"`
	TrustedIPs []string `description:"Trust only selected IPs." json:"trustedIPs,omitempty" toml:"trustedIPs,omitempty" yaml:"trustedIPs,omitempty"`
}

// EntryPoints holds the HTTP entry point list.
type EntryPoints map[string]*EntryPoint
