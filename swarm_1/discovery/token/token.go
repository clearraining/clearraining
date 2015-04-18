package token

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/docker/swarm_1/discovery"
)

const DiscoveryURL = "http://discovery-stage.hub.docker.com/v1"

type Discovery struct 
	heartbeat uint64
	url       string
	token     string
}

func init() {
	discovery.Register("token", &Discovery{})
}

func (s *Discovery) Initialize(urltoken string, heartbeat uint64) error { //初始化url和token，以及heartbeat
	if i := strings.LastIndex(urltoken, "/"); i != -1 {
		s.url = "https://" + urltoken[:i]
		s.token = urltoken[i+1:]
	} else {
		s.url = DiscoveryURL
		s.token = urltoken
	}

	if s.token == "" {
		return errors.New("token is empty")
	}
	s.heartbeat = heartbeat

	return nil
}

func (s *Discovery) Register(addr string) error {
	buf := strings.NewReader(addr) //创建一个从addr读取数据的*reader

	resp, err := http.Post(fmt.Sprintf("%s/%s/%s", s.url, "clusters", s.token), "application", buf)

	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

func (s *Discovery) CreateCluster() (string, error) {
	resp, err := http.Post(fmt.Sprintf("%s/%s", s.url, "clusters"), "", nil) //返回的resp.body中包含token
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	token, err := ioutil.ReadAll(resp.Body)
	return string(token), err

}
