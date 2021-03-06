// Copyright 2016 CNI authors
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

package fakes

import "context"

type RawExec struct {
	ExecPluginCall struct {
		Received struct {
			PluginPath string
			StdinData  []byte
			Environ    []string
		}
		Returns struct {
			ResultBytes []byte
			Error       error
		}
	}
	FindInPathCall struct {
		Received struct {
			Plugin string
			Paths  []string
		}
		Returns struct {
			Path  string
			Error error
		}
	}
}

func (e *RawExec) ExecPlugin(ctx context.Context, pluginPath string, stdinData []byte, environ []string) ([]byte, error) {
	e.ExecPluginCall.Received.PluginPath = pluginPath
	e.ExecPluginCall.Received.StdinData = stdinData
	e.ExecPluginCall.Received.Environ = environ
	return e.ExecPluginCall.Returns.ResultBytes, e.ExecPluginCall.Returns.Error
}

func (e *RawExec) FindInPath(plugin string, paths []string) (string, error) {
	e.FindInPathCall.Received.Plugin = plugin
	e.FindInPathCall.Received.Paths = paths
	return e.FindInPathCall.Returns.Path, e.FindInPathCall.Returns.Error
}
