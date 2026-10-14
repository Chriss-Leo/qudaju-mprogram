package repo

import (
	"context"
	"gorm.io/gorm"
)

type txKey struct{}
type BaseRepo struct{ Connection *gorm.DB }

func (r BaseRepo) DB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return r.Connection.WithContext(ctx)
}
func WithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}
