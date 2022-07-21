package redis

import (
	"fmt"
)

func (d *Datastore) SetThingStatus(thingID int, status bool) error {

	key := fmt.Sprintf(thingStatusKeyFormat, thingID)

	var statusValue int

	if status {
		statusValue = 1
	} else {
		statusValue = 0
	}

	_, err := d.client.Set(d.ctx, key, statusValue, 0).Result()
	if err != nil {
		return err
	}

	return nil
}
