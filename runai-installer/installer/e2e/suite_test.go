//go:build e2e

/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package e2e_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/e2e/utils"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/logging"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testEnv *utils.TestEnv

// installerBinaryPath is the installer binary compiled once for the whole suite
// in BeforeSuite and shared by every spec that drives the CLI.
var installerBinaryPath string

func buildInstallerBinary() (string, error) {
	binPath := filepath.Join(os.TempDir(), "installer-e2e-test")
	cmd := exec.Command("go", "build", "-tags", "runai_internal", "-o", binPath, "./cmd/")
	cmd.Dir = filepath.Join("..")
	cmd.Stdout = GinkgoWriter
	cmd.Stderr = GinkgoWriter
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build installer binary: %w", err)
	}
	return binPath, nil
}

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Installer E2E Suite")
}

var _ = BeforeSuite(func() {
	var err error
	testEnv, err = utils.SetupTestEnv()
	Expect(err).ToNot(HaveOccurred())

	err = logging.InitFromDir(testEnv.TempDir)
	Expect(err).ToNot(HaveOccurred())

	installerBinaryPath, err = buildInstallerBinary()
	Expect(err).ToNot(HaveOccurred())
})

var _ = ReportAfterEach(func(report SpecReport) {
	if report.Failed() {
		utils.CollectDiagnostics(testEnv.TempDir)
	}
})

var _ = AfterSuite(func() {
	if installerBinaryPath != "" {
		os.Remove(installerBinaryPath)
	}
	if testEnv != nil {
		testEnv.Cleanup()
	}
})
