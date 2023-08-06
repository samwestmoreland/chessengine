go_binary(
    name = "chessengine",
    srcs = glob(
        ["*.go"],
        exclude = [
            "*_test.go",
        ],
    ),
    deps = [
        "///third_party/go/github.com_integrii_flaggy//:flaggy",
        "///third_party/go/github.com_sirupsen_logrus//:logrus",
        "//src/position",
    ],
)
