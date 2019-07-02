package cache

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/bookun/sandbox/go/judge-clientip/entity"
)

type Firebase struct {
	client *firestore.Client
}

func NewFirebase(ctx context.Context, projectID string) (*Firebase, error) {
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &Firebase{client: client}, nil

}

func (f *Firebase) Add(ctx context.Context, ip string, result entity.Result) error {
	_, err := f.client.Collection("ocn-ip").Doc(ip).Set(ctx, result)
	return err
}

func (f *Firebase) Get(ctx context.Context, ip string) (entity.Result, error) {
	result := entity.Result{}
	cache, err := f.client.Collection("ocn-ip").Doc(ip).Get(ctx)
	if !cache.Exists() {
		return result, nil
	}
	if err != nil {
		return result, err
	}
	if err := cache.DataTo(&result); err != nil {
		return result, err
	}
	return result, nil
}
