go_library(
    name = "position",
    srcs = glob(["*.go"]),
    visibility = ["PUBLIC"],
    deps = [
        "//src/board",
        "//src/moves",
    ],
)

go_test(
    name = "position_test",
    srcs = glob(["*_test.go"]),
    deps = [
        ":position",
        "//src/board",
    ],
)
