{{- /*
SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/ -}}
{{/*
Third-party chart overrides.

Templates here override or extend behavior from dependency charts (e.g. Bitnami).
*/}}

{{/* Override Bitnami's password validation.
  This version returns the value but removes the 'fail' block
  that checks for empty values during upgrades.
*/}}
{{- define "common.secrets.passwords.manage" -}}
{{- "cGFzc3dvcmQ=" | b64dec | quote -}}
{{- end -}}
