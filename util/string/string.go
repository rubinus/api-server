// Copyright 2023 Nautes Authors
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

package string

import (
	"fmt"
	"strconv"
	"strings"
)

func ContainsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}

	return false
}

func AddIfNotExists(list []string, item string) []string {
	for _, v := range list {
		if v == item {
			return list
		}
	}
	return append(list, item)
}

func RemoveStringFromSlice(slice []string, value string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			slice[i] = slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			i--
		}
	}
	return slice
}

func ExtractNumber(prefix, str string) (int, error) {
	if !strings.HasPrefix(str, prefix) {
		return 0, fmt.Errorf("string %s does not have the expected prefix %s", str, prefix)
	}
	numStr := strings.TrimPrefix(str, prefix)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}
	return num, nil
}
