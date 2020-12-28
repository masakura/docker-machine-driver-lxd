package utils

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperationCreatedIsStarting(t *testing.T) {
	assert.Equal(t, state.Starting, GetDockerMachineState(State(api.OperationCreated)))
}

func TestStartedIsStarting(t *testing.T) {
	assert.Equal(t, state.Starting, GetDockerMachineState(State(api.Starting)))
}

func TestStoppedIsStopped(t *testing.T) {
	assert.Equal(t, state.Stopped, GetDockerMachineState(State(api.Stopped)))
}

func TestRunningIsRunning(t *testing.T) {
	assert.Equal(t, state.Running, GetDockerMachineState(State(api.Running)))
}

func TestPendingIsStarting(t *testing.T) {
	assert.Equal(t, state.Starting, GetDockerMachineState(State(api.Pending)))
}

func TestStartingIsStarting(t *testing.T) {
	assert.Equal(t, state.Starting, GetDockerMachineState(State(api.Starting)))
}

func TestStoppingIsStopping(t *testing.T) {
	assert.Equal(t, state.Stopping, GetDockerMachineState(State(api.Stopping)))
}

func TestAbortingIsStopping(t *testing.T) {
	assert.Equal(t, state.Stopping, GetDockerMachineState(State(api.Aborting)))
}

func TestFreezingIsPaused(t *testing.T) {
	assert.Equal(t, state.Paused, GetDockerMachineState(State(api.Freezing)))
}

func TestFrozenIsPaused(t *testing.T) {
	assert.Equal(t, state.Paused, GetDockerMachineState(State(api.Frozen)))
}

func TestThawedIsStarting(t *testing.T) {
	assert.Equal(t, state.Starting, GetDockerMachineState(State(api.Thawed)))
}

func TestErrorIsError(t *testing.T) {
	assert.Equal(t, state.Error, GetDockerMachineState(State(api.Error)))
}

func State(statusCode api.StatusCode) *api.ContainerState {
	return &api.ContainerState{
		StatusCode: statusCode,
	}
}
