package service

import (
	"sync"

	"github.com/kataras/iris/core/errors"
)

type IdGeneratorService interface {
	GetId(serverId string) (int, error)
}

var server1Id = -1
var server2Id = -1
var mutex = sync.Mutex{}
var invalidServer = errors.New("Invalid serverId")

type DefaultIdGenerator struct {
}

func (g DefaultIdGenerator) GetId(serverId string) (id int, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	if serverId == "server1" {
		server1Id++
		id = server1Id
	} else if serverId == "server2" {
		server2Id++
		id = server2Id
	} else {
		err = invalidServer
	}

	return
}
