package commons

type MyConfig struct {
	Current HostSpecification `json:"current"`
}
type HostSpecification struct {
	Endpoint string `json:"endpoint,omitempty"`
	Target   string `json:"target,omitempty"`
}
type Clusters map[string]HostSpecification

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
