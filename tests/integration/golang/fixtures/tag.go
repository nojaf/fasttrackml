package fixtures

import (
	"context"
	"time"

	"github.com/rotisserie/eris"

	"github.com/G-Research/fasttrackml/pkg/api/mlflow/dao/models"
	"github.com/G-Research/fasttrackml/pkg/database"
)

// TagFixtures represents data fixtures object.
type TagFixtures struct {
	baseFixtures
}

// NewTagFixtures creates new instance of TagFixtures.
func NewTagFixtures(databaseDSN string) (*TagFixtures, error) {
	db, err := database.ConnectDB(
		databaseDSN,
		1*time.Second,
		20,
		false,
		false,
		"",
	)
	if err != nil {
		return nil, eris.Wrap(err, "error connection to database")
	}
	return &TagFixtures{
		baseFixtures: baseFixtures{db: db.DB},
	}, nil
}

// CreateTag creates new test Tag.
func (f TagFixtures) CreateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	if err := f.baseFixtures.db.WithContext(ctx).Create(tag).Error; err != nil {
		return nil, eris.Wrap(err, "error creating test tag")
	}
	return tag, nil
}