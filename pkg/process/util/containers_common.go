// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package util

import "github.com/DataDog/datadog-agent/pkg/util/containers/metrics"

// ContainerRateMetrics holds previous values for a container,
// in order to compute rates
type ContainerRateMetrics struct {
	CPU        *metrics.ContainerCPUStats
	IO         *metrics.ContainerIOStats
	NetworkSum *metrics.InterfaceNetStats
	Network    metrics.ContainerNetStats
}

// NullContainerRates can be safely used for containers that have no
// previous rate values stored (new containers)
var NullContainerRates = ContainerRateMetrics{
	CPU: &metrics.ContainerCPUStats{
		User:   -1,
		System: -1,
		Shares: -1,
	},
	IO:         &metrics.ContainerIOStats{},
	NetworkSum: &metrics.InterfaceNetStats{},
	Network:    metrics.ContainerNetStats{},
}
