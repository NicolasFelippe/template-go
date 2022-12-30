package user_service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"template-go/internal/core/domain"
	cryptomock "template-go/mocks/pkg/crypto"
	uidgenmock "template-go/mocks/pkg/uidgen"
	userrepositorymock "template-go/mocks/repositories"
	"template-go/pkg/util"
	"testing"
)

func randomUser(t *testing.T, crypt *cryptomock.MockCrypto) (user *domain.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := crypt.HashPassword(password)
	require.NoError(t, err)

	user = &domain.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}

func TestCreateUserSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCrypto := cryptomock.NewMockCrypto(ctrl)
	mockCrypto.EXPECT().HashPassword(gomock.Any()).Return("@hashedPassword", nil).AnyTimes()

	mockUidGen := uidgenmock.NewMockUIDGen(ctrl)
	mockUidGen.EXPECT().New().Return(uuid.UUID{})

	user, password := randomUser(t, mockCrypto)

	mockRepository := userrepositorymock.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().CreateUser(gomock.Any()).Return(user, nil).Times(1)

	userService := New(mockRepository, mockUidGen, mockCrypto)
	createUser, err := userService.CreateUser(user.Username, password, user.FullName, user.Email)

	require.NoError(t, err)
	require.NotEmpty(t, createUser)
	require.Equal(t, createUser.FullName, user.FullName)
	require.Equal(t, createUser.Username, user.Username)
	require.Equal(t, createUser.Email, user.Email)
	require.Equal(t, createUser.HashedPassword, user.HashedPassword)
}

func TestCreateUserHashedPasswordError(t *testing.T) {
	msgErrorHashedPassword := "failed to has crypto"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCrypto := cryptomock.NewMockCrypto(ctrl)
	mockCrypto.EXPECT().HashPassword(gomock.Any()).Return("", errors.New(msgErrorHashedPassword)).AnyTimes()

	mockUidGen := uidgenmock.NewMockUIDGen(ctrl)
	mockUidGen.EXPECT().New().Return(uuid.UUID{}).AnyTimes()

	mockRepository := userrepositorymock.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().CreateUser(gomock.Any()).Return(nil, nil).AnyTimes()

	userService := New(mockRepository, mockUidGen, mockCrypto)
	createUser, err := userService.CreateUser(
		util.RandomOwner(),
		util.RandomString(6),
		util.RandomOwner(),
		util.RandomEmail(),
	)

	require.Error(t, err)
	require.Empty(t, createUser)
	require.ErrorContains(t, err, msgErrorHashedPassword)
}

func TestCreateUserError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCrypto := cryptomock.NewMockCrypto(ctrl)
	mockCrypto.EXPECT().HashPassword(gomock.Any()).Return("@hashedPassword", nil).AnyTimes()

	mockUidGen := uidgenmock.NewMockUIDGen(ctrl)
	mockUidGen.EXPECT().New().Return(uuid.UUID{}).AnyTimes()

	mockRepository := userrepositorymock.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().CreateUser(gomock.Any()).Return(nil, &pq.Error{Code: "23505"}).AnyTimes()

	userService := New(mockRepository, mockUidGen, mockCrypto)
	createUser, err := userService.CreateUser(
		util.RandomOwner(),
		util.RandomString(6),
		util.RandomOwner(),
		util.RandomEmail(),
	)
	pqError := err.(*pq.Error)
	require.Error(t, err)
	require.Empty(t, createUser)
	require.Equal(t, pqError.Code.Name(), "unique_violation")
	require.Equal(t, string(pqError.Code), "23505")
}
