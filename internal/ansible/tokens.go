// Copyright 2025 TechDivision GmbH
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

package ansible

import "strings"

type KnownFlag struct {
	Long, Short string
}

var KnownFlags = struct {
	Verbose KnownFlag
	Debug   KnownFlag
}{
	Verbose: KnownFlag{Long: "--verbose", Short: "-v"},
	Debug:   KnownFlag{Long: "--debug", Short: "-d"},
}

func SplitTokens(tokens []string) (args, opts []string, verbose, debug bool) {
	for _, tok := range tokens {
		switch tok {
		case KnownFlags.Verbose.Long, KnownFlags.Verbose.Short:
			verbose = true
		case KnownFlags.Debug.Long, KnownFlags.Debug.Short:
			debug = true
		default:
			if strings.HasPrefix(tok, "-") {
				opts = append(opts, tok)
			} else {
				args = append(args, tok)
			}
		}
	}
	return args, opts, verbose, debug
}
