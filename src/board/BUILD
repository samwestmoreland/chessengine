go_library(
    name = "board",
    srcs = glob(["*.go"]),
    visibility = ["PUBLIC"],
    deps = ["///third_party/go/github.com_stretchr_testify//assert"],
)

go_test(
    name = "board_test",
    srcs = glob(["*_test.go"]),
    visibility = ["PUBLIC"],
    deps = [
        ":board",
        "///third_party/go/github.com_stretchr_testify//assert",
    ],
)
