package wsFramework

type Broadcast struct {
}

func (b *Broadcast) BroadcastToOthers(roomName string, clientID string, event string, data interface{}) error {
	room, err := rooms.getRoom(roomName)
	if err != nil {
		return err
	}

	for id, client := range room.Clients {
		if id == clientID {
			continue
		}

		err := client.Emit("", event, data)
		if err != nil {
			return err
		}
	}

	return nil
}
