package keys

import (
	"context"
	"fmt"
	"testing"

	mock3 "github.com/consensys/quorum-key-manager/src/auth/mock"
	"github.com/consensys/quorum-key-manager/src/auth/types"

	"github.com/consensys/quorum-key-manager/src/infra/log/testutils"
	mock2 "github.com/consensys/quorum-key-manager/src/stores/database/mock"
	"github.com/consensys/quorum-key-manager/src/stores/entities"
	testutils2 "github.com/consensys/quorum-key-manager/src/stores/entities/testutils"
	"github.com/consensys/quorum-key-manager/src/stores/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListKey(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedErr := fmt.Errorf("error")

	store := mock.NewMockKeyStore(ctrl)
	db := mock2.NewMockKeys(ctrl)
	logger := testutils.NewMockLogger(ctrl)
	auth := mock3.NewMockAuthorizator(ctrl)

	connector := NewConnector(store, db, auth, logger)

	t.Run("should list keys successfully", func(t *testing.T) {
		keyOne := testutils2.FakeKey()
		keyTwo := testutils2.FakeKey()

		auth.EXPECT().CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceKey}).Return(nil)
		db.EXPECT().GetAll(gomock.Any()).Return([]*entities.Key{keyOne, keyTwo}, nil)

		keyIDs, err := connector.List(ctx)

		assert.NoError(t, err)
		assert.Equal(t, keyIDs, []string{keyOne.ID, keyTwo.ID})
	})

	t.Run("should fail with same error if authorization fails", func(t *testing.T) {
		auth.EXPECT().CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceKey}).Return(expectedErr)

		_, err := connector.List(ctx)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("should fail to list keys if db fails", func(t *testing.T) {
		auth.EXPECT().CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceKey}).Return(nil)
		db.EXPECT().GetAll(gomock.Any()).Return(nil, expectedErr)

		_, err := connector.List(ctx)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})
}

func TestListDeletedKey(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedErr := fmt.Errorf("error")

	store := mock.NewMockKeyStore(ctrl)
	db := mock2.NewMockKeys(ctrl)
	logger := testutils.NewMockLogger(ctrl)
	auth := mock3.NewMockAuthorizator(ctrl)

	connector := NewConnector(store, db, auth, logger)

	t.Run("should list deleted key successfully", func(t *testing.T) {
		keyOne := testutils2.FakeKey()
		keyTwo := testutils2.FakeKey()

		auth.EXPECT().CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceKey}).Return(nil)
		db.EXPECT().GetAllDeleted(gomock.Any()).Return([]*entities.Key{keyOne, keyTwo}, nil)

		keyIDs, err := connector.ListDeleted(ctx)

		assert.NoError(t, err)
		assert.Equal(t, keyIDs, []string{keyOne.ID, keyTwo.ID})
	})

	t.Run("should fail with same error if authorization fails", func(t *testing.T) {
		auth.EXPECT().CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceKey}).Return(expectedErr)

		_, err := connector.ListDeleted(ctx)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("should fail to list deleted key if db fails", func(t *testing.T) {
		auth.EXPECT().CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceKey}).Return(nil)
		db.EXPECT().GetAllDeleted(gomock.Any()).Return(nil, expectedErr)

		_, err := connector.ListDeleted(ctx)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})
}