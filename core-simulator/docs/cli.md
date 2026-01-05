# CLI User Guide

The Core Simulator CLI (`cnsim-cli`) is an interactive shell for configuring and controlling simulations.

## Commands

### init
Configure the simulator.

- Sends configuration via `/configure` API.  
- If `initOnStartup: true` in the YAML config, the simulator ignores CLI parameters and simply resets the simulation.  

### start
Start the simulation.  
- Requires the simulator to be initialized first (`init` or `initOnStartup`).  

### status
Display the current status of the simulation.  
- Returns whether the simulation is running, stopped, or awaiting configuration.  

### stop
Stop the simulation.  
- After stopping, you must run `init` before running `start` again.  

### loadprofile
Load a simulation profile from the `cnsim-profile.yaml` file in the current working directory.  
- Multiple profiles can be defined in the YAML file.  
- The user is prompted to choose which profile to activate.  

## Profiles

`cnsim-profile.yaml` allows predefined simulation profiles. Example:

```yaml
profiles:
  default:
    plmn:
      mcc: "001"
      mnc: "06"
    dnn: "intenet"
    slice:
      sst: 1
      sd: "010203"
    numUe: 1
    gNBs: 40
    rate: 1

  basic:
    plmn:
      mcc: "001"
      mnc: "06"
    dnn: "intenet"
    slice:
      sst: 1
      sd: "010203"
    numUe: 10
    gNBs: 2
    rate: 1

  heavy:
    plmn:
      mcc: "001"
      mnc: "06"
    dnn: "intenet"
    slice:
      sst: 1
      sd: "010203"
    numUe: 500
    gNBs: 40
    rate: 10

```

Users can load a profile interactively:

```text
cnsim> loadprofile heavy
Profile 'heavy' loaded.
```

## Example CLI Workflow

```text
$ ./cnsim-cli
cnsim> loadprofile basic
Profile 'basic' loaded.
cnsim> init
Simulation configured with profile 'basic'.
cnsim> start
Simulation started.
cnsim> status
Simulation running with 10 UEs, 2 gNBs.
cnsim> stop
Simulation stopped.
```
