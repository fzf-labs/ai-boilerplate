# AIBoilerplate

`ai-boilerplate-ios` is the ready-to-use iOS starter project for this
repository. It is generated from
[nimblehq/ios-templates](https://github.com/nimblehq/ios-templates), then kept
as a real project in this directory instead of a template wrapper.

## Stack

- SwiftUI
- Swift 6.1+
- Xcode 26+
- Tuist
- Swift Package Manager
- Ruby 3.2+ for Fastlane and related automation
- mise for local tool version management

## Getting Started

From this directory:

```bash
mise install
bundle install
bundle exec arkana
tuist generate
```

Then open the generated Xcode project/workspace and build with Xcode.

## Project Identity

- Project name: `AIBoilerplate`
- Bundle identifier: `com.fzflabs.aiboilerplate`

Update the bundle identifier in the `.xcconfig`/Tuist settings before shipping a
real app.

## Upstream

The project is based on `nimblehq/ios-templates`. See `UPSTREAM.md` and
`LICENSE-NIMBLEHQ` for source and license details.
