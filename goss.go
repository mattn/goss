package main

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
	"mime"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func usage() {
	fmt.Printf(`Usage:
    goss ls sakura://<bucket>/path/to/
    goss put file.txt sakura://<bucket>/path/to/
    goss get sakura://<bucket>/path/to/file.txt
    goss del sakura://<bucket>/path/to/file.txt
    goss cat sakura://<bucket>/path/to/file.txt
`)
	os.Exit(1)
}

var pat = regexp.MustCompile("^sakura://(.+)/(.*)")

func main() {
	auth := aws.Auth{
		AccessKey: os.Getenv("SAKURA_STORAGE_USER_ID"),
		SecretKey: os.Getenv("SAKURA_STORAGE_ACCESS_TOKEN"),
	}
	region := aws.Region{
		"sakura-storage",
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
		"",
		true,
		true,
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
	}

	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "ls":
		if len(os.Args) != 3 {
			usage()
		}
		matches := pat.FindStringSubmatch(os.Args[2])
		if len(matches) == 0 {
			usage()
		}
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		res, err := bucket.List(matches[2], "/", "", 1000)
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range res.Contents {
			fmt.Println(c.Key)
		}
	case "put":
		if len(os.Args) != 4 {
			usage()
		}
		matches := pat.FindStringSubmatch(os.Args[2])
		if len(matches) == 0 {
			usage()
		}
		b, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		typ := mime.TypeByExtension(filepath.Ext(os.Args[2]))
		if typ == "" {
			typ = "application/octet-stream"
		}
		file := path.Join(matches[2], os.Args[3])
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		err = bucket.Put(file, b, typ, s3.BucketOwnerFull)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "get":
		if len(os.Args) != 3 {
			usage()
		}
		matches := pat.FindStringSubmatch(os.Args[2])
		if len(matches) == 0 {
			usage()
		}
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		b, err := bucket.Get(matches[2])
		if err != nil {
			log.Fatal(err.Error())
		}
		p := path.Base(matches[2])
		err = ioutil.WriteFile(p, b, 0664)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "cat":
		if len(os.Args) != 3 {
			usage()
		}
		matches := pat.FindStringSubmatch(os.Args[2])
		if len(matches) == 0 {
			usage()
		}
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		b, err := bucket.Get(matches[2])
		if err != nil {
			log.Fatal(err.Error())
		}
		os.Stdout.Write(b)
	case "del":
		if len(os.Args) != 3 {
			usage()
		}
		matches := pat.FindStringSubmatch(os.Args[2])
		if len(matches) == 0 {
			usage()
		}
		s := s3.New(auth, region)
		bucket := s.Bucket(matches[1])
		err := bucket.Del(matches[2])
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		usage()
	}
}
