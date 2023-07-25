package protocols

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	"github.com/codeclout/AccountEd/notifications/ports/framework/drivers"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type Adapter struct {
	config   map[string]interface{}
	monitor  monitoring.Adapter
	protocol drivers.EmailDriverPort
}

func NewAdapter(config map[string]interface{}, protocol drivers.EmailDriverPort, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:   config,
		monitor:  monitor,
		protocol: protocol,
	}
}

func (a *Adapter) Run() {
	var options []grpc.ServerOption

	listener, e := net.Listen("tcp", a.getPort())
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		os.Exit(1)
	}

	server := grpc.NewServer(options...)
	pb.RegisterEmailNotificationServiceServer(server, a.protocol)
	reflection.Register(server)

	if e := server.Serve(listener); e != nil {
		a.monitor.LogGenericError(e.Error())
		os.Exit(1)
	}
}

func (a *Adapter) getPort() string {
	p, ok := a.config["Port"].(string)
	n, _ := strconv.Atoi(p)

	if ok && len(strings.TrimSpace(p)) >= 4 && n >= 1024 && n <= 65535 {
		return fmt.Sprintf(":%d", n)
	}

	return ":8088"
}
