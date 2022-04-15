package bar

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/require"
)

func TestFindUsersByBarsNoRowsErrorReturned(t *testing.T) {
	var (
		testName = "NoRows"
	)
	db, rowMock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		_ = db.Close()
	}()
	rowMock.ExpectQuery(FindUserByBars).WithArgs(testName).WillReturnError(pgx.ErrNoRows)

	application, err := newrelic.NewApplication(newrelic.ConfigEnabled(false))
	require.NoError(t, err, "dummy NewRelic Application should be create without error")
	repo := New(application, db)
	bars, err := repo.FindUserByBars(context.Background(), testName)
	require.Nil(t, bars, "no rows should be found")
	require.NoError(t, err, "no rows is not an error")
}

func TestFindUsersByBarsUsersFound(t *testing.T) {
	var (
		testName = "FewRows"
	)
	db, rowMock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		_ = db.Close()
	}()

	rows := rowMock.NewRows([]string{"name"}).
		AddRow("Ann").
		AddRow("Boris")
	rowMock.ExpectQuery(FindUserByBars).WithArgs(testName).WillReturnRows(rows)

	application, err := newrelic.NewApplication(newrelic.ConfigEnabled(false))
	require.NoError(t, err, "dummy NewRelic Application should be create without error")

	repo := New(application, db)
	actual, err := repo.FindUserByBars(context.Background(), testName)
	require.NoError(t, err, "no error expected")
	assert.ElementsMatch(t, []*User{{Name: "Boris"}, {Name: "Ann"}}, actual)
}
