{
  inputs= {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    sntlGw.url = "github:linecard/sentential-gw/flake";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, sntlGw, ... }:
  let system = flake-utils.lib.system;
  in flake-utils.lib.eachSystem [
    "x86_64-linux"
    "aarch64-linux"
  ] (system:
    let
      pkgs = nixpkgs.legacyPackages.${system};

      build-sntlGw = (pkgs:
        # Build the go binary for each target
        let
          binary = pkgs.stdenv.mkDerivation {
            __noChroot = true;
            name = "sentential-gw";
            src = sntlGw;
            buildInputs = [ pkgs.go_1_21 ];

            buildPhase = ''
              cd src
              export HOME=$TMPDIR
              mkdir -p $out/bin

              go build -o $out/bin/sntl-gw
            '';
          };
        in
          pkgs.dockerTools.buildLayeredImage {
            name = "sentential-gw";
            tag = "latest";
            contents = [
              binary
            ];
            config = {
              WorkingDir = "/src";
              Entrypoint = [ "/bin/sntl-gw" ];
            };
          }
      );
  in rec {
    packages = rec {
      sntlGw = build-sntlGw pkgs;
      sntlGw-cross-aarch64-linux =
        build-sntlGw pkgs.pkgsCross.aarch64-multiplatform;
    };

    defaultPackage = packages.sntlGw;
  });
}