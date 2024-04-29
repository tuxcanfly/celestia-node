package da

import (
	"context"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/rollkit/go-da"
)

//go:generate mockgen -destination=mocks/api.go -package=mocks . Module
type Module interface {
	da.DA
	RegisterErrors(*jsonrpc.Errors)
}

// API is a wrapper around Module for the RPC.
// TODO(@distractedm1nd): These structs need to be autogenerated.
type API struct {
	Internal struct {
		MaxBlobSize func(ctx context.Context) (uint64, error)                                            `perm:"read"`
		Get         func(ctx context.Context, ids []da.ID, ns da.Namespace) ([]da.Blob, error)           `perm:"read"`
		GetIDs      func(ctx context.Context, height uint64, ns da.Namespace) ([]da.ID, error)           `perm:"read"`
		GetProofs   func(ctx context.Context, ids []da.ID, ns da.Namespace) ([]da.Proof, error)          `perm:"read"`
		Commit      func(ctx context.Context, blobs []da.Blob, ns da.Namespace) ([]da.Commitment, error) `perm:"read"`
		Validate    func(context.Context, []da.ID, []da.Proof, da.Namespace) ([]bool, error)             `perm:"read"`
		Submit      func(context.Context, []da.Blob, float64, da.Namespace) ([]da.ID, error)             `perm:"write"`
	}
}

func (api *API) MaxBlobSize(ctx context.Context) (uint64, error) {
	return api.Internal.MaxBlobSize(ctx)
}

func (api *API) Get(ctx context.Context, ids []da.ID, ns da.Namespace) ([]da.Blob, error) {
	return api.Internal.Get(ctx, ids, ns)
}

func (api *API) GetIDs(ctx context.Context, height uint64, ns da.Namespace) ([]da.ID, error) {
	return api.Internal.GetIDs(ctx, height, ns)
}

func (api *API) GetProofs(ctx context.Context, ids []da.ID, ns da.Namespace) ([]da.Proof, error) {
	return api.Internal.GetProofs(ctx, ids, ns)
}

func (api *API) Commit(ctx context.Context, blobs []da.Blob, ns da.Namespace) ([]da.Commitment, error) {
	return api.Internal.Commit(ctx, blobs, ns)
}

func (api *API) Validate(ctx context.Context, ids []da.ID, proofs []da.Proof, ns da.Namespace) ([]bool, error) {
	return api.Internal.Validate(ctx, ids, proofs, ns)
}

func (api *API) Submit(ctx context.Context, blobs []da.Blob, gasPrice float64, ns da.Namespace) ([]da.ID, error) {
	return api.Internal.Submit(ctx, blobs, gasPrice, ns)
}
