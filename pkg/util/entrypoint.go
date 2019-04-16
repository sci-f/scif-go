// Copyright (C) 2019 Vanessa Sochat.

// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public
// License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package util

import (
	"github.com/google/shlex"
	"strings"
)

// ParseEntrypoint will handle special characters in the entrypoint command
// the input should be a string from the command line
//     Special characters in the entrypoint should be replaced
//            [e] in the command or entrypoint: environment vars --> $
//            [out] in the command or entrypoint: environment vars --> >
//            [in] in the command or entrypoint: environment vars --> <
//            [pipe] in the command or entrypoint: environment vars --> |

func ParseEntrypoint(entrypoint string) []string {

	entrypoint = strings.Replace(entrypoint, "[e]", "$", -1)
	entrypoint = strings.Replace(entrypoint, "[out]", ">", -1)
	entrypoint = strings.Replace(entrypoint, "[in]", "<", -1)
	entrypoint = strings.Replace(entrypoint, "[pipe]", "|", -1)
	entrypoint = strings.Replace(entrypoint, "[append]", "|", -1)

	entrylist, _ := shlex.Split(entrypoint)
	return entrylist
}
