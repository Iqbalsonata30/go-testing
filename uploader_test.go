package gotesting

import (
	"context"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGS struct {
	mock.Mock
}

func (m *MockGS) NewWriter(ctx context.Context) *storage.Writer {
	args := m.Called(ctx)
	return args.Get(0).(*storage.Writer)
}

func TestUploader_Upload(t *testing.T) {
	m := new(MockGS)
	u := Uploader{
		svc: m,
	}
	m.On("NewWriter", mock.Anything).Return(&storage.Writer{})
	err := u.Upload()
	assert.NoError(t, err)
	m.AssertExpectations(t)
}
