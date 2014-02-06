package main

import (
	"github.com/gonuts/commander"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"os"
)

func make_cmd_del(auth aws.Auth, region aws.Region) *commander.Command {
	cmd_del := func(cmd *commander.Command, args []string) error {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}
		matches := pat.FindStringSubmatch(args[0])
		if len(matches) == 0 {
			cmd.Usage()
			os.Exit(1)
		}
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		err := bucket.Del(matches[2])
		if err != nil {
			return err
		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_del,
		UsageLine: "del [options] sakura://<bucket>/path/file",
		Short:     "delete file",
	}
}
