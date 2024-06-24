go_library(
    name = "engine",
    srcs = ["engine.go"],
    visibility = [
        "//src:all",
        "//cmd:all",
    ],
    deps = [
        "//src/eval",
        "//src/moves",
        "//src/position",
    ],
)

go_test(
    name = "engine_test",
    srcs = ["engine_internal_test.go"],
    deps = [":engine"],
)
