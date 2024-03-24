package audit

import (
	"errors"
	"time"
)

const (
	ENTITY_USER = "USER"
	ENTITY_BOOK = "BOOK"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

var (
	entities = map[string]Entities{
		ENTITY_USER: Entities_USER,
		ENTITY_BOOK: Entities_BOOK,
	}

	actions = map[string]Actions{
		ACTION_CREATE:   Actions_CREATE,
		ACTION_UPDATE:   Actions_UPDATE,
		ACTION_GET:      Actions_GET,
		ACTION_DELETE:   Actions_DELETE,
		ACTION_REGISTER: Actions_REGISTER,
		ACTION_LOGIN:    Actions_LOGIN,
	}
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbAction(action string) (Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, errors.New("invalid action")
	}

	return val, nil
}
