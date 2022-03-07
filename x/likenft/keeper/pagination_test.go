package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/likecoin/likechain/x/likenft/keeper"
	"github.com/stretchr/testify/require"
)

func TestPaginationNormalOffset(t *testing.T) {
	var actualPage1 []int
	res1, err := keeper.PaginateArray(10, &query.PageRequest{
		Limit: 3,
	}, func(i int) error {
		actualPage1 = append(actualPage1, i)
		return nil
	}, 5, 10)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: []byte("3"),
		Total:   uint64(10),
	}, res1)
	require.Equal(t, []int{0, 1, 2}, actualPage1)

	var actualPage2 []int
	res2, err := keeper.PaginateArray(10, &query.PageRequest{
		Limit:  3,
		Offset: 3,
	}, func(i int) error {
		actualPage2 = append(actualPage2, i)
		return nil
	}, 5, 10)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: []byte("6"),
		Total:   uint64(10),
	}, res2)
	require.Equal(t, []int{3, 4, 5}, actualPage2)

	var actualPage3 []int
	res3, err := keeper.PaginateArray(10, &query.PageRequest{
		Limit:  3,
		Offset: 6,
	}, func(i int) error {
		actualPage3 = append(actualPage3, i)
		return nil
	}, 5, 10)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: []byte("9"),
		Total:   uint64(10),
	}, res3)
	require.Equal(t, []int{6, 7, 8}, actualPage3)

	var actualPage4 []int
	res4, err := keeper.PaginateArray(10, &query.PageRequest{
		Limit:  3,
		Offset: 9,
	}, func(i int) error {
		actualPage4 = append(actualPage4, i)
		return nil
	}, 5, 10)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: nil,
		Total:   uint64(10),
	}, res4)
	require.Equal(t, []int{9}, actualPage4)
}

func TestPaginationNormalKey(t *testing.T) {
	var actualPage1 []int
	res1, err := keeper.PaginateArray(10, &query.PageRequest{
		Limit: 4,
	}, func(i int) error {
		actualPage1 = append(actualPage1, i)
		return nil
	}, 5, 10)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: []byte("4"),
		Total:   uint64(10),
	}, res1)
	require.Equal(t, []int{0, 1, 2, 3}, actualPage1)

	var actualPage2 []int
	res2, err := keeper.PaginateArray(10, &query.PageRequest{
		Limit: 4,
		Key:   []byte("4"),
	}, func(i int) error {
		actualPage2 = append(actualPage2, i)
		return nil
	}, 5, 10)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: []byte("8"),
		Total:   uint64(10),
	}, res2)
	require.Equal(t, []int{4, 5, 6, 7}, actualPage2)

	var actualPage3 []int
	res3, err := keeper.PaginateArray(10, &query.PageRequest{
		Limit: 4,
		Key:   []byte("8"),
	}, func(i int) error {
		actualPage3 = append(actualPage3, i)
		return nil
	}, 5, 10)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: nil,
		Total:   uint64(10),
	}, res3)
	require.Equal(t, []int{8, 9}, actualPage3)
}
