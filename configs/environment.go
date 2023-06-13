package configs

// Environment is a type for application environment
type Environment string

const (
	// Testing is the environment for testing
	Testing Environment = "testing"
	// Production is the environment for production
	Production Environment = "production"
	// Staging is the environment for staging
	Staging Environment = "staging"
	// Local is the environment for local
	Local Environment = "local"
)

// IsTesting returns true if the environment is testing
func (e Environment) IsTesting() bool {
	return e == Testing
}

// IsProduction returns true if the environment is production
func (e Environment) IsProduction() bool {
	return e == Production
}

// IsStaging returns true if the environment is staging
func (e Environment) IsStaging() bool {
	return e == Staging
}

// IsLocal returns true if the environment is local
func (e Environment) IsLocal() bool {
	return e == Local
}
