// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package monitors

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/mapval"
	"github.com/elastic/beats/libbeat/testing/mapvaltest"
)

func TestURLFields(t *testing.T) {
	tests := []struct {
		name string
		u    string
		want common.MapStr
	}{
		{
			"simple-http",
			"http://elastic.co",
			common.MapStr{
				"full":   "http://elastic.co",
				"scheme": "http",
				"domain": "elastic.co",
			},
		},
		{
			"simple-https",
			"https://elastic.co",
			common.MapStr{
				"full":   "https://elastic.co",
				"scheme": "https",
				"domain": "elastic.co",
			},
		},
		{
			"fancy-proto",
			"tcp+ssl://elastic.co",
			common.MapStr{
				"full":   "tcp+ssl://elastic.co",
				"scheme": "tcp+ssl",
				"domain": "elastic.co",
			},
		},
		{
			"complex",
			"tcp+ssl://myuser:mypass@elastic.co/foo/bar?q=dosomething&x=y",
			common.MapStr{
				"full":     "tcp+ssl://myuser:%3Chidden%3E@elastic.co/foo/bar?q=dosomething&x=y",
				"scheme":   "tcp+ssl",
				"domain":   "elastic.co",
				"path":     "/foo/bar",
				"query":    "q=dosomething&x=y",
				"username": "myuser",
				"password": "<hidden>",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := url.Parse(tt.u)
			require.NoError(t, err)

			got := URLFields(parsed)
			mapvaltest.Test(t, mapval.MustCompile(mapval.Map(tt.want)), got)
		})
	}
}
