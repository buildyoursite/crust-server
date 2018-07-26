package rest

import (
	"github.com/pkg/errors"

	"github.com/titpetric/factory/resputil"

	"context"
	"github.com/crusttech/crust/crm/rest/server"
	"github.com/crusttech/crust/crm/types"
)

var _ = errors.Wrap

type (
	Module struct {
		svc moduleService
	}

	moduleService interface {
		FindByID(context.Context, uint64) (*types.Module, error)
		Find(context.Context) ([]*types.Module, error)

		Create(context.Context, *types.Module) (*types.Module, error)
		Update(context.Context, *types.Module) (*types.Module, error)
		DeleteById(context.Context, uint64) error
	}
)

func (Module) New(moduleSvc moduleService) *Module {
	var ctrl = &Module{}
	ctrl.svc = moduleSvc
	return ctrl
}

func (c *Module) List(ctx context.Context, r *server.ModuleListRequest) (interface{}, error) {
	return c.svc.Find(ctx)
}

func (c *Module) Read(ctx context.Context, r *server.ModuleReadRequest) (interface{}, error) {
	return c.svc.FindByID(ctx, r.ID)
}

func (c *Module) Delete(ctx context.Context, r *server.ModuleDeleteRequest) (interface{}, error) {
	return resputil.OK(), c.svc.DeleteById(ctx, r.ID)
}

func (c *Module) Create(ctx context.Context, r *server.ModuleCreateRequest) (interface{}, error) {
	m := types.Module{}.New()
	m.SetName(r.Name)
	return c.svc.Create(ctx, m)
}

func (c *Module) Edit(ctx context.Context, r *server.ModuleEditRequest) (interface{}, error) {
	m := types.Module{}.New()
	m.SetID(r.ID).SetName(r.Name)
	return c.svc.Update(ctx, m)
}

func (*Module) ContentList(ctx context.Context, r *server.ModuleContentListRequest) (interface{}, error) {
	return nil, errors.New("Not implemented: Module.content/edit")
}

func (*Module) ContentRead(ctx context.Context, r *server.ModuleContentReadRequest) (interface{}, error) {
	return nil, errors.New("Not implemented: Module.content/edit")
}

func (*Module) ContentCreate(ctx context.Context, r *server.ModuleContentCreateRequest) (interface{}, error) {
	return nil, errors.New("Not implemented: Module.content/edit")
}

func (*Module) ContentEdit(ctx context.Context, r *server.ModuleContentEditRequest) (interface{}, error) {
	return nil, errors.New("Not implemented: Module.content/edit")
}

func (*Module) ContentDelete(ctx context.Context, r *server.ModuleContentDeleteRequest) (interface{}, error) {
	return nil, errors.New("Not implemented: Module.content/delete")
}
