import httpx
import time
import cmd
import typer
import yaml
from dataclasses import dataclass, field


BASE_URL = "http://localhost:8081/core-simulator/v1"
client = httpx.Client()


@dataclass
class Plmn:
    mcc: str
    mnc: str

@dataclass
class Slice:
    sst: int
    sd: str

@dataclass
class SimulationProfile:
    plmn: Plmn
    dnn: str
    slice: Slice
    numUe: int
    gNBs: int
    rate: int

@dataclass
class SimulationConfig:
    profiles: dict[str, SimulationProfile]
    current_profile: str = "default"

def load_config() -> SimulationConfig:
    try:
        with open("cnsim-profile.yaml", "r") as f:
            raw = yaml.safe_load(f)

        profiles_raw = raw["profiles"]
        profiles = {}
        for name, data in profiles_raw.items():
            plmn = Plmn(**data["plmn"])
            slice_info = Slice(**data["slice"])
            profiles[name] = SimulationProfile(
                plmn=plmn,
                dnn=data["dnn"],
                slice=slice_info,
                numUe=data["numUe"],
                gNBs=data["gNBs"],
                rate=data["rate"]
            )

        return SimulationConfig(profiles=profiles, current_profile="default")
    except Exception as e:
        typer.echo(f"failed to load cnsim-profile.yaml: {e}")
        raise typer.Exit()




def send_request(method: str, endpoint: str, body: dict = None):
    url = f"{BASE_URL}{endpoint}"
    try:
        response = client.request(method, url, json=body)
        print(f"{method} {url} → {response.status_code}")
        print(response.text)
    except httpx.RequestError as e:
        print(f"request failed: {e}")

class SimCtlShell(cmd.Cmd):
    def __init__(self, config: SimulationConfig):
        super().__init__()
        self.config = config

    def get_current_profile(self):
        return self.config.profiles[self.config.current_profile]

    def do_showconfig(self, arg):
        "Show the active simulation profile"
        profile = self.get_current_profile()
        name = self.config.current_profile
        print(f"active profile: '{name}'")
        print(f"  plmn: {profile.plmn.mcc}-{profile.plmn.mnc}")
        print(f"  dnn: {profile.dnn}")
        print(f"  slice: SST={profile.slice.sst}, SD={profile.slice.sd}")
        print(f"  number of UEs: {profile.numUe}")
        print(f"  gNBs: {profile.gNBs}")
        print(f"  rate: {profile.rate}")


    def do_listprofiles(self, arg):
        "List all available profiles"
        print("Available profiles:")
        for name in self.config.profiles:
            mark = "← current" if name == self.config.current_profile else ""
            print(f"  - {name} {mark}")

    def do_switch(self, profile_name):
        "Switch to a different simulation profile: switch <name>"
        if profile_name in self.config.profiles:
            self.config.current_profile = profile_name
            print(f"✅ Switched to profile '{profile_name}'")
        else:
            print(f"profile '{profile_name}' not found. Use 'listprofiles' to view available ones.")

    intro = "Welcome to simctl. Type help or ? to list commands.\n"
    prompt = "simctl > "

    def do_init(self, arg):
        "Send POST /configure to configure the simulation"
        profile = self.get_current_profile()
        data = {}
        data["numOfUe"] = profile.numUe
        data["arrivalRate"] = profile.rate
        data["plmn"] =  {"mcc": profile.plmn.mcc, "mnc": profile.plmn.mnc}
        data["numOfgNB"] = profile.gNBs
        data["slice"] = {"sst": profile.slice.sst, "sd": profile.slice.sd}
        data["dnn"] = profile.dnn
        send_request("POST", "/configure", data)

    def do_start(self, arg):
        "Send POST /start to start the simulation"
        send_request("POST", "/start")

    def do_stop(self, arg):
        "Send POST /stop to stop the simulation"
        send_request("POST", "/stop")

    def do_status(self, arg):
        "Send GET /status to check current simulation status"
        send_request("GET", "/status")

    def do_loop(self, arg):
        "Continuously check /status every 3 seconds"
        print("Looping status check (Ctrl+C to stop)...")
        try:
            while True:
                self.do_status(arg)
                time.sleep(3)
        except KeyboardInterrupt:
            print("\nStopped status loop.")

    def do_exit(self, arg):
        "Exit the shell"
        print("Bye!")
        return True

    def do_quit(self, arg):
        "Exit the shell"
        return self.do_exit(arg)

    def emptyline(self):
        pass

    def default(self, line):
        print(f"Unknown command: {line}. Type 'help'.")

def main():
    config = load_config()
    SimCtlShell(config).cmdloop()

if __name__ == "__main__":
    typer.run(main)