package main

import (
	"github.com/gonuts/commander"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"os"
	"path"
)

func make_cmd_get(auth aws.Auth, region aws.Region) *commander.Command {
	cmd_get := func(cmd *commander.Command, args []string) error {
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
		b, err := bucket.Get(matches[2])
		if err != nil {
			return err
		}
		p := path.Base(matches[2])
		err = ioutil.WriteFile(p, b, 0664)
		if err != nil {
			return err
		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_get,
		UsageLine: "get [options] sakura://<bucket>/path/file",
		Short:     "get file",
	}
}
