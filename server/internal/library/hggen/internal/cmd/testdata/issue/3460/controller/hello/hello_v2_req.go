package hello

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/internal/library/hggen/internal/cmd/testdata/issue/3460/api/hello/v2"
)

func (c *ControllerV2) DictTypeAddPage(ctx context.Context, req *v2.DictTypeAddPageReq) (res *v2.DictTypeAddPageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV2) DictTypeAdd(ctx context.Context, req *v2.DictTypeAddReq) (res *v2.DictTypeAddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV2) DictTypeEditPage(ctx context.Context, req *v2.DictTypeEditPageReq) (res *v2.DictTypeEditPageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV2) DictTypeEdit(ctx context.Context, req *v2.DictTypeEditReq) (res *v2.DictTypeEditRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
