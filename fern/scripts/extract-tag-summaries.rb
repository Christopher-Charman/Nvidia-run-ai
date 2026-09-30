#!/usr/bin/env ruby
# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# POC: move OpenAPI tag descriptions into markdown files and wire them into a
# Fern product layout as `summary:`, so tag copy + doc links live in docs/, not the spec.
# usage: extract-tag-summaries.rb <openapi.yml> <product.yml> <api-name> <out-dir-under-docs>
require "yaml"; require "fileutils"
spec, product, api_name, out = ARGV
tags = YAML.load_file(spec)["tags"].select { |t| t["description"].to_s.strip != "" }
id_of = ->(n) { w = n.split(/[^A-Za-z0-9]+/).reject(&:empty?).map { |x| x =~ /^[A-Z0-9]+$/ ? x.downcase : x }; w[0][0].downcase + w[0][1..] + w[1..].map { |x| x[0].upcase + x[1..] }.join }
slug_of = ->(n) { n.gsub(/([a-z])([A-Z])/, '\1-\2').downcase.gsub(/[^a-z0-9]+/, "-").gsub(/^-|-$/, "") }
by_id = tags.to_h { |t| [id_of.(t["name"]), t] }
lines = File.readlines(product); in_api = false; done = []; n = 0
lines = lines.flat_map do |l|
  in_api = true if l =~ /api-name:\s*#{Regexp.escape(api_name)}\s*$/
  in_api = false if in_api && l =~ /^\s*- (tab|section|api):/ && l !~ /^\s{8,}/ && done.any? && false
  if in_api && (m = l.match(/^(\s*)- (\w+)\s*$/)) && (t = by_id[m[2]])
    file = "docs/#{out}/#{slug_of.(t['name'])}.md"
    FileUtils.mkdir_p(File.dirname(file)); File.write(file, t["description"].strip + "\n")
    n += 1; done << m[2]
    [ "#{m[1]}- #{m[2]}:\n", "#{m[1]}    summary: ../../#{file}\n" ]
  else [l] end
end
File.write(product, lines.join)
puts "#{n} tag summaries written (#{tags.size} tags with descriptions in spec); unmatched: #{(by_id.keys - done).join(', ')}"
