// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"strings"
	"testing"
)

func TestReadShouldFailWhenNoConfigProvided(t *testing.T) {
	_, err := Read(nil)
	if err == nil {
		t.Fatalf("should fail with when given no config")
	}
}

const validConfig = `
auth:
  region: us-phoenix-1
  tenancy: ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq
  user: ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q
  key: |
    -----BEGIN RSA PRIVATE KEY-----
    -----END RSA PRIVATE KEY-----
  fingerprint: 97:84:f7:26:a3:7b:74:d0:bd:4e:08:a7:79:c9:d0:1d

compartment: ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq

loadBalancer:
  disableSecurityListManagement: false
  subnet1: ocid1.subnet.oc1.phx.aaaaaaaasa53hlkzk6nzksqfccegk2qnkxmphkblst3riclzs4rhwg7rg57q
  subnet2: ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq
`
const validConfigNoRegion = `
auth:
  tenancy: ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq
  compartment: ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq
  user: ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q
  key: |
    -----BEGIN RSA PRIVATE KEY-----
    -----END RSA PRIVATE KEY-----
  fingerprint: 97:84:f7:26:a3:7b:74:d0:bd:4e:08:a7:79:c9:d0:1d

loadBalancer:
  disableSecurityListManagement: false
  subnet1: ocid1.subnet.oc1.phx.aaaaaaaasa53hlkzk6nzksqfccegk2qnkxmphkblst3riclzs4rhwg7rg57q
  subnet2: ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq
`

func TestReadShouldSucceedWhenProvidedValidConfig(t *testing.T) {
	_, err := Read(strings.NewReader(validConfig))
	if err != nil {
		t.Fatalf("expected no error but got '%+v'", err)
	}
}

func TestReadShouldHaveNoDefaultRegionIfNoneSpecified(t *testing.T) {
	config, err := Read(strings.NewReader(validConfigNoRegion))
	if err != nil {
		t.Fatalf("expected no error but got '%+v'", err)
	}
	if config.Auth.Region != "" {
		t.Errorf("expected no region but got %s", config.Auth.Region)
	}
}

func TestReadShouldSetCompartmentOCIDWhenProvidedValidConfig(t *testing.T) {
	cfg, err := Read(strings.NewReader(validConfig))
	if err != nil {
		t.Fatalf("expected no error but got '%+v'", err)
	}
	expected := "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq"

	if cfg.CompartmentID != expected {
		t.Errorf("Got CompartmentID = %s; want CompartmentID = %s",
			cfg.CompartmentID, expected)
	}
}
