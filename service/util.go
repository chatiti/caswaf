// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"crypto/tls"
	"net/url"

	"github.com/casbin/caswaf/object"
)

func joinPath(a string, b string) string {
	res, err := url.JoinPath(a, b)
	if err != nil {
		panic(err)
	}

	return res
}

func getCertificateForDomain(domain string) (*tls.Certificate, error) {
	site := object.GetSiteByDomain(domain)
	tlsCert, certErr := tls.X509KeyPair([]byte(site.SslCertObj.Certificate), []byte(site.SslCertObj.PrivateKey))

	return &tlsCert, certErr
}
