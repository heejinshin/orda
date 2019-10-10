package service

import (
	"context"
	"github.com/knowhunger/ortoo/commons/log"
	"github.com/knowhunger/ortoo/commons/model"
	"github.com/knowhunger/ortoo/server/mongodb/schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

//ProcessClient processes ClientRequest and returns ClientResponse
func (o *OrtooService) ProcessClient(ctx context.Context, in *model.ClientRequest) (*model.ClientResponse, error) {
	transferredDoc := schema.ClientModelToBson(in.Client)

	collectionDoc, err := o.mongo.GetCollection(ctx, transferredDoc.Collection)
	if err != nil {
		return nil, model.NewRPCError(model.RPCErrMongoDB)
	}
	if collectionDoc == nil {
		return nil, log.OrtooError(status.New(codes.InvalidArgument, "fail to find collection").Err())
	}

	storedDoc, err := o.mongo.GetClient(ctx, transferredDoc.CUID)
	if err != nil {
		return nil, log.OrtooErrorf(err, "fail to get client")
	}
	if storedDoc == nil {
		transferredDoc.CreatedAt = time.Now()
		log.Logger.Infof("A new client is created:%+v", transferredDoc)
		if _, err := o.mongo.GetOrCreateCollectionSnapshot(ctx, transferredDoc.Collection); err != nil {
			return nil, model.NewRPCError(model.RPCErrMongoDB)
		}
	} else {
		if storedDoc.Collection != transferredDoc.Collection {
			return nil, model.NewRPCError(model.RPCErrClientInconsistentCollection, storedDoc.Collection, transferredDoc.Collection)
		}
		log.Logger.Infof("Client will be updated:%+v", transferredDoc)
	}
	if err = o.mongo.UpdateClient(ctx, transferredDoc); err != nil {
		return nil, model.NewRPCError(model.RPCErrMongoDB)
	}

	return model.NewClientResponse(in.Header, model.StateOfResponse_OK), nil
}
