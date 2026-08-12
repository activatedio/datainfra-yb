with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "datainfra-yb";
  buildInputs = with pkgs; [
    zstd
    go
    gnumake
  ];
  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
  '';
  hardeningDisable = [ "fortify" ];
}
