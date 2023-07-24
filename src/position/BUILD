go_library(
    name = "position",
    srcs = glob(["*.go"]),
    visibility = ["PUBLIC"],
    deps = [
        "//src/board",
        "//src/pieces",
    ],
)

go_test(
    name = "position_test",
    srcs = glob(["*_test.go"]),
    deps = [":position"],
)
