package protocols

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
	"github.com/codeclout/AccountEd/session/ports/framework/drivers/cloud"
	"github.com/codeclout/AccountEd/session/ports/framework/drivers/member"
)

type Adapter struct {
	config       map[string]interface{}
	log          *slog.Logger
	cloudDriver  cloud.AWSDriverPort
	memberDriver member.SessionDriverMemberPort
}

func NewAdapter(config map[string]interface{}, cloud cloud.AWSDriverPort, m member.SessionDriverMemberPort, log *slog.Logger, ) *Adapter {
	return &Adapter{
		config:       config,
		cloudDriver:  cloud,
		log:          log,
		memberDriver: m,
	}
}

func (a *Adapter) Run() {
	var options []grpc.ServerOption

	listener, e := net.Listen("tcp", a.getPort())
	if e != nil {
		a.log.Error(e.Error())
		os.Exit(1)
	}

	server := grpc.NewServer(options...)
	awspb.RegisterAWSResourceClientServiceServer(server, a.cloudDriver)
	pb.RegisterMemberSessionServer(server, a.memberDriver)
	reflection.Register(server)

	if e := server.Serve(listener); e != nil {
		a.log.Error(e.Error())
		os.Exit(1)
	}
}

func (a *Adapter) getPort() string {
	p, ok := a.config["Port"].(string)
	n, _ := strconv.Atoi(p)

	if ok && len(strings.TrimSpace(p)) >= 4 && n >= 1024 && n <= 65535 {
		return fmt.Sprintf(":%d", n)
	}

	return ":9001"
}
