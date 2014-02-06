package main

import (
	"github.com/gonuts/commander"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func make_cmd_put(auth aws.Auth, region aws.Region) *commander.Command {
	cmd_put := func(cmd *commander.Command, args []string) error {
		if len(args) != 2 {
			cmd.Usage()
			os.Exit(1)
		}
		matches := pat.FindStringSubmatch(args[1])
		if len(matches) == 0 {
			cmd.Usage()
			os.Exit(1)
		}
		b, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}
		typ := mime.TypeByExtension(filepath.Ext(args[0]))
		if typ == "" {
			typ = "application/octet-stream"
		}
		file := matches[2]
		if strings.HasSuffix(matches[2], "/") {
			file = path.Join(matches[2], filepath.Base(args[0]))
		}
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		err = bucket.Put(file, b, typ, s3.BucketOwnerFull)
		if err != nil {
			return err
		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_put,
		UsageLine: "put file sakura://<bucket>/path/",
		Short:     "put file",
	}
}
