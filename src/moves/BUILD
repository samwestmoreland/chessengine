go_library(
    name = "moves",
    srcs = [
        "move_list.go",
        "moves.go",
    ],
    visibility = [
        "//plz-out/gen/src/position:all",
        "//plz-out/tmp/src/position/_position_test#lib._build/src/position:all",
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
