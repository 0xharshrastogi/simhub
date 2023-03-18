package common

type Encyclopedia interface {
	// fetches information for a resource
	GetResourceById(id string) (*Resource, error)

	// fetches information for a building
	GetBuildingById(id string) (*Building, error)
}
