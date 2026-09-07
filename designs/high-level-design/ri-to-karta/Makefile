# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

PUML_FILES := $(wildcard *.puml)
SVG_FILES  := $(PUML_FILES:.puml=.svg)

.PHONY: all clean

all: $(SVG_FILES)

%.svg: %.puml
	plantuml -tsvg $<

clean:
	rm -f $(SVG_FILES)
