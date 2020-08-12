/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kubernetes

import (
	"bytes"

	"github.com/appvia/kube-devx/pkg/kev/log"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	. "github.com/onsi/gomega"
)

var hook *test.Hook

func init() {
	// Use mem buffer in test instead of Stdout
	logBuffer := &bytes.Buffer{}
	log.SetOutput(logBuffer)
	hook = test.NewLocal(log.GetLogger())
}

func assertLog(level logrus.Level, message string, fields map[string]string) {
	Expect(hook.LastEntry().Level).To(Equal(level))
	Expect(hook.LastEntry().Message).To(Equal(message))
	for k, v := range fields {
		Expect(hook.LastEntry().Data).To(HaveKeyWithValue(k, v))
	}
}
