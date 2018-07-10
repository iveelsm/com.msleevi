load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/msleevi/com.msleevi",
)

go_library(
    name = "api_library",
    deps = [
        "github.com/gorilla/mux",
        "github.com/olivere/elastic",
        "github.com/google/uuid",
        "github.com/sirupsen/logrus",
        "github.com/go-kivik/couchdb",
        "golang.org/x/crypto/bcrypt",
    ],
    importpath = "github.com/msleevi/com.msleevi",
    visibility = ["//visibility:public"],
)


go_binary(
    name = "main",
    srcs = ["src/api/main.go"],
)