go_library(
    name = "tables",
    srcs = [
        "constants.go",
        "knight_tables.go",
        "pawn_tables.go",
        "king_tables.go",
    ],
    deps = [
        "//src/bitboard",
        "//src/colours",
    ],
    visibility = ["//cmd:all"],
)

go_test(
    name = "tables_test",
    srcs = [
        "pawn_tables_internal_test.go",
        "knight_tables_internal_test.go",
        "king_tables_internal_test.go",
    ],
    deps = [
        ":tables",
        "//src/squares",
        "//src/bitboard",
        "//src/colours",
    ],
)
