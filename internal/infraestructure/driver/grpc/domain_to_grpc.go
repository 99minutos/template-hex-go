package grpc

/*
func TimeToGrpcTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
func DomainAccountToGrpc(ctx context.Context, o *domain.Account) *pb.Account {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Account/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Account{
			ApiKey:   o.ApiKey,
			Email:    o.Email,
			ClientId: o.ClientId,
			Company:  o.Company,
		}
	}
	return nil
}

func DomainActorToGrpc(ctx context.Context, o *domain.Actor) *pb.Actor {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Actor/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Actor{
			FirstName: o.FirstName,
			LastName:  o.LastName,
			Email:     o.Email,
			Phone:     o.Phone,
			Address:   o.Address,
		}
	}
	return nil
}
func DomainLocationToGrpc(ctx context.Context, o *domain.Location) *pb.Location {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Location/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Location{
			Lat:  o.Latitude,
			Lon:  o.Longitude,
			Type: o.GeoType,
		}
	}
	return nil
}
func DomainPickupToGrpc(ctx context.Context, o *domain.Pickup) *pb.Pickup {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Pickup/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Pickup{
			Actor:    DomainActorToGrpc(ctx, o.Actor),
			Attempts: int32(o.Attempts),
			Location: DomainLocationToGrpc(ctx, o.Location),
			PickupAt: TimeToGrpcTime(o.PickupAt),
		}
	}
	return nil
}
func DomainDeliveryToGrpc(ctx context.Context, o *domain.Delivery) *pb.Delivery {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Delivery/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Delivery{
			Actor:      DomainActorToGrpc(ctx, o.Actor),
			Attempts:   int32(o.Attempts),
			Location:   DomainLocationToGrpc(ctx, o.Location),
			DeliveryAt: TimeToGrpcTime(o.DeliveryAt),
		}
	}
	return nil
}
func DomainReturnToGrpc(ctx context.Context, o *domain.Return) *pb.Return {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Return/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Return{
			Actor:    DomainActorToGrpc(ctx, o.Actor),
			Attempts: int32(o.Attempts),
			Location: DomainLocationToGrpc(ctx, o.Location),
			ReturnAt: TimeToGrpcTime(o.ReturnAt),
		}
	}
	return nil
}
func DomainCodToGrpc(ctx context.Context, o *domain.Cod) *pb.Cod {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Cod/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Cod{
			Amount:     o.Amount,
			AmountPaid: o.AmountPaid,
			Reference:  o.Reference,
		}
	}
	return nil
}
func DomainStationToGrpc(ctx context.Context, o *domain.Station) *pb.Station {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Station/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Station{
			Previous: o.Previous,
			Current:  o.Current,
			Location: DomainLocationToGrpc(ctx, o.Location),
		}
	}
	return nil
}

func DomainProviderToGrpc(ctx context.Context, o *domain.Provider) *pb.Provider {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Provider/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Provider{
			StaffId:      o.StaffId,
			Name:         o.Name,
			Email:        o.Email,
			TradeName:    o.TradeName,
			Nomenclature: o.Nomenclature,
		}
	}
	return nil
}
func DomainDriverToGrpc(ctx context.Context, o *domain.Driver) *pb.Driver {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Driver/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Driver{
			DriverId: o.DriverId,
			StaffId:  o.StaffId,
			Name:     o.Name,
			Nickname: o.Nickname,
			Type:     pb.DriverResourceType(pb.DriverResourceType_value[string(o.ResourceType)]),
			Provider: DomainProviderToGrpc(ctx, o.Provider),
		}
	}
	return nil
}
func DomainDimensionsToGrpc(ctx context.Context, o *domain.Dimensions) *pb.Dimensions {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Dimensions/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Dimensions{
			Length:      o.Length,
			Width:       o.Width,
			Height:      o.Height,
			Weight:      o.Weight,
			PackageSize: o.PackageSize,
			ExecutedAt:  TimeToGrpcTime(o.ExecutedAt),
		}
	}
	return nil
}
func DomainMetadataToGrpc(ctx context.Context, o *domain.Metadata) *pb.Metadata {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Metadata/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.Metadata{
			IsPriority:    o.IsPriority,
			IsReturn:      o.IsReturn,
			IsFulfillment: o.IsFulfillment,
			IsSelfService: o.IsSelfService,
			IsIntegration: o.IsIntegration,
		}
	}
	return nil
}
func DomainLegacyMetadataToGrpc(ctx context.Context, o *domain.LegacyMetadata) *pb.LegacyMetadata {
	ctx, span := tracer.Tracer.Start(ctx, "Order/LegacyMetadata/ToGrpc")
	defer span.End()
	if o != nil {
		return &pb.LegacyMetadata{
			Status:      o.Status,
			Description: o.Description,
		}
	}
	return nil
}

func DomainOrderToGrpc(ctx context.Context, o *domain.Order) (*pb.Order, error) {
	ctx, span := tracer.Tracer.Start(ctx, "Order/Order/ToGrpc")
	defer span.End()

	order := &pb.Order{
		Id:             o.Id,
		TrackingId:     o.TrackingId,
		InternalKey:    o.InternalKey,
		DeliveryType:   o.DeliveryType,
		Status:         int32(o.Status),
		StatusName:     o.StatusName,
		Country:        o.Country,
		Account:        DomainAccountToGrpc(ctx, o.Account),
		Pickup:         DomainPickupToGrpc(ctx, o.Pickup),
		Delivery:       DomainDeliveryToGrpc(ctx, o.Delivery),
		Return:         DomainReturnToGrpc(ctx, o.Return),
		Cod:            DomainCodToGrpc(ctx, o.Cod),
		Station:        DomainStationToGrpc(ctx, o.Station),
		Driver:         DomainDriverToGrpc(ctx, o.Driver),
		Dimensions:     DomainDimensionsToGrpc(ctx, o.Dimensions),
		Metadata:       DomainMetadataToGrpc(ctx, o.Metadata),
		LegacyMetadata: DomainLegacyMetadataToGrpc(ctx, o.LegacyMetadata),
	}

	if o.CreatedSs != nil {
		order.CreatedSs = timestamppb.New(*o.CreatedSs)
	}
	if o.UpdatedSs != nil {
		order.UpdatedSs = timestamppb.New(*o.UpdatedSs)
	}
	if o.CreatedAt != nil {
		order.CreatedAt = timestamppb.New(*o.CreatedAt)
	}
	if o.UpdatedAt != nil {
		order.UpdatedAt = timestamppb.New(*o.UpdatedAt)
	}
	if o.CanceledAt != nil {
		order.CanceledAt = timestamppb.New(*o.CanceledAt)
	}

	return order, nil
}
*/
