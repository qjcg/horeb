{
  description = "A Nix flake for horeb.";

  inputs = {
    devshell = {
      url = "github:numtide/devshell";
      inputs.nixpkgs.follows = "pkgs";
    };
    flake-utils.url = "github:numtide/flake-utils";
    pkgs.url = "github:nixos/nixpkgs/release-22.05";
  };

  outputs = {self, ...} @ inputs:
    (inputs.flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import inputs.pkgs {
        inherit system;
        overlays = [inputs.devshell.overlay];
      };
    in {
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

      devShells = {
        default = pkgs.devshell.fromTOML ./devshell.toml;
      };

      packages = let
        inherit (pkgs) buildGoModule;
        inherit (pkgs.lib) fakeSha256 licenses;

        pname = "horeb";
        version = "0.14.0";
      in {
        default = buildGoModule {
          inherit pname version;

          src = self;
          packages = ["cmd/..."];
          vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";

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
    }))
    # Avoids nix flake check error: overlay does not take an argument named 'final'
    # See e.g. https://github.com/ivanovs-4/haskell-flake-utils/issues/2
    // {
      overlays = {
        default = final: prev: {
          horeb = self.packages.${prev.system}.default;
        };
      };
    };
}
