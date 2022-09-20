let pkgs = import <nixpkgs> {};
in pkgs.mkShell {
  name = "golang-dev";
  buildInputs = with pkgs; [
    go_1_18 gopls go-protobuf protobuf delve
  ];
}
