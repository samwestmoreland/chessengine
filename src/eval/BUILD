go_library(
    name = "eval",
    srcs = [
        "eval.go",
        "shannon.go",
    ],
    visibility = ["//src/engine:all"],
    deps = [
        "//src/board",
        "//src/piece",
        "//src/position",
    ],
)

go_test(
    name = "eval_test",
    srcs = ["shannon_internal_test.go"],
    deps = [
        ":eval",
        "//src/position",
    ],
)
