package fixture

const (
	// YFinQuotes are the yfin quote responses.
	YFinQuotes ResourceID = "quote"
	// YFinChart are the yfin chart responses.
	YFinChart ResourceID = "chart"
	// YFinOptions are the yfin options responses.
	YFinOptions ResourceID = "options"
	// YFinCrumb are the yfin crumb responses.
	YFinCrumb CrumbID = "crumb"
	// ServiceYFin is the yfin service.
	ServiceYFin ServiceID = "yfin"
)

// Path is a url path.
type Path string

// ResourceID is just an identifier for a resource.
type ResourceID string

// ServiceID is just an identifier for a service.
type ServiceID string

// CrumbID is just an identifier for a service.
type CrumbID string

// Resources alias for resource map.
type Resources map[ResourceID]interface{}

// Fixtures is a collection of resources.
type Fixtures struct {
	Resources map[ServiceID]Resources `json:"resources"`
}

// Spec specification of services.
type Spec struct {
	Services map[ServiceID]*Service `yaml:"services"`
}

// Service is a collection of url paths and resources.
type Service struct {
	Paths map[Path]*Operation `yaml:"paths"`
}

// Operation defines a service operation.
type Operation struct {
	Parameters []*Parameter `yaml:"parameters"`
	ResourceID ResourceID   `yaml:"resource"`
}

// Parameter describes a url parameter.
type Parameter struct {
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
	Required    bool   `yaml:"required"`
}
