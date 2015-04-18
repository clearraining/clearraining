package discovery

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Entry struct {
	Host string
	Port string
}

type Discovery interface {
	Initialize(string, uint64) error
	Fetch() ([]*Entry, error)
	Watch(WatchCallback)
	Register(string) error
}

var (
	discoveries map[string]Discovery

	ErrSupported = errors.New("discovery service not supported")
)

func init() {
	discoveries = make(map[string]Discovery)
}

func Register(scheme string, d Discovery) error {
	if _, exists := discoveries[scheme]; exists {
		return fmt.Errorf("scheme already registered %s", scheme)
	}
	log.WithField("name", scheme).Debug("Registering discovery service")
	discoveries[scheme] = d

	return nil
}

func parse(rawurl string) (string, string) {
	parts := strings.Split(rawurl, "://", 2) //对rawurl进行切割，返回切片，2表明切片的数目

	if len(parts) == 1 {
		return "nodes", parts[0]
	}
	return parts[0], parts[1]

}

func New(rawurl string, heartbeat uint64) (Discovery, error) {
	scheme, uri := parse(rawurl)

	if discovery, exists := discoveries[scheme]; exists { //scheme 是token 
		log.WithFields(log.Fields{"name": scheme}, "uri": rui).Debug("Initializing discovery servie")
		err := discovery.Initialize(uri, heartbeat)
		return discovery, err
	}

	return nil, ErrSupported
}











