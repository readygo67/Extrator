package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/readygo67/ExtractSafeMoon/config"
	"github.com/readygo67/ExtractSafeMoon/runner"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb/util"
	"os"
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "safemoond",
		Short: "safemoond Daemon (server)",
	}

	rootCmd.AddCommand(StartCmd())
	rootCmd.AddCommand(VersionCmd())
	rootCmd.AddCommand(queryCmd())
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("err:%v", err)
		os.Exit(1)
	}

}

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func StartCmd() *cobra.Command {
	var configFile string
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start the safemoond search",
		Long:  `start the safemoond search`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("starting safemoond bot")
			cfg, err := config.New(configFile)
			if err != nil {
				panic(err)
			}

			runner.Run(cfg)
			return nil
		},
	}
	cmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./config.yml", "config file (default is ./config.yaml)")
	return cmd
}

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get safemoond version",
		Long:  `Get safemoond version`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("safemoond v0.1")
			return nil
		},
	}
	return cmd
}

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func queryCmd() *cobra.Command {
	var configFile string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list subcommnd",
	}

	cmd.AddCommand(
		listTokenCommand(&configFile),
		listSafeMoonCommand(&configFile),
		listVictimCommand(&configFile),
	)
	cmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./config.yml", "config file (default is ./config.yml)")
	return cmd
}

func listTokenCommand(configFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "list tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(*configFile)
			if err != nil {
				return err
			}

			if !fileExists(cfg.DB) {
				return fmt.Errorf("db does not exist")
			}

			db, err := runner.NewDB(cfg.DB)
			defer db.Close()
			if err != nil {
				return err
			}

			count := 0
			iter := db.NewIterator(util.BytesPrefix(runner.TokenPrefix), nil)
			defer iter.Release()
			for iter.Next() {
				addr := common.BytesToAddress(iter.Value())
				log.Info(addr)
				count++
			}
			log.Infof("total:%v", count)
			return nil
		},
	}
	return cmd
}

func listSafeMoonCommand(configFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "safemoon",
		Short: "list safemoon",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(*configFile)
			if err != nil {
				return err
			}

			if !fileExists(cfg.DB) {
				return fmt.Errorf("db does not exist")
			}

			db, err := runner.NewDB(cfg.DB)
			defer db.Close()
			if err != nil {
				return err
			}

			count := 0
			iter := db.NewIterator(util.BytesPrefix(runner.SafeMoonPrefix), nil)
			defer iter.Release()
			for iter.Next() {
				addr := common.BytesToAddress(iter.Value())
				log.Info(addr)
				count++
			}
			log.Infof("total:%v", count)
			return nil
		},
	}
	return cmd
}

func listVictimCommand(configFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "victim",
		Short: "list victim",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(*configFile)
			if err != nil {
				return err
			}

			if !fileExists(cfg.DB) {
				return fmt.Errorf("db does not exist")
			}

			db, err := runner.NewDB(cfg.DB)
			defer db.Close()
			if err != nil {
				return err
			}

			count := 0
			iter := db.NewIterator(util.BytesPrefix(runner.VictimPrefix), nil)
			defer iter.Release()
			for iter.Next() {
				addr := common.BytesToAddress(iter.Value())
				log.Info(addr)
				count++
			}
			log.Infof("total:%v", count)
			return nil
		},
	}
	return cmd
}

// filesExists reports whether the named file or directory exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
