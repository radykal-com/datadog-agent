// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// +build linux

package model

import (
	"golang.org/x/sys/unix"
)

var (
	ptraceArchConstants = map[string]uint32{
		"PTRACE_PEEKMTETAGS":       unix.PTRACE_PEEKMTETAGS,
		"PTRACE_POKEMTETAGS":       unix.PTRACE_POKEMTETAGS,
		"PTRACE_SYSEMU":            unix.PTRACE_SYSEMU,
		"PTRACE_SYSEMU_SINGLESTEP": unix.PTRACE_SYSEMU_SINGLESTEP,
	}

	mmapFlagArchConstants = map[string]int{}
)
