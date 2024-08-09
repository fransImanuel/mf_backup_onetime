package tr_tasklist

import (
	"context"
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/tr_tasklist/model"
	"mf_backup_onetime/schemas"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	BulkExportOneTimeService() []string
	ExportPDFTasklist(ctx context.Context, input dto.GetTasklistByID) (*string, string, schemas.SchemaDatabaseError)
}

type Repository interface {
	BulkMongoExportOneTime() []string
	GetTasklistByIdRepository(id primitive.ObjectID) (model.TRTasklist, schemas.SchemaDatabaseError)
}
