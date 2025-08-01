// Copyright 2020 Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	k8sclient "github.com/aws-controllers-k8s/test-infra/experimental/prow/pkg/k8s"
)

// submitProwJob submits a ProwJob to the Kubernetes cluster
func (s *Server) submitProwJob(prowJob *k8sclient.ProwJob) error {
	ctx, cancel := ContextWithDefaultTimeout()
	defer cancel()
	return s.k8sProwClient.SubmitProwJob(ctx, prowJob, s.prowJobNamespace)
}
