package shell

import (
	"context"
	"io"

	tar "github.com/ipfs/boxo/tar"
	shell "github.com/ipfs/go-ipfs-api"
	"ipfs-cluster/internal/logger"
)

func Cat(s *shell.Shell, path string) (io.ReadCloser, error) {
	logger.GetIPFSLogger().Debugf("starting cat %v from ipfs node", path)
	resp, err := s.Request("cat", path).
		Option("decrypt", true).
		Send(context.Background())
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Output, nil
}

func Get(s *shell.Shell, hash, outdir string) error {
	logger.GetIPFSLogger().Debugf("starting get %v from ipfs node", hash)
	resp, err := s.Request("get", hash).
		Option("create", true).
		Option("decrypt", true).
		Send(context.Background())
	if err != nil {
		return err
	}
	defer resp.Close()

	if resp.Error != nil {
		return resp.Error
	}

	extractor := &tar.Extractor{Path: outdir}
	return extractor.Extract(resp.Output)
}
