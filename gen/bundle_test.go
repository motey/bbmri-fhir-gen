// Copyright © 2019 The Samply Development Community
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

package gen

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBundle(t *testing.T) {
	bundle := Bundle(rand.New(rand.NewSource(0)), 0, 1, false)
	assert.Equal(t, 25, len(bundle["entry"].(Array)))
}

func TestEntry(t *testing.T) {
	entry := entry(Biobank())
	assert.Equal(t, "http://example.com/Organization/biobank-0", entry["fullUrl"].(string))
}
