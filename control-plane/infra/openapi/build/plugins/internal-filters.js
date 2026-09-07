/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

module.exports = {
  id: 'internal-filters',
  decorators: {
    oas3: {
      'filter-internal-enum-values': () => ({
        Schema: {
          leave(schema) {
            const internalValues = schema['x-enum-internal'];
            delete schema['x-enum-internal'];
            if (Array.isArray(internalValues) && Array.isArray(schema.enum)) {
              schema.enum = schema.enum.filter(v => !internalValues.includes(v));
            }
          }
        }
      })
    }
  }
};
