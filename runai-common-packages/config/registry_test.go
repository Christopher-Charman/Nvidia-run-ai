package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfig struct {
	TestVal1 string `mapstructure:"test_val1"`
	TestVal2 string `mapstructure:"test_val2"`
	TestInt  int
	TestMap  map[string]string `mapstructure:"test_map"`
}

func (t *TestConfig) SetDefaults() {
	t.TestVal1 = "test"
}

type SecondTestConfig struct {
	TestVal3 string `mapstructure:"test_val3" validate:"required"`
}

func (t *SecondTestConfig) SetDefaults() {}

type PointerTestConfig struct {
	TestVal1 *bool `mapstructure:"test_val1"`
}

func (t *PointerTestConfig) SetDefaults() {}

func TestLoad_PointerConfig_True(t *testing.T) {
	config := `
test_val1: true
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)
	registry.Register(&PointerTestConfig{})
	result := registry.Get(&PointerTestConfig{}).(*PointerTestConfig)
	assert.Equal(t, true, *result.TestVal1)
}

func TestLoad_PointerConfig_False(t *testing.T) {
	config := `
test_val1: false
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)
	registry.Register(&PointerTestConfig{})
	result := registry.Get(&PointerTestConfig{}).(*PointerTestConfig)
	assert.Equal(t, false, *result.TestVal1)
}

func TestLoad_PointerConfig_Nil(t *testing.T) {
	config := `
test_val2: false
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)
	registry.Register(&PointerTestConfig{})
	result := registry.Get(&PointerTestConfig{}).(*PointerTestConfig)
	assert.Nil(t, result.TestVal1)
}

func TestSet_SetsTheStruct(t *testing.T) {
	registry := newRegistry()

	testConf := TestConfig{TestVal1: "blala"}
	registry.Set(&testConf)
	result := registry.Get(&TestConfig{}).(*TestConfig)
	assert.Equal(t, "blala", result.TestVal1)
}

func TestSet_SetsWithDefaults(t *testing.T) {
	registry := newRegistry()

	testConf := TestConfig{TestVal2: "bla"}
	registry.SetWithDefaults(&testConf)
	result := registry.Get(&TestConfig{}).(*TestConfig)
	assert.Equal(t, "test", result.TestVal1)
	assert.Equal(t, "bla", result.TestVal2)
}

func TestSetConfigFiles_Loads(t *testing.T) {
	config := `
test_val1: bla
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	registry.Register(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, "bla", result.TestVal1)
}

func TestSetConfigFiles_LoadsAndEnvOverrides(t *testing.T) {
	config := `
test_val1: bla
test_val2: bla
`
	t.Setenv("TEST_VAL2", "overrides")

	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	registry.Register(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, "bla", result.TestVal1)
	assert.Equal(t, "overrides", result.TestVal2)
}

func TestSetConfigFiles_LoadsInt(t *testing.T) {
	config := `
testint: 1
`

	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	registry.Register(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, 1, result.TestInt)
}

func TestSetConfigFiles_LoadsIntWithRoot(t *testing.T) {
	config := `
test:
  testint: 1
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	registry.RegisterRoot(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, 1, result.TestInt)
}

func TestSetConfigFiles_RegisterRootReloads(t *testing.T) {
	config := `
test:
  testint: 1
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	change := make(chan struct{}, 1)
	registry.OnConfigChange(func() {
		change <- struct{}{}
	})

	registry.RegisterRoot(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, 1, result.TestInt)

	time.Sleep(500 * time.Millisecond) // allow time for the watch go routine to start
	config_change := `
test:
  testint: 11
`
	rewriteFile(t, f.Name(), config_change)

	<-change

	result = registry.Get(&TestConfig{}).(*TestConfig)
	assert.Equal(t, 11, result.TestInt)
}

func TestSetConfigFiles_RegisterReloads(t *testing.T) {
	config := `
testint: 1
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	change := make(chan struct{}, 1)
	registry.OnConfigChange(func() {
		change <- struct{}{}
	})

	registry.Register(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, 1, result.TestInt)

	time.Sleep(500 * time.Millisecond) // allow time for the watch go routine to start
	config_change := `
testint: 11
`
	rewriteFile(t, f.Name(), config_change)

	<-change

	result = registry.Get(&TestConfig{}).(*TestConfig)
	assert.Equal(t, 11, result.TestInt)
}

func TestOnConfigChange_CallsCallback(t *testing.T) {
	config := `
test:
  testint: 1
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	change := make(chan struct{}, 1)
	registry.OnConfigChange(func() {
		change <- struct{}{}
	})

	registry.RegisterRoot(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, 1, result.TestInt)

	time.Sleep(500 * time.Millisecond) // allow time for the watch go routine to start
	config_change := `
test:
  testint: 11
`
	rewriteFile(t, f.Name(), config_change)

	<-change
	assert.Len(t, change, 0)
}

func TestWatchConfigFiles_NotifiesChangeForMultipleFiles(t *testing.T) {
	f1, err := createTempFile(`
test:
  testint: 1
`)
	assert.Nil(t, err)
	defer os.Remove(f1.Name())

	f2, err := createTempFile(`
second_test:
  test_val3: bla
`)
	assert.Nil(t, err)
	defer os.Remove(f2.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f1.Name(), f2.Name())
	require.Nil(t, err)

	change := make(chan struct{}, 1)
	registry.OnConfigChange(func() { change <- struct{}{} })

	registry.RegisterRoot(&TestConfig{})
	assert.Equal(t, 1, registry.Get(&TestConfig{}).(*TestConfig).TestInt)

	registry.RegisterRoot(&SecondTestConfig{})
	assert.Equal(t, "bla", registry.Get(&SecondTestConfig{}).(*SecondTestConfig).TestVal3)

	time.Sleep(500 * time.Millisecond)

	rewriteFile(t, f1.Name(), `
test:
  testint: 11
`)
	<-change
	assert.Equal(t, 11, registry.Get(&TestConfig{}).(*TestConfig).TestInt)

	rewriteFile(t, f2.Name(), `
second_test:
  test_val3: bla1
`)
	<-change
	assert.Equal(t, "bla1", registry.Get(&SecondTestConfig{}).(*SecondTestConfig).TestVal3)
	assert.Len(t, change, 0)
}

func TestWatchConfigFiles_WatcherResetsWhenSetFilesCalledAgain(t *testing.T) {
	f1, err := createTempFile(`
test:
  testint: 1
`)
	assert.Nil(t, err)
	defer os.Remove(f1.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f1.Name())
	require.Nil(t, err)

	calledCount := 0
	registry.OnConfigChange(func() { calledCount++ })

	registry.RegisterRoot(&TestConfig{})
	assert.Equal(t, 1, registry.Get(&TestConfig{}).(*TestConfig).TestInt)

	f2, err := createTempFile(`
test:
  testint: 11
`)
	assert.Nil(t, err)
	defer os.Remove(f2.Name())

	err = registry.SetConfigFiles(f2.Name())
	require.Nil(t, err)
	registry.RegisterRoot(&TestConfig{})

	assert.Equal(t, 11, registry.Get(&TestConfig{}).(*TestConfig).TestInt)

	rewriteFile(t, f1.Name(), `
test:
  testint: 2
`)

	assert.Equal(t, 0, calledCount)
}

func TestSetConfigFiles_LoadsIntWithRootAndEnv(t *testing.T) {
	config := `
test:
  testint: 1
`
	t.Setenv("TEST_TESTINT", "2")
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	registry.RegisterRoot(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, 2, result.TestInt)
}

func TestSetConfigFiles_WithoutRootInEnvWontOverride(t *testing.T) {
	config := `
test:
  testint: 1
`
	t.Setenv("TESTINT", "2")
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	registry.RegisterRoot(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, 1, result.TestInt)
}

func TestSetConfigFiles_FilePathTakenFromEnv(t *testing.T) {
	config := `
test_val1: bla
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	t.Setenv(envConfigFilePath, f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles()
	require.Nil(t, err)

	registry.Register(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, "bla", result.TestVal1)
}

func TestSetConfigFiles_WithoutConfig(t *testing.T) {
	registry := newRegistry()
	err := registry.SetConfigFiles()
	require.Equal(t, err.Error(), "no config file paths provided")
}

func TestRegister_ValidationFails(t *testing.T) {
	registry := newRegistry()
	err := registry.Register(&SecondTestConfig{})
	require.Contains(t, err.Error(), "failed to validate config struct SecondTestConfig")
}

func TestRegister_ValidationSucceeds(t *testing.T) {
	f, err := createTempFile(`
test_val3: bla
`)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	err = registry.Register(&SecondTestConfig{})
	require.NoError(t, err)
}

func TestRegisterRoot_ValidationFails(t *testing.T) {
	registry := newRegistry()
	err := registry.RegisterRoot(&SecondTestConfig{})
	require.Contains(t, err.Error(), "failed to validate config struct SecondTestConfig")
}

func TestRegisterRoot_ValidationSucceeds(t *testing.T) {
	f, err := createTempFile(`
second_test:
  test_val3: bla
`)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()
	err = registry.SetConfigFiles(f.Name())
	require.Nil(t, err)

	err = registry.RegisterRoot(&SecondTestConfig{})
	require.NoError(t, err)
}

func createTempFile(data string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "temp-config-*.yaml")
	if err != nil {
		return nil, err
	}

	if _, err := tempFile.Write([]byte(data)); err != nil {
		return nil, err
	}
	if err := tempFile.Close(); err != nil {
		return nil, err
	}
	return tempFile, nil
}

func TestGetEnvOrDefault_ReturnsEnvVarValue(t *testing.T) {
	t.Setenv("TEST_VAR", "test")

	val := getEnvOrDefault("TEST_VAR", "default")
	assert.Equal(t, val, "test")
}

func TestGetEnvOrDefault_ReturnsDefaultValue(t *testing.T) {
	val := getEnvOrDefault("TEST_VAR", "default")
	assert.Equal(t, val, "default")
}

func TestIsDir_ExistingFolder(t *testing.T) {
	baseDir, err := os.MkdirTemp("", "test_isdir_*")
	assert.NoError(t, err)
	defer os.RemoveAll(baseDir)

	folderPath := filepath.Join(baseDir, "test_folder")
	err = os.Mkdir(folderPath, 0755)
	assert.NoError(t, err)

	isDirectory, err := isDir(baseDir, "test_folder")
	assert.True(t, isDirectory)
	assert.NoError(t, err)
}

func TestIsDir_ExistingFile(t *testing.T) {
	baseDir, err := os.MkdirTemp("", "test_isdir_*")
	assert.NoError(t, err)
	defer os.RemoveAll(baseDir)

	filePath := filepath.Join(baseDir, "test_file")
	err = os.WriteFile(filePath, []byte("content"), 0644)
	assert.NoError(t, err)

	isDirectory, err := isDir(baseDir, "test_file")
	assert.False(t, isDirectory)
	assert.NoError(t, err)
}

func TestIsDir_NonExistingEntry(t *testing.T) {
	isDirectory, err := isDir(os.TempDir(), "non_existent_123")
	assert.False(t, isDirectory)
	assert.Error(t, err)
}

func TestTrimConfigSuffixFromPascal(t *testing.T) {
	assert.Equal(t, "test", trimConfigSuffixFromPascal("test_config"))
	assert.Equal(t, "test", trimConfigSuffixFromPascal("test_cfg"))
	assert.Equal(t, "test", trimConfigSuffixFromPascal("test_conf"))
	assert.Equal(t, "test", trimConfigSuffixFromPascal("test_configuration"))
	assert.Equal(t, "test_configuration", trimConfigSuffixFromPascal("test_configuration_config"))

}

func TestSetConfigFiles_WithCustomDelimiter(t *testing.T) {
	t.Setenv(envConfigKeyDelimiter, ";")
	t.Setenv("TEST_VAL1", "bla")

	config := `
test_map:
  1.1.1: value1
  1.2.1: value2
`
	f, err := createTempFile(config)
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	registry := newRegistry()

	err = registry.SetConfigFiles(f.Name())
	require.NoError(t, err)

	registry.Register(&TestConfig{})
	result := registry.Get(&TestConfig{}).(*TestConfig)

	assert.Equal(t, result.TestVal1, "bla")
	assert.Equal(t, result.TestMap["1.1.1"], "value1")
	assert.Equal(t, result.TestMap["1.2.1"], "value2")
}

func rewriteFile(t *testing.T, filename, config_change string) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0)
	require.NoError(t, err)
	_, err = f.WriteAt([]byte(config_change), 0)
	require.NoError(t, err)
	err = f.Close()
	require.NoError(t, err)
}
