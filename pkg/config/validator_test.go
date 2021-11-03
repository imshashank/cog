package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateConfig(t *testing.T) {
	config := &Config{
		Build: &Build{
			GPU:           true,
			PythonVersion: "3.8",
			PythonPackages: []string{
				"tensorflow==1.15.0",
				"foo==1.0.0",
			},
			CUDA: "10.0",
		},
	}
	err := ValidateConfig(config, "1.0")
	require.NoError(t, err)
}

func TestValidateSuccess(t *testing.T) {
	config := `build:
  gpu: true
  system_packages:
    - "libgl1-mesa-glx"
    - "libglib2.0-0"
  python_version: "3.8"
  python_packages:
    - "torch==1.8.1"`

	err := Validate(config, "1.0")
	require.NoError(t, err)
<<<<<<< HEAD
}

func TestValidatePythonVersionNumerical(t *testing.T) {
	config := `build:
  gpu: true
  system_packages:
    - "libgl1-mesa-glx"
    - "libglib2.0-0"
  python_version: 3.8
  python_packages:
    - "torch==1.8.1"`

	err := Validate(config, "1.0")
	require.NoError(t, err)
=======
>>>>>>> bf8b1be (Fix broken tests)
}

func TestValidateBuildIsRequired(t *testing.T) {
	config := `buildd:
  gpu: true
  system_packages:
    - "libgl1-mesa-glx"
    - "libglib2.0-0"
  python_version: "3.8"
  python_packages:
    - "torch==1.8.1"`

	err := Validate(config, "1.0")
	require.Error(t, err)
	require.Contains(t, err.Error(), "build is required")
}

func TestValidatePythonVersionIsRequired(t *testing.T) {
	config := `build:
  gpu: true
  system_packages:
    - "libgl1-mesa-glx"
    - "libglib2.0-0"
  python_packages:
    - "torch==1.8.1"`

	err := Validate(config, "1.0")
	require.Error(t, err)
	require.Contains(t, err.Error(), "python_version is required")
}
