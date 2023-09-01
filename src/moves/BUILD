go_library(
    name = "moves",
    srcs = [
        "move_list.go",
        "moves.go",
    ],
    visibility = [
        "//src/engine:all",
        "//src/pieces:all",
        "//src/position:all",
    ],
    deps = [
        "//src/board",
        "//src/piece",
    ],
)

go_test(
    name = "moves_test",
    srcs = [
        "move_list_internal_test.go",
        "moves_internal_test.go",
    ],
    deps = [
        ":moves",
        "//src/board",
        "//src/piece",
    ],
)
