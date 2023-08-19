go_library(
    name = "position",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = ["PUBLIC"],
    deps = [
        "///third_party/go/github.com_sirupsen_logrus//:logrus",
        "//src/board",
        "//src/moves",
        "//src/piece",
    ],
)

go_test(
    name = "position_test",
    srcs = glob(["*_test.go"]),
    deps = [
        ":position",
        "//src/board",
        "//src/moves",
        "//src/piece",
    ],
)

go_benchmark(
    name = "position_benchmark",
    srcs = glob(["*_benchmark.go"]),
    deps = [
        ":position",
        "//src/board",
    ],
)
