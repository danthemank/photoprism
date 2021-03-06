package context

import (
	"testing"

	"github.com/photoprism/photoprism/internal/fsutil"
	"github.com/photoprism/photoprism/internal/test"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	ctx := test.CliContext()

	assert.True(t, ctx.IsSet("assets-path"))
	assert.False(t, ctx.Bool("debug"))

	c := NewConfig(ctx)

	assert.IsType(t, new(Config), c)

	assert.Equal(t, fsutil.ExpandedFilename("../../assets"), c.GetAssetsPath())
	assert.False(t, c.IsDebug())
}

func TestConfig_SetValuesFromFile(t *testing.T) {
	c := NewConfig(test.CliContext())

	c.SetValuesFromFile(fsutil.ExpandedFilename("../../configs/photoprism.yml"))

	assert.Equal(t, "/srv/photoprism", c.GetAssetsPath())
	assert.Equal(t, "/srv/photoprism/cache", c.GetCachePath())
	assert.Equal(t, "/srv/photoprism/cache/thumbnails", c.GetThumbnailsPath())
	assert.Equal(t, "/srv/photoprism/photos/originals", c.GetOriginalsPath())
	assert.Equal(t, "/srv/photoprism/photos/import", c.GetImportPath())
	assert.Equal(t, "/srv/photoprism/photos/export", c.GetExportPath())
	assert.Equal(t, "mysql", c.GetDatabaseDriver())
	assert.Equal(t, "photoprism:photoprism@tcp(database:3306)/photoprism?parseTime=true", c.GetDatabaseDsn())
}
