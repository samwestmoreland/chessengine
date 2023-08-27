go_library(
    name = "engine",
    srcs = ["engine.go"],
    deps = ["//src/position"],
)

go_test(
    name = "engine_test",
    srcs = ["engine_internal_test.go"],
    deps = [":engine"],
)
