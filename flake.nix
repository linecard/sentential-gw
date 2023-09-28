{
  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
  let
    system = "aarch64-linux";
    pkgs = nixpkgs.legacyPackages.${system};

    sententialGw = pkgs.stdenv.mkDerivation {
      __noChroot = true;
      name = "sentential-gw";
      src = ./src;
      buildInputs = [ pkgs.go_1_21 ];

      buildPhase = ''
        export HOME=$TMPDIR
        mkdir -p $out/bin
        go build -o $out/bin/sntl-gw
      '';
    };

    image = pkgs.dockerTools.buildLayeredImage {
      name = "sentential-gw";
      tag = "latest";
      contents = [
        sententialGw
      ];
      config = {
        WorkingDir = "/src";
        Entrypoint = [ "/bin/sntl-gw" ];
      };
    };
  in
  {
    packages.${system} = {
      default = image;
    };
  };
}