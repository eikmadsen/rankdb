// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "rankdb": backup Resource Client
//
// Command:
// $ goagen
// --design=github.com/Vivino/rankdb/api/design
// --out=$(GOPATH)/src/github.com/Vivino/rankdb/api
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// DeleteBackupPath computes a request path to the delete action of backup.
func DeleteBackupPath(backupID string) string {
	param0 := backupID

	return fmt.Sprintf("/backup/%s", param0)
}

// cancel backup
func (c *Client) DeleteBackup(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteBackupRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteBackupRequest create the request corresponding to the delete action endpoint of the backup resource.
func (c *Client) NewDeleteBackupRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// StatusBackupPath computes a request path to the status action of backup.
func StatusBackupPath(backupID string) string {
	param0 := backupID

	return fmt.Sprintf("/backup/%s", param0)
}

// Return backup progress
func (c *Client) StatusBackup(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStatusBackupRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatusBackupRequest create the request corresponding to the status action endpoint of the backup resource.
func (c *Client) NewStatusBackupRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
