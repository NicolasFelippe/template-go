package userservice

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"template-go/internal/core/domain/users"
	cryptomock "template-go/mocks/pkg/crypto"
	uidgenmock "template-go/mocks/pkg/uidgen"
	userrepositorymock "template-go/mocks/repositories"
	"template-go/pkg/random"
	"testing"
)

var mockRepository *userrepositorymock.MockUserRepository
var mockCrypto *cryptomock.MockCrypto
var mockUidGen *uidgenmock.MockUIDGen
var userService *UserService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCrypto = cryptomock.NewMockCrypto(ctrl)
	mockUidGen = uidgenmock.NewMockUIDGen(ctrl)
	mockRepository = userrepositorymock.NewMockUserRepository(ctrl)
	userService = New(mockRepository, mockUidGen, mockCrypto)
	return func() {
		userService = nil
		defer ctrl.Finish()
	}
}

func randomUser(crypt *cryptomock.MockCrypto) (user *users.User, hashedPassword string, err error) {
	password := random.String(6)
	hashedPassword, err = crypt.HashPassword(password)

	user = &users.User{
		Username:       random.Owner(),
		HashedPassword: password,
		FullName:       random.Owner(),
		Email:          random.Email(),
	}
	return
}

func Test_Should_CreateUser_When_SendParameters_Then_Success(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockCrypto.EXPECT().HashPassword(gomock.Any()).Return("@hashedPassword", nil).AnyTimes()
	mockUidGen.EXPECT().New().Return(uuid.UUID{})

	user, hashedPassword, err := randomUser(mockCrypto)
	require.NoError(t, err)

	mockRepository.EXPECT().CreateUser(gomock.Any()).Return(user, nil).Times(1)

	createUser, err := userService.CreateUser(user)

	require.NoError(t, err)
	require.NotEmpty(t, createUser)
	require.Equal(t, createUser.FullName, user.FullName)
	require.Equal(t, createUser.Username, user.Username)
	require.Equal(t, createUser.Email, user.Email)
	require.Equal(t, createUser.HashedPassword, hashedPassword)
}

func Test_Should_CreateUser_When_SendParameters_Then_FailedCrypto(t *testing.T) {
	msgErrorHashedPassword := "failed to has crypto"
	teardown := setup(t)
	defer teardown()

	mockCrypto.EXPECT().HashPassword(gomock.Any()).Return("", errors.New(msgErrorHashedPassword)).AnyTimes()
	user, _, err := randomUser(mockCrypto)
	require.Error(t, err)

	mockUidGen.EXPECT().New().Return(uuid.UUID{}).AnyTimes()

	mockRepository.EXPECT().CreateUser(gomock.Any()).Return(nil, nil).AnyTimes()

	createUser, err := userService.CreateUser(user)

	require.Error(t, err)
	require.Empty(t, createUser)
	require.ErrorContains(t, err, msgErrorHashedPassword)
}

func Test_Should_CreateUserDuplicate_When_SendDuplicateName_Then_ErrorUniqueViolation(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockCrypto.EXPECT().HashPassword(gomock.Any()).Return("@hashedPassword", nil).AnyTimes()

	mockUidGen.EXPECT().New().Return(uuid.UUID{}).AnyTimes()

	mockRepository.EXPECT().CreateUser(gomock.Any()).Return(nil, &pq.Error{Code: "23505"}).AnyTimes()
	user, _, err := randomUser(mockCrypto)

	require.NoError(t, err)

	createUser, err := userService.CreateUser(user)
	pqError := err.(*pq.Error)

	require.Error(t, err)
	require.Empty(t, createUser)
	require.Equal(t, pqError.Code.Name(), "unique_violation")
	require.Equal(t, string(pqError.Code), "23505")
}
