package cmd

import (
	"errors"

	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/ignite/cli/ignite/pkg/yaml"
	"github.com/spf13/cobra"

	"github.com/toschdev/ignite-testnet/network"
)

func NewNetworkShowInfo() *cobra.Command {
	c := &cobra.Command{
		Use:   "info [launch-id]",
		Short: "Show info details of the chain",
		Args:  cobra.ExactArgs(1),
		RunE:  networkChainShowInfoHandler,
	}
	return c
}

func networkChainShowInfoHandler(cmd *cobra.Command, args []string) error {
	session := cliui.New(cliui.StartSpinner())
	defer session.End()

	nb, launchID, err := networkChainLaunch(cmd, args, session)
	if err != nil {
		return err
	}
	n, err := nb.Network()
	if err != nil {
		return err
	}

	chainLaunch, err := n.ChainLaunch(cmd.Context(), launchID)
	if err != nil {
		return err
	}

	reward, err := n.ChainReward(cmd.Context(), launchID)
	if err != nil && !errors.Is(err, network.ErrObjectNotFound) {
		return err
	}
	chainLaunch.Reward = reward.RemainingCoins.String()

	info, err := yaml.Marshal(cmd.Context(), chainLaunch)
	if err != nil {
		return err
	}

	return session.Print(info)
}
