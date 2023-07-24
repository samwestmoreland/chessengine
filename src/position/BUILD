go_library(
    name = "position",
    srcs = glob(["*.go"]),
    visibility = ["PUBLIC"],
    deps = ["//src/board"],
)

go_test(
    name = "position_test",
    srcs = glob(["*_test.go"]),
    deps = [":position"],
)
