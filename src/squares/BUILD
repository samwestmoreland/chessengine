go_library(
    name = "squares",
    srcs = ["squares.go"],
    visibility = [
        "//src/bitboard:all",
        "//src/tables:all",
    ],
)

go_test(
    name = "squares_test",
    external = True,
    srcs = ["squares_test.go"],
    deps = [":squares"],
)
