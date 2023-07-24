go_library(
    name = "board",
    srcs = glob(["*.go"]),
    visibility = ["PUBLIC"],
)

go_test(
    name = "board_test",
    srcs = glob(["*_test.go"]),
    visibility = ["PUBLIC"],
    deps = [":board"],
)
