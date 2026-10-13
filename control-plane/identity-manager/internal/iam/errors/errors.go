/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrInternal     = errors.New("internal error")
	ErrUnauthorized = errors.New("unauthorized")
	ErrConflict     = errors.New("conflict")
)

func HandleKeyCloakError(err error) error {
	if err == nil {
		return nil
	}
	keycloakError, ok := err.(*gocloak.APIError)
	if !ok {
		return fmt.Errorf("keycloak error: %s: %w", err.Error(), ErrInternal)
	}

	if keycloakError.Code == http.StatusNotFound || strings.Contains(keycloakError.Message, "not found") {
		return fmt.Errorf("keycloak error: %s: %w", err.Error(), ErrNotFound)
	}
	if keycloakError.Code == http.StatusConflict || strings.Contains(keycloakError.Message, "conflict") {
		return fmt.Errorf("keycloak error: %s: %w", err.Error(), ErrConflict)
	}
	if keycloakError.Code == http.StatusBadRequest || strings.Contains(keycloakError.Message, "bad request") {
		return fmt.Errorf("keycloak error: %s: %w", err.Error(), ErrInvalidInput)
	}
	if keycloakError.Code == http.StatusUnauthorized || strings.Contains(keycloakError.Message, "unauthorized") {
		return fmt.Errorf("keycloak error: %s: %w", err.Error(), ErrUnauthorized)
	}
	if keycloakError.Code == http.StatusForbidden || strings.Contains(keycloakError.Message, "forbidden") {
		return fmt.Errorf("keycloak error: %s: %w", err.Error(), ErrUnauthorized)
	}
	return fmt.Errorf("keycloak error: %s: %w", err.Error(), ErrInternal)
}
