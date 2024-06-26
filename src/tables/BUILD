go_library(
    name = "tables",
    srcs = ["tables.go"],
    deps = [
        "//src/bitboard",
        "//src/colours",
    ],
    visibility = ["//cmd:all"],
)

go_test(
    name = "tables_test",
    srcs = ["tables_internal_test.go"],
    deps = [
        ":tables",
        "//src/squares",
        "//src/bitboard",
        "//src/colours",
    ],
)
