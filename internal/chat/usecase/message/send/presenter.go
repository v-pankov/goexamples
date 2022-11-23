package send

import (
	"context"
)

type Presenter interface {
	Present(context.Context, *ResponseModel)
}

func NewPresenter(
	viewer Viewer,
) Presenter {
	return presenter{
		viewer: viewer,
	}
}

type presenter struct {
	viewer Viewer
}

func (p presenter) Present(ctx context.Context, rsp *ResponseModel) {
	p.viewer.View(
		ctx,
		&ViewModel{
			MessageID:       rsp.MessageID,
			MessageContents: rsp.MessageContents,
			CreatedAt:       rsp.CreatedAt,
		},
	)
}
