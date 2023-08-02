package grpc

/*
type SotGrpcServer struct {
	server     *grpc.Server
	SOTService ports.ISOTService
}

func NewSotGrpcServer(sotService ports.ISOTService) *SotGrpcServer {
	grpcServer := grpc.NewServer(
		grpc.MaxSendMsgSize(1024*1024*10),
		grpc.MaxRecvMsgSize(1024*1024*10),
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	sotgrpc.RegisterSourceOfTruthServer(grpcServer, NewSotGrpcActions(sotService))

	reflection.Register(grpcServer)

	return &SotGrpcServer{
		server:     grpcServer,
		SOTService: sotService,
	}
}

func (s *SotGrpcServer) Start(is net.Listener) {
	err := s.server.Serve(is)
	if err != nil {
		logger.Logger.Fatal(err)
	}
}
*/
