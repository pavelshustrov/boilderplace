package bar

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"boilerplate/internal/repositiories/bar"
	"boilerplate/internal/services/bar/mocks"
)

func TestFindUsersSuccess(t *testing.T) {
	var (
		testName = "success_test"
		repoMock = new(mocks.Repository)
	)

	repoMock.On("FindUserByBars", mock.Anything, testName).
		Return([]*bar.User{
			{
				Name: "first",
			},
			{
				Name: "second",
			},
		}, nil)

	service := New(repoMock)
	names, err := service.Bar(context.Background(), testName)
	require.NoError(t, err, "no error should happened on success test")
	assert.ElementsMatch(t, []string{"second", "first"}, names, "names should match and order could be changed")
}

func TestFindUsersErrorReturned(t *testing.T) {
	var (
		repoMock = new(mocks.Repository)
		testName = "error_test"
	)

	repoMock.On("FindUserByBars", mock.Anything, testName).Return(nil, fmt.Errorf("some error"))

	service := New(repoMock)
	names, err := service.Bar(context.Background(), testName)
	require.NotNil(t, err, "error be return")
	require.Nil(t, names, "nothing as names should be return")
}
