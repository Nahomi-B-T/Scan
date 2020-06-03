package provider

// Provider abstracts a scan provider
type Provider interface {
	Version() (string, error)
	Scan(image string) error
}
