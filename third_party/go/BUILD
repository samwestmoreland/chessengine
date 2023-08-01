subinclude("///go//build_defs:go")

package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    install_std = True,
    version = "1.20.6",
)

go_repo(
    module = "github.com/integrii/flaggy",
    version = "v1.5.2",
)

go_repo(
    module = "github.com/google/go-cmp",
    version = "v0.5.6",
)

go_repo(
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20191204190536-9bdfabe68543",
)
