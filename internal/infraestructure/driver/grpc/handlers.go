package grpc

/*
type SOTGrpc struct {
	sotgrpc.UnimplementedSourceOfTruthServer
	SOTService ports.ISOTService
}

func NewSotGrpcActions(sotService ports.ISOTService) sotgrpc.SourceOfTruthServer {
	return &SOTGrpc{
		SOTService: sotService,
	}
}
func (s *SOTGrpc) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	ctx, span := tracer.Tracer.Start(ctx, "GRPC/GetOrder")
	defer span.End()

	order, err := s.SOTService.GetOrder(ctx, in.TrackingId, in.SearchType.String())
	if err != nil {
		return nil, err
	}

	grpcOrder, err := DomainOrderToGrpc(ctx, order)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponse{
		Order: grpcOrder,
	}, nil
}
*/
