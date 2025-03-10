package ports

import (
	"context"
	stdockpb "github.com/zandala88/gorder/internal/common/genproto/stockpb"
	"github.com/zandala88/gorder/stock/app"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stdockpb.GetItemsRequest) (*stdockpb.GetItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stdockpb.CheckIfItemsInStockRequest) (*stdockpb.CheckIfItemsInStockResponse, error) {
	//TODO implement me
	panic("implement me")
}
