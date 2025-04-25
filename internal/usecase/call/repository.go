package call

import (
	"context"
	//entity "internal/entity"
)

type CallRepository interface {
	Save(ctx context.Context, c *entity.Call)
}
