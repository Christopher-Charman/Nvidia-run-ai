# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

import sys
import json

if __name__ == '__main__':
    data = json.load(sys.stdin)

    unsuccessful_jobs = []

    for job in data:
        if data[job]['result'] != 'success' and data[job]['result'] != 'skipped': 
            unsuccessful_jobs.append(job)

    if len(unsuccessful_jobs) > 0:
        print(f"Unsuccessful jobs: {unsuccessful_jobs}")
        sys.exit(1)