package commands

import (
	"fmt"
	cmds "github.com/ipfs/go-ipfs-cmds"
	"github.com/ipfs/go-ipfs-provider/batched"
	"github.com/ipfs/go-ipfs/core/commands/cmdenv"
	"time"
)

type providerStat struct {
	TotalToProvide, NumRecentlyProvided, TotalProvidesLifetime int
	AvgProvideTime                                             time.Duration
}

var statProvideCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Returns statistics about the node's (re)provider system.",
		ShortDescription: `
Returns statistics about the content the node is advertising.

This interface is not stable and may change from release to release.
`,
	},
	Arguments: []cmds.Argument{},
	Options:   []cmds.Option{},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		nd, err := cmdenv.GetNode(env)
		if err != nil {
			return err
		}

		if !nd.IsOnline {
			return ErrNotOnline
		}

		sys, ok := nd.Provider.(*batched.BatchProvidingSystem)
		if !ok {
			return fmt.Errorf("can only return stats if the experimental DHT client is enabled")
		}

		total, recentlyProvided, totalProvides, avgProvideTime, err := sys.Stat(req.Context)
		if err != nil {
			return err
		}

		if err := res.Emit(providerStat{
			TotalToProvide:        total,
			NumRecentlyProvided:   recentlyProvided,
			TotalProvidesLifetime: totalProvides,
			AvgProvideTime:        avgProvideTime,
		}); err != nil {
			return err
		}

		return nil
	},
	Encoders: cmds.EncoderMap{},
	Type:     providerStat{},
}
