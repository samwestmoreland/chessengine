go_library(
    name = "tables",
    srcs = [
        "constants.go",
        "knight_tables.go",
        "pawn_tables.go",
        "king_tables.go",
        "bishop_tables.go",
        "rook_tables.go",
        "tables.go",
    ],
    deps = [
        "//src/bitboard",
        "//src/colours",
        "//magic:magic_data",
    ],
    visibility = [
        "//cmd:all",
        "//cmd/magics:all",
        "//cmd/uci:all",
    ],
)

go_test(
    name = "tables_test",
    srcs = [
        "pawn_tables_internal_test.go",
        "knight_tables_internal_test.go",
        "king_tables_internal_test.go",
        "bishop_tables_internal_test.go",
        "rook_tables_internal_test.go",
    ],
    deps = [
        ":tables",
        "//src/squares",
        "//src/bitboard",
        "//src/colours",
        "//magic:magic_data",
    ],
)
