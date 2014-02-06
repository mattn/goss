package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"os"
)

func make_cmd_ls(auth aws.Auth, region aws.Region) *commander.Command {
	cmd_ls := func(cmd *commander.Command, args []string) error {
		if len(args) != 2 {
			cmd.Usage()
			os.Exit(1)
		}
		matches := pat.FindStringSubmatch(args[1])
		if len(matches) == 0 {
			cmd.Usage()
			os.Exit(1)
		}
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		res, err := bucket.List(matches[2], "/", "", 1000)
		if err != nil {
			return err
		}
		for _, c := range res.CommonPrefixes {
			fmt.Println(c)
		}
		for _, c := range res.Contents {
			fmt.Println(c.Key)
		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_ls,
		UsageLine: "ls [options] [path]",
		Short:     "list files",
	}
}
