go_binary(
    name = "chessengine",
    srcs = glob(
        ["*.go"],
        exclude = [
            "*_test.go",
        ],
    ),
    deps = [
        "//src/engine",
        "//src/position",
    ],
)
