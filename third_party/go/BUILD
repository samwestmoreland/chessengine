subinclude("///go//build_defs:go")

package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    install_std = True,
    version = "1.20.7",
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

go_repo(
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
)

go_repo(
    module = "github.com/stretchr/objx",
    version = "v0.1.0",
)

go_repo(
    module = "gopkg.in/yaml.v3",
    version = "v3.0.0-20200313102051-9f266ea9e77c",
)

go_repo(
    module = "gopkg.in/check.v1",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repo(
    module = "golang.org/x/sys",
    version = "v0.0.0-20220715151400-c0bba94af5f8",
)

go_repo(
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
)

go_repo(
    module = "github.com/stretchr/testify",
    version = "v1.7.0",
)

go_repo(
    module = "github.com/sirupsen/logrus",
    version = "v1.9.3",
)

remote_file(
    name = "golangci-lint_install_script",
    url = "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh",
    binary = True,
    out = "install.sh",
)

genrule(
    name = "golangci-lint",
    srcs = [":golangci-lint_install_script"],
    outs = ["golangci-lint"],
    cmd = "cat $(location :golangci-lint_install_script) | sh -s -- -b . v1.54.2",
    binary = True,
)
