// Copyright 2017 The Kubernetes Authors.
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

package auth

import (
	authAPI "github.com/kubernetes/dashboard/src/app/backend/auth/api"
	"testing"
)

func TestNewBasicAuthenticator(t *testing.T) {
	auth := NewBasicAuthenticator(&authAPI.LoginSpec{
		Username:   "username",
		Password:   "password",
		Token:      "",
		KubeConfig: "",
	})

	if _, err := auth.GetAuthInfo(); err != nil {
		t.Errorf("Failed: %v", err)
	}
}
