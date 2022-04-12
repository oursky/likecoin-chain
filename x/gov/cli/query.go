package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gcutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
	"github.com/cosmos/cosmos-sdk/x/gov/types"

	originalcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"

	bech32utils "github.com/likecoin/likechain/bech32-migration/utils"
)

// This is to override the query commands which query results by Tx query, so that the returned addresses could be
// further processed before displaying

func convertAddr(s string) (string, error) {
	addr, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

func GetQueryCmd() *cobra.Command {
	govQueryCmd := originalcli.GetQueryCmd()
	govQueryCmd.RemoveCommand(govQueryCmd.Commands()...)
	govQueryCmd.AddCommand(
		originalcli.GetCmdQueryProposal(),
		originalcli.GetCmdQueryProposals(),
		GetCmdQueryVote(),
		GetCmdQueryVotes(),
		originalcli.GetCmdQueryParam(),
		originalcli.GetCmdQueryParams(),
		GetCmdQueryProposer(),
		GetCmdQueryDeposit(),
		GetCmdQueryDeposits(),
		originalcli.GetCmdQueryTally(),
	)
	return govQueryCmd
}

func GetCmdQueryVote() *cobra.Command {
	cmd := originalcli.GetCmdQueryVote()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx, err := client.GetClientQueryContext(cmd)
		if err != nil {
			return err
		}
		queryClient := types.NewQueryClient(clientCtx)

		// validate that the proposal id is a uint
		proposalID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
		}

		// check to see if the proposal is in the store
		ctx := cmd.Context()
		_, err = queryClient.Proposal(
			ctx,
			&types.QueryProposalRequest{ProposalId: proposalID},
		)
		if err != nil {
			return fmt.Errorf("failed to fetch proposal-id %d: %s", proposalID, err)
		}

		voterAddr, err := sdk.AccAddressFromBech32(args[1])
		if err != nil {
			return err
		}

		res, err := queryClient.Vote(
			ctx,
			&types.QueryVoteRequest{ProposalId: proposalID, Voter: args[1]},
		)
		if err != nil {
			return err
		}

		vote := res.GetVote()
		if vote.Empty() {
			params := types.NewQueryVoteParams(proposalID, voterAddr)
			resByTxQuery, err := gcutils.QueryVoteByTxQuery(clientCtx, params)

			if err != nil {
				return err
			}

			if err := clientCtx.Codec.UnmarshalJSON(resByTxQuery, &vote); err != nil {
				return err
			}
			vote.Voter = bech32utils.ConvertAccAddr(vote.Voter)
		}

		return clientCtx.PrintProto(&res.Vote)
	}
	return cmd
}

func GetCmdQueryVotes() *cobra.Command {
	cmd := originalcli.GetCmdQueryVotes()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx, err := client.GetClientQueryContext(cmd)
		if err != nil {
			return err
		}
		queryClient := types.NewQueryClient(clientCtx)

		// validate that the proposal id is a uint
		proposalID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
		}

		// check to see if the proposal is in the store
		ctx := cmd.Context()
		proposalRes, err := queryClient.Proposal(
			ctx,
			&types.QueryProposalRequest{ProposalId: proposalID},
		)
		if err != nil {
			return fmt.Errorf("failed to fetch proposal-id %d: %s", proposalID, err)
		}

		propStatus := proposalRes.GetProposal().Status
		if !(propStatus == types.StatusVotingPeriod || propStatus == types.StatusDepositPeriod) {
			page, _ := cmd.Flags().GetInt(flags.FlagPage)
			limit, _ := cmd.Flags().GetInt(flags.FlagLimit)

			params := types.NewQueryProposalVotesParams(proposalID, page, limit)
			resByTxQuery, err := gcutils.QueryVotesByTxQuery(clientCtx, params)
			if err != nil {
				return err
			}

			var votes types.Votes
			// TODO migrate to use JSONCodec (implement MarshalJSONArray
			// or wrap lists of proto.Message in some other message)
			clientCtx.LegacyAmino.MustUnmarshalJSON(resByTxQuery, &votes)
			for i := range votes {
				votes[i].Voter = bech32utils.ConvertAccAddr(votes[i].Voter)
			}
			return clientCtx.PrintObjectLegacy(votes)

		}

		pageReq, err := client.ReadPageRequest(cmd.Flags())
		if err != nil {
			return err
		}

		res, err := queryClient.Votes(
			ctx,
			&types.QueryVotesRequest{ProposalId: proposalID, Pagination: pageReq},
		)

		if err != nil {
			return err
		}

		return clientCtx.PrintProto(res)

	}

	return cmd
}

func GetCmdQueryDeposit() *cobra.Command {
	cmd := originalcli.GetCmdQueryDeposit()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx, err := client.GetClientQueryContext(cmd)
		if err != nil {
			return err
		}
		queryClient := types.NewQueryClient(clientCtx)

		// validate that the proposal id is a uint
		proposalID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("proposal-id %s not a valid uint, please input a valid proposal-id", args[0])
		}

		// check to see if the proposal is in the store
		ctx := cmd.Context()
		proposalRes, err := queryClient.Proposal(
			ctx,
			&types.QueryProposalRequest{ProposalId: proposalID},
		)
		if err != nil {
			return fmt.Errorf("failed to fetch proposal-id %d: %s", proposalID, err)
		}

		depositorAddr, err := sdk.AccAddressFromBech32(args[1])
		if err != nil {
			return err
		}

		var deposit types.Deposit
		propStatus := proposalRes.Proposal.Status
		if !(propStatus == types.StatusVotingPeriod || propStatus == types.StatusDepositPeriod) {
			params := types.NewQueryDepositParams(proposalID, depositorAddr)
			resByTxQuery, err := gcutils.QueryDepositByTxQuery(clientCtx, params)
			if err != nil {
				return err
			}
			clientCtx.Codec.MustUnmarshalJSON(resByTxQuery, &deposit)
			deposit.Depositor = bech32utils.ConvertAccAddr(deposit.Depositor)
			return clientCtx.PrintProto(&deposit)
		}

		res, err := queryClient.Deposit(
			ctx,
			&types.QueryDepositRequest{ProposalId: proposalID, Depositor: args[1]},
		)
		if err != nil {
			return err
		}

		return clientCtx.PrintProto(&res.Deposit)
	}
	return cmd
}

func GetCmdQueryDeposits() *cobra.Command {
	cmd := originalcli.GetCmdQueryDeposits()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx, err := client.GetClientQueryContext(cmd)
		if err != nil {
			return err
		}
		queryClient := types.NewQueryClient(clientCtx)

		// validate that the proposal id is a uint
		proposalID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("proposal-id %s not a valid uint, please input a valid proposal-id", args[0])
		}

		// check to see if the proposal is in the store
		ctx := cmd.Context()
		proposalRes, err := queryClient.Proposal(
			ctx,
			&types.QueryProposalRequest{ProposalId: proposalID},
		)
		if err != nil {
			return fmt.Errorf("failed to fetch proposal-id %d: %s", proposalID, err)
		}

		propStatus := proposalRes.GetProposal().Status
		if !(propStatus == types.StatusVotingPeriod || propStatus == types.StatusDepositPeriod) {
			params := types.NewQueryProposalParams(proposalID)
			resByTxQuery, err := gcutils.QueryDepositsByTxQuery(clientCtx, params)
			if err != nil {
				return err
			}

			var dep types.Deposits
			// TODO migrate to use JSONCodec (implement MarshalJSONArray
			// or wrap lists of proto.Message in some other message)
			clientCtx.LegacyAmino.MustUnmarshalJSON(resByTxQuery, &dep)

			for i := range dep {
				dep[i].Depositor = bech32utils.ConvertAccAddr(dep[i].Depositor)
			}

			return clientCtx.PrintObjectLegacy(dep)
		}

		pageReq, err := client.ReadPageRequest(cmd.Flags())
		if err != nil {
			return err
		}

		res, err := queryClient.Deposits(
			ctx,
			&types.QueryDepositsRequest{ProposalId: proposalID, Pagination: pageReq},
		)

		if err != nil {
			return err
		}

		return clientCtx.PrintProto(res)
	}
	return cmd
}

func GetCmdQueryProposer() *cobra.Command {
	cmd := originalcli.GetCmdQueryProposer()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx, err := client.GetClientQueryContext(cmd)
		if err != nil {
			return err
		}

		// validate that the proposalID is a uint
		proposalID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("proposal-id %s is not a valid uint", args[0])
		}

		prop, err := gcutils.QueryProposerByTxQuery(clientCtx, proposalID)
		if err != nil {
			return err
		}

		prop.Proposer = bech32utils.ConvertAccAddr(prop.Proposer)

		return clientCtx.PrintObjectLegacy(prop)
	}
	return cmd
}
