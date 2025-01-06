package minio

import "go.uber.org/fx"

var Module = fx.Provide(
	NewMinIO,
)
