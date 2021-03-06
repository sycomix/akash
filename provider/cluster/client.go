package cluster

import (
	"bufio"
	"context"
	"errors"
	"io"
	"sync"

	"github.com/ovrclk/akash/manifest"
	atypes "github.com/ovrclk/akash/types"
	"github.com/ovrclk/akash/types/unit"
	mquery "github.com/ovrclk/akash/x/market/query"
	mtypes "github.com/ovrclk/akash/x/market/types"
)

var ErrNoDeployments = errors.New("no deployments")

type Client interface {
	Deploy(mtypes.LeaseID, *manifest.Group) error
	TeardownLease(mtypes.LeaseID) error
	Deployments() ([]Deployment, error)
	LeaseStatus(mtypes.LeaseID) (*LeaseStatus, error)
	ServiceStatus(mtypes.LeaseID, string) (*ServiceStatus, error)
	ServiceLogs(context.Context, mtypes.LeaseID, int64, bool) ([]*ServiceLog, error)
	Inventory() ([]Node, error)
}

type Node interface {
	ID() string
	Available() atypes.Unit
}

type node struct {
	id        string
	available atypes.Unit
}

func NewNode(id string, available atypes.Unit) Node {
	return &node{id: id, available: available}
}

func (n *node) ID() string {
	return n.id
}

func (n *node) Available() atypes.Unit {
	return n.available
}

type Deployment interface {
	LeaseID() mtypes.LeaseID
	ManifestGroup() manifest.Group
}

type ServiceLog struct {
	Name    string
	Stream  io.ReadCloser
	Scanner *bufio.Scanner
}

const (
	// 5 CPUs, 5Gi memory for null client.
	nullClientCPU    = 5000
	nullClientMemory = 32 * unit.Gi
	nullClientDisk   = 512 * unit.Gi
)

type nullClient struct {
	leases map[string]*manifest.Group
	mtx    sync.Mutex
}

func NewServiceLog(name string, stream io.ReadCloser) *ServiceLog {
	return &ServiceLog{
		Name:    name,
		Stream:  stream,
		Scanner: bufio.NewScanner(stream),
	}
}

func NullClient() Client {
	return &nullClient{
		leases: make(map[string]*manifest.Group),
		mtx:    sync.Mutex{},
	}
}

func (c *nullClient) Deploy(lid mtypes.LeaseID, mgroup *manifest.Group) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.leases[mquery.LeasePath(lid)] = mgroup
	return nil
}

func (c *nullClient) LeaseStatus(lid mtypes.LeaseID) (*LeaseStatus, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	mgroup, ok := c.leases[mquery.LeasePath(lid)]
	if !ok {
		return nil, nil
	}

	resp := &LeaseStatus{}
	for _, svc := range mgroup.Services {
		resp.Services = append(resp.Services, &ServiceStatus{
			Name:      svc.Name,
			Available: int32(svc.Count),
			Total:     int32(svc.Count),
		})
	}

	return resp, nil
}

func (c *nullClient) ServiceStatus(_ mtypes.LeaseID, _ string) (*ServiceStatus, error) {
	return nil, nil
}

func (c *nullClient) ServiceLogs(_ context.Context, _ mtypes.LeaseID, _ int64, _ bool) ([]*ServiceLog, error) {
	return nil, nil
}

func (c *nullClient) TeardownLease(lid mtypes.LeaseID) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.leases, mquery.LeasePath(lid))
	return nil
}

func (c *nullClient) Deployments() ([]Deployment, error) {
	return nil, nil
}

func (c *nullClient) Inventory() ([]Node, error) {
	return []Node{
		NewNode("solo", atypes.Unit{
			CPU:     nullClientCPU,
			Memory:  nullClientMemory,
			Storage: nullClientDisk,
		}),
	}, nil
}
