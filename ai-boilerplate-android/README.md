# AI Boilerplate Android

`ai-boilerplate-android` is the native Android starter project for this
repository. It is based on
[its-me-debk007/kotlin-android-mvvm-template](https://github.com/its-me-debk007/kotlin-android-mvvm-template)
and kept as a real project in this directory instead of a template wrapper.

## Stack

- Kotlin 2.x
- Android Gradle Plugin
- Jetpack Compose and Material 3
- Compose Navigation
- MVVM architecture
- Dagger Hilt
- Retrofit with kotlinx.serialization
- Coil
- Gradle Version Catalog
- detekt

## Getting Started

From this directory:

```bash
./gradlew detekt test assembleDebug -x validateSigningDebug
```

The app points at the local AI Boilerplate backend by default. Android
emulators reach the host machine through `10.0.2.2`, so the built-in default is
`http://10.0.2.2:8000/api/`. To point the app at another API, create a local
`local.properties` file and set the host that is reachable from the emulator or
device. For example, for a backend running on another machine in your LAN:

```properties
BASE_URL=http://192.168.1.20:8000/api/
```

## Project Identity

- App name: `AI Boilerplate Android`
- Package/application ID: `com.fzflabs.aiboilerplate.android`
- Root Gradle project: `AI Boilerplate Android`

Update the application ID and API base URL before shipping a real app.

## Verification And Shipping Notes

Use this command for local verification:

```bash
./gradlew detekt test assembleDebug -x validateSigningDebug
```

Keep keystores, signing passwords, Firebase service files, production package
IDs, and production API hosts outside the repository. Use `local.properties` or
platform secrets for machine-specific values.

## Project Structure

The project includes one Android application module, `app`, with these main
packages:

- `di`: Hilt dependency injection modules.
- `network`: Retrofit API service definitions.
- `model`: kotlinx.serialization response models.
- `repository`: Repository interfaces and implementations.
- `presentation`: Compose UI, theme, and ViewModels.
- `util`: Shared state helpers.

## Upstream

The project is based on `its-me-debk007/kotlin-android-mvvm-template`. See
`UPSTREAM.md` and `LICENSE-ITS-ME-DEBK007` for source and license details.

## More Documentation

- Root operations workflow: `../docs/operations.md`
- Template selection and tradeoffs: `../docs/technical-decisions.md`
- Environment configuration: `../docs/environment.md`
- Release checklist: `../docs/release.md`
- Troubleshooting: `../docs/troubleshooting.md`
- Verification matrix: `../docs/verification.md`
