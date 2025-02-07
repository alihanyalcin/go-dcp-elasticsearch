package couchbase

import "time"

type Event struct {
	EventTime      time.Time
	CollectionName string
	Key            []byte
	Value          []byte
	Cas            uint64
	IsDeleted      bool
	IsExpired      bool
	IsMutated      bool
}

func NewDeleteEvent(key []byte, value []byte, collectionName string, cas uint64, eventTime time.Time) Event {
	return Event{
		Key:            key,
		Value:          value,
		IsDeleted:      true,
		CollectionName: collectionName,
		Cas:            cas,
		EventTime:      eventTime,
	}
}

func NewExpireEvent(key []byte, value []byte, collectionName string, cas uint64, eventTime time.Time) Event {
	return Event{
		Key:            key,
		Value:          value,
		IsExpired:      true,
		CollectionName: collectionName,
		Cas:            cas,
		EventTime:      eventTime,
	}
}

func NewMutateEvent(key []byte, value []byte, collectionName string, cas uint64, eventTime time.Time) Event {
	return Event{
		Key:            key,
		Value:          value,
		IsMutated:      true,
		CollectionName: collectionName,
		Cas:            cas,
		EventTime:      eventTime,
	}
}
