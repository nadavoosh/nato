// Copyright © 2018 NADAV RECCA <ami@nadavrec.ca>
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

package cmd

import (
	"fmt"
	"os"
	"strings"
	"sort"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "nato",
	Short: "Simple NATO alphabet phoneticizer",
	Long: `This translates a message into the NATO Alphabet. For example:

	$ nato Good Morning
	GOLD OSCAR OSCAR DELTA
	MIKE OSCAR ROMEO NOVEMBER INDIA NOVEMBER GOLD`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				printNATOMap();
				return
			}
			m := getNATOMap(true)
			for _, x := range(args) {
				fmt.Println(phoneticize(x, m))
			}
		},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printNATOMap() {
	printMap(getNATOMap(false))
}

func getNATOMap(withExtras bool) map[string]string {
	m := make(map[string]string)
	m["a"] = "Alfa"
	m["b"] = "Bravo"
	m["c"] = "Charlie"
	m["d"] = "Delta"
	m["e"] = "Echo"
	m["f"] = "Foxtrot"
	m["g"] = "Gold"
	m["h"] = "Hotel"
	m["i"] = "India"
	m["j"] = "Juliett"
	m["k"] = "Kilo"
	m["l"] = "Lima"
	m["m"] = "Mike"
	m["n"] = "November"
	m["o"] = "Oscar"
	m["p"] = "Papa"
	m["q"] = "Quebec"
	m["r"] = "Romeo"
	m["s"] = "Sierra"
	m["t"] = "Tango"
	m["u"] = "Uniform"
	m["v"] = "Victor"
	m["w"] = "Whiskey"
	m["x"] = "X-ray"
	m["y"] = "Yankee"
	m["z"] = "Zulu"

	if withExtras {
		m["0"] = "Zero"
		m["1"] = "One"
		m["2"] = "Two"
		m["3"] = "Three"
		m["4"] = "Four"
		m["5"] = "Five"
		m["6"] = "Six"
		m["7"] = "Seven"
		m["8"] = "Eight"
		m["9"] = "Niner"

		m[" "] = "Space"
	}

	return m
}

func printMap(m map[string]string) {
	var maxLenKey int
	 var keys []string

	 for k := range m {
		 keys = append(keys, k)
		 if len(k) > maxLenKey {
            maxLenKey = len(k)
        }
	 }

	 sort.Strings(keys)

	 for _, k := range keys {
		fmt.Println(k + ": " + strings.Repeat(" ", maxLenKey - len(k)) + strings.ToUpper(m[k]))
	 }
}
func phoneticize(s string, alphabet map[string]string) string {
	var r string
	for _, letter := range(strings.ToLower(s)) {
		if val, ok := alphabet[string(letter)]; ok {
			r = r + val + " "
		} else {
			r = r + string(letter)
		}
	}
	return strings.ToUpper(r)
}
