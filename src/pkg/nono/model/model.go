package model

import (
	"redis"
	"strings"
)

type Model struct {
	db     int
	client redis.Client
	name   string
}

const keySeparator = ":"

func NewModel(name string, db int) *Model {
	s := redis.DefaultSpec().Db(db)
	m := new(Model)
	m.db = db
	m.name = name
	m.client, _ = redis.NewSynchClientWithSpec(s)
	// TODO
	// if e != nil { log.Stderr ("failed to create the client", e); return "failed" }
	return m
}

func (m *Model) FullKey(key string) string {
	return m.name + keySeparator + key
}

func (m *Model) Get(key string) string {
	value, _ := m.client.Get(m.FullKey(key))
	// TODO
	// if e!= nil { log.Stderr ("error on Get", e); return "failed 2" }
	return string(value);
}

func (m *Model) Set(key string, value string) {
	m.client.Set(m.FullKey(key), strings.Bytes(value))
}

