go_binary(
    name = "chessengine",
    srcs = glob(
        ["*.go"],
        exclude = [
            "*_test.go",
        ],
    ),
    deps = [
        "//src/position",
    ],
)

go_test(
    name = "chessengine_test",
    srcs = glob(
        ["*_test.go"],
    ),
    deps = [
        ":chessengine",
    ],
)
