/*
Copyright 2018 Cai Gwatkin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secrets

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"slate/internal/pkg/errors"

	cloudkms "google.golang.org/api/cloudkms/v1"
)

type Client struct {
	cloudKMSService *cloudkms.Service
	secrets         map[string]string
}

func NewClient(ctx context.Context) (*Client, error) {
	cloudKMSService, err := newCloudkmsService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed creating cloudkms service")
	}
	s := Client{
		cloudKMSService: cloudKMSService,
		secrets:         make(map[string]string),
	}
	return &s, nil
}

func (s *Client) value(filename string, loadAndDecrypt func(s *Client, filename string) (string, error)) (string, error) {
	if v, ok := s.secrets[filename]; ok {
		return v, nil
	}
	secret, err := loadAndDecrypt(s, filename)
	if err != nil {
		return "", errors.Wrap(err, "Failed decrypting secret")
	}
	s.secrets[filename] = string(secret)
	return s.secrets[filename], nil
}

func load(filename string, dst interface{}) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.Wrap(err, "Failed to retrieve current directory")
	}
	f := fmt.Sprintf("%s/%s", dir, filename)
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return errors.Wrapf(err, "Failed read file %q", f)
	}
	if err := json.Unmarshal(buf, &dst); err != nil {
		return errors.Wrapf(err, "Failed unmarshal file %q into dst", f)
	}
	return nil
}
