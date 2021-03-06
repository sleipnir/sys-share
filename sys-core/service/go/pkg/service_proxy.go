package pkg

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// Server side
type SysEmailProxyService struct {
	SysMail *mailProxyService
}

func NewSysMailProxyService(e EmailService) *SysEmailProxyService {
	return &SysEmailProxyService{SysMail: newMailProxyService(e)}
}

func (s *SysEmailProxyService) RegisterSvc(server *grpc.Server) {
	s.SysMail.registerSvc(server)
}

type SysBusProxyService struct {
	SysBus *busProxyService
}

func NewSysBusProxyService(bs BusService) *SysBusProxyService {
	return &SysBusProxyService{newBusServiceProxy(bs)}
}

func (s *SysBusProxyService) RegisterSvc(server *grpc.Server) {
	s.SysBus.registerSvc(server)
}

type SysCoreProxyService struct {
	SysCore *sysCoreService
}

func NewSysCoreProxyService(ds DbAdminService) *SysCoreProxyService {
	return &SysCoreProxyService{SysCore: newSysCoreService(ds)}
}

func (s *SysCoreProxyService) RegisterSvc(server *grpc.Server) {
	s.SysCore.registerSvc(server)
}

// Client side
type SysMailProxyServiceClient struct {
	*emailClientProxy
}

func NewSysMailProxyServiceClient(cc grpc.ClientConnInterface) *SysMailProxyServiceClient {
	return &SysMailProxyServiceClient{newEmailClientProxy(cc)}
}

type SysBusProxyServiceClient struct {
	*busClientProxy
}

func NewSysBusProxyServiceClient(cc grpc.ClientConnInterface) *SysBusProxyServiceClient {
	return &SysBusProxyServiceClient{newBusClientProxy(cc)}
}

type SysCoreProxyServiceClient struct {
	*sysCoreClientProxy
}

func NewSysCoreProxyServiceClient(cc grpc.ClientConnInterface) *SysCoreProxyServiceClient {
	return &SysCoreProxyServiceClient{newSysCoreClientProxy(cc)}
}

// CLI
type SysCoreProxyClient struct {
	SysCoreClient *sysCoreClient
}

func NewSysCoreProxyClient() *SysCoreProxyClient {
	return &SysCoreProxyClient{SysCoreClient: newSysCoreClient()}
}

func (s *SysCoreProxyClient) CobraCommand() *cobra.Command {
	return s.SysCoreClient.cobraCommand()
}

type SysBusProxyClient struct {
	SysBusClient *busClient
}

func NewSysBusProxyClient() *SysBusProxyClient {
	return &SysBusProxyClient{SysBusClient: newBusClient()}
}

func (sb *SysBusProxyClient) CobraCommand() *cobra.Command {
	return sb.SysBusClient.cobraCommand()
}

type SysMailProxyClient struct {
	SysMailProxyClient *mailClient
}

func NewSysMailProxyClient() *SysMailProxyClient {
	return &SysMailProxyClient{
		SysMailProxyClient: newMailClient(),
	}
}

func (sm *SysMailProxyClient) CobraCommand() *cobra.Command {
	return sm.SysMailProxyClient.cobraCommand()
}
