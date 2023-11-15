package category

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type I18nDetail struct {
	Locale string
	Slug   string
}

type Repository interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	Update(ctx context.Context, entity *Entity) *i18np.Error
	Delete(ctx context.Context, categoryUUID string) *i18np.Error
	Disable(ctx context.Context, categoryUUID string) *i18np.Error
	Enable(ctx context.Context, categoryUUID string) *i18np.Error
	FindChild(ctx context.Context, categoryUUID string) ([]*Entity, *i18np.Error)
	Find(ctx context.Context, categoryUUID string) (*Entity, *i18np.Error)
	FindAll(ctx context.Context) ([]*Entity, *i18np.Error)
	AdminFindAll(ctx context.Context, onlyMains bool) ([]*Entity, *i18np.Error)
	AdminFindChild(ctx context.Context, categoryUUID string) ([]*Entity, *i18np.Error)
	UpdateOrder(ctx context.Context, categoryUUID string, order int16) *i18np.Error
	FindBySlug(ctx context.Context, i18n I18nDetail) (*Entity, *i18np.Error)
	FindFieldsByUUIDs(ctx context.Context, categoryUUIDs []string) ([]*Entity, *i18np.Error)
	FindByUUIDs(ctx context.Context, categoryUUIDs []string) ([]*Entity, *i18np.Error)
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepo(collection *mongo.Collection, factory Factory) Repository {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func validate(collection *mongo.Collection) {
	if collection == nil {
		panic("collection is nil")
	}
}

func (r *repo) Create(ctx context.Context, e *Entity) (*Entity, *i18np.Error) {
	e.BeforeCreate()
	res, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	e.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}

func (r *repo) Update(ctx context.Context, e *Entity) *i18np.Error {
	e.BeforeUpdate()
	id, err := mongo2.TransformId(e.UUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID("update")
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.MainUUIDs:   e.MainUUIDs,
			fields.Images:      e.Images,
			fields.Meta:        e.Meta,
			fields.InputGroups: e.InputGroups,
			fields.Inputs:      e.Inputs,
			fields.Rules:       e.Rules,
			fields.Alerts:      e.Alerts,
			fields.Validators:  e.Validators,
			fields.UpdatedAt:   time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Delete(ctx context.Context, categoryUUID string) *i18np.Error {
	id, err := mongo2.TransformId(categoryUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID("disable")
	}
	filter := bson.M{
		fields.UUID: id,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsDeleted: true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Disable(ctx context.Context, categoryUUID string) *i18np.Error {
	id, err := mongo2.TransformId(categoryUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID("disable")
	}
	filter := bson.M{
		fields.UUID:     id,
		fields.IsActive: true,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive:  false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) UpdateOrder(ctx context.Context, categoryUUID string, order int16) *i18np.Error {
	id, err := mongo2.TransformId(categoryUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID("re-order")
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.Order:     order,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Enable(ctx context.Context, categoryUUID string) *i18np.Error {
	id, err := mongo2.TransformId(categoryUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID("enable")
	}
	filter := bson.M{
		fields.UUID:     id,
		fields.IsActive: false,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive:  true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) FindChild(ctx context.Context, categoryUUID string) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		fields.MainUUIDs: bson.M{
			"$in": []string{categoryUUID},
		},
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
		fields.IsActive: true,
	}
	return r.helper.GetListFilter(ctx, filter, r.listOptions())
}

func (r *repo) AdminFindChild(ctx context.Context, categoryUUID string) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		fields.MainUUID: categoryUUID,
	}
	return r.helper.GetListFilter(ctx, filter, r.adminListOptions())
}

func (r *repo) Find(ctx context.Context, categoryUUID string) (*Entity, *i18np.Error) {
	id, _err := mongo2.TransformId(categoryUUID)
	if _err != nil {
		return nil, r.factory.Errors.InvalidUUID("disable")
	}
	filter := bson.M{
		fields.UUID: id,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) FindBySlug(ctx context.Context, i18n I18nDetail) (*Entity, *i18np.Error) {
	filter := bson.M{
		metaField(i18n.Locale, metaFields.Slug): i18n.Slug,
		fields.IsActive:                         true,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	e, exist, err := r.helper.GetFilter(ctx, filter, r.viewOptions())
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) FindFieldsByUUIDs(ctx context.Context, categoryUUIDs []string) ([]*Entity, *i18np.Error) {
	ids, err := mongo2.TransformIds(categoryUUIDs)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID("find by uuids")
	}
	filter := bson.M{
		fields.UUID: bson.M{
			"$in": ids,
		},
		fields.IsActive: true,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	return r.helper.GetListFilter(ctx, filter, r.fieldOptions())
}

func (r *repo) FindAll(ctx context.Context) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		fields.MainUUIDs: bson.M{
			"$size": 0,
		},
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
		fields.IsActive: true,
	}
	return r.helper.GetListFilter(ctx, filter, r.listOptions())
}

func (r *repo) AdminFindAll(ctx context.Context, onlyMains bool) ([]*Entity, *i18np.Error) {
	filter := bson.M{}
	if onlyMains {
		filter[fields.MainUUIDs] = bson.M{
			"$size": 0,
		}
	}
	return r.helper.GetListFilter(ctx, filter, r.adminListOptions())
}

func (r *repo) FindByUUIDs(ctx context.Context, categoryUUIDs []string) ([]*Entity, *i18np.Error) {
	ids, err := mongo2.TransformIds(categoryUUIDs)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID("find by uuids")
	}
	filter := bson.M{
		fields.UUID: bson.M{
			"$in": ids,
		},
		fields.IsActive: true,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	return r.helper.GetListFilter(ctx, filter, r.listOptions())
}

func (r *repo) AdminFindByUUIDs(ctx context.Context, categoryUUIDs []string) ([]*Entity, *i18np.Error) {
	ids, err := mongo2.TransformIds(categoryUUIDs)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID("find by uuids")
	}
	filter := bson.M{
		fields.UUID: bson.M{
			"$in": ids,
		},
	}
	return r.helper.GetListFilter(ctx, filter, r.adminListOptions())
}

func (r *repo) adminListOptions() *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:      1,
		fields.MainUUIDs: 1,
		fields.Images:    1,
		fields.Meta:      1,
		fields.IsActive:  1,
		fields.IsDeleted: 1,
		fields.UpdatedAt: 1,
	})
	return opts
}

func (r *repo) fieldOptions() *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:        1,
		fields.Inputs:      1,
		fields.InputGroups: 1,
		fields.Meta:        1,
		fields.Alerts:      1,
		fields.Rules:       1,
	})
	opts.SetSort(bson.M{
		fields.Order: 1,
	})
	return opts
}

func (r *repo) listOptions() *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:      1,
		fields.MainUUIDs: 1,
		fields.Images:    1,
		fields.Meta:      1,
	})
	opts.SetSort(bson.M{
		fields.Order: 1,
	})
	return opts
}

func (r *repo) viewOptions() *options.FindOneOptions {
	opts := &options.FindOneOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:      1,
		fields.MainUUIDs: 1,
		fields.Images:    1,
		fields.Meta:      1,
		fields.CreatedAt: 1,
	})
	return opts
}
