"""Dependencies of Tink ObjC."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def tink_objc_deps():
    """Dependencies for Tink Objective C."""
    if not native.existing_rule("build_bazel_rules_apple"):
        # Release from 2022-12-21.
        http_archive(
            name = "build_bazel_rules_apple",
            sha256 = "9e26307516c4d5f2ad4aee90ac01eb8cd31f9b8d6ea93619fc64b3cbc81b0944",
            url = "https://github.com/bazelbuild/rules_apple/releases/download/2.2.0/rules_apple.2.2.0.tar.gz",
        )
    if not native.existing_rule("build_bazel_rules_swift"):
        # Release from 2022-09-16.
        http_archive(
            name = "build_bazel_rules_swift",
            sha256 = "bf2861de6bf75115288468f340b0c4609cc99cc1ccc7668f0f71adfd853eedb3",
            url = "https://github.com/bazelbuild/rules_swift/releases/download/1.7.1/rules_swift.1.7.1.tar.gz",
        )
    if not native.existing_rule("build_bazel_apple_support"):
        # Release from 2022-10-31.
        http_archive(
            name = "build_bazel_apple_support",
            sha256 = "9a2338d6f8dad3244f823f2dc6084a03e4d0fbb27ca892dc970e6890d5b48184",
            url = "https://github.com/bazelbuild/apple_support/releases/download/1.5.0/apple_support.1.5.0.tar.gz",
        )

    # Currently required by ios_unit_test
    if not native.existing_rule("xctestrunner"):
        # Release from 2021-11-01.
        http_archive(
            name = "xctestrunner",
            strip_prefix = "xctestrunner-0.2.15",
            sha256 = "03ce1088f74d85e23d14a09e533383bd06368d2b453c962e6ce66f80b833feae",
            url = "https://github.com/google/xctestrunner/archive/refs/tags/0.2.15.zip",
        )

    # Subpar is a utility for creating self-contained python executables. It is designed to work well with Bazel.
    # Currently required by @xctestrunner
    if not native.existing_rule("subpar"):
        # Release from 2019-05-14.
        http_archive(
            name = "subpar",
            strip_prefix = "subpar-2.0.0",
            sha256 = "8876244a984d75f28b1c64d711b6e5dfab5f992a3b741480e63cfc5e26acba93",
            url = "https://github.com/google/subpar/archive/refs/tags/2.0.0.zip",
        )
