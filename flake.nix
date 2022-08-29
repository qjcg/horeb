{
  description = "A Nix flake for horeb.";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    pkgs.url = "github:nixos/nixpkgs/nixos-22.05";
  };

  outputs = {self, ...} @ inputs:
    inputs.flake-utils.lib.eachDefaultSystem (system: {
      apps = with self.packages.${system}; {
        default = self.apps.${system}.horeb;

        horeb = {
          program = "${default}/bin/horeb";
          type = "app";
        };
      };

      checks = {
        build = self.packages.${system}.default;
      };

      overlays = {
        default = final: prev: {
          horeb = self.packages.${prev.system}.default;
        };
      };

      packages = let
        inherit (pkgs) buildGo118Module;
        inherit (pkgs.lib) fakeSha256 licenses;

        pkgs = import inputs.pkgs {inherit system;};
        version = "0.12.1";
      in {
        default = buildGo118Module {
          inherit version;

          pname = "horeb";
          src = self;
          packages = ["cmd/..."];
          vendorSha256 = "sha256-wvmz1jzRxPCldS/1VHdPoT4hNSSoPTEEYezjDCjRqMw=";

          ldflags = [
            "-s"
            "-w"
            "-X github.com/qjcg/horeb/pkg/horeb.Version=${version}"
          ];

          meta = {
            description = "Speaking in tongues via stdout.";
            homepage = "https://github.com/qjcg/horeb";
            license = licenses.mit;
          };
        };
      };
    });
}
