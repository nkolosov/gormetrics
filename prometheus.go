// Copyright 2019 Profects Group B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gormetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"math"
)

// counterVecCreator allows for mass creation of counter vectors in the same
// Prometheus namespace and with equal constant labels.
type counterVecCreator struct {
	namespace string
	labels    []string
}

// new creates a new prometheus.CounterVec based on the specified name and
// values in the counterVecCreator.
func (c counterVecCreator) new(
	name string,
	help string,
) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: c.namespace,
			Name:      name,
			Help:      help,
		},
		c.labels,
	)
}

// gaugeVecCreator allows for mass creation of gauge vectors in the same
// Prometheus namespace and with equal constant labels.
type gaugeVecCreator struct {
	namespace string
	labels    []string
}

// new creates a new prometheus.GaugeVec based on the specified name and
// values in the counterCreator.
func (c gaugeVecCreator) new(
	name string,
	help string,
) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: c.namespace,
			Name:      name,
			Help:      help,
		},
		c.labels,
	)
}

// histogramVecCreator allows for mass creation of histogram vectors in the same
// Prometheus namespace and with equal constant labels.
type histogramVecCreator struct {
	namespace string
	labels    []string
}

// new creates a new prometheus.GaugeVec based on the specified name and
// values in the counterCreator.
func (c histogramVecCreator) new(
	name string,
	help string,
) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: c.namespace,
			Name:      name,
			Help:      help,
			Buckets:   []float64{.005, .01, .025, .05, .1, .15, .2, .25, .5, .75, 1, 1.5, 2, 2.5, 3, 4, 5, 10, 15, 30, 60, math.Inf(1)},
		},
		c.labels,
	)
}
