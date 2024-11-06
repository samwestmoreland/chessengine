go_library(
    name = "bitboard",
    srcs = ["bitboard.go"],
    visibility = [
        "//src/tables:all",
        "//cmd:all",
        "//cmd/magics:all",
    ],
)

go_test(
    name = "bitboard_test",
    srcs = ["bitboard_internal_test.go"],
    deps = [
        ":bitboard",
        "//src/squares",
    ],
)
