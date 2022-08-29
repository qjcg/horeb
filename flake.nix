{
  description = "A Nix flake for horeb.";

  inputs = {
    devshell.url = "github:numtide/devshell";
    flake-utils.url = "github:numtide/flake-utils";
    pkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = {self, ...} @ inputs:
    inputs.flake-utils.lib.eachDefaultSystem (system: let
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

      overlays = {
        default = final: prev: {
          horeb = self.packages.${prev.system}.default;
        };
      };

      packages = let
        inherit (pkgs) buildGo118Module protobuf protoc-gen-go protoc-gen-go-grpc;
        inherit (pkgs.lib) fakeSha256 licenses;

        version = "0.14.0";
      in {
        default = buildGo118Module {
          inherit version;

          pname = "horeb";
          src = self;
          packages = ["cmd/..."];
          vendorSha256 = "sha256-wvmz1jzRxPCldS/1VHdPoT4hNSSoPTEEYezjDCjRqMw=";

          nativeBuildInputs = [protobuf protoc-gen-go protoc-gen-go-grpc];

          # See https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code
          preBuild = ''
            ${protobuf}/bin/protoc -I $src/proto/ $src/proto/horeb.proto --go_opt=paths=source_relative --go_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=.
          '';

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
