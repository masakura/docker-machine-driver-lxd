package utils

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/lxc/lxd/shared/api"
)

func GetDockerMachineState(containerState *api.ContainerState) state.State {
	switch containerState.StatusCode {
	case api.Running:
		return state.Running
	case api.Stopped:
		return state.Stopped
	case api.OperationCreated, api.Thawed, api.Pending, api.Starting:
		return state.Starting
	case api.Stopping, api.Aborting:
		return state.Stopping
	case api.Freezing, api.Frozen:
		return state.Paused
	case api.Error:
		return state.Error
	default:
		return state.None
	}
}
