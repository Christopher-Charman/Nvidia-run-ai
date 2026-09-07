/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package mocks

import (
	"net"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/cluster-api/internal/websocket/connection"
)

type MockWebSocket struct {
	net.Conn
}

func NewMockWebSocket() *MockWebSocket {
	conn, _ := net.Pipe()
	return &MockWebSocket{conn}
}

func (mws *MockWebSocket) WriteText(p []byte) error {
	return nil
}

func (mws *MockWebSocket) SetMessageHandler(_ connection.MessageHandler) {
}
