// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"time"

	"golang.org/x/net/context"

	"github.com/kasworld/go-profile"
)

func main() {
	var rundur = flag.Int("rundur", 3600*24*365, "run sec")

	profile.AddArgs()

	flag.Parse()

	if profile.IsCpu() {
		fn := profile.StartCPUProfile()
		defer fn()
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(*rundur)*time.Second)
	DoMain(ctx)

	if profile.IsMem() {
		profile.WriteHeapProfile()
	}
}

func DoMain(ctx context.Context) {
	ch1sec := time.After(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ch1sec:
			ch1sec = time.After(1 * time.Second)
		}
	}
}
