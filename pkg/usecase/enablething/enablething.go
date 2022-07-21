package enablething

import (
	"github.com/thisdougb/cleango/pkg/entity/thing"
)

// EnableThing set the status of a Thing
func (s *Service) EnableThing(thingID int) error {

	// an example of using an entity in a usecase
	thing := thing.NewThing(thingID)

	// this is mockable because of dependency injection
	err := s.repo.SetThingStatus(thing.Identifier, true)
	if err != nil {
		return err
	}

	// do some more stuff

	return nil
}
