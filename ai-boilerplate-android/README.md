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

The sample API uses `https://dummyjson.com/` by default. To point the app at
another API, create a local `local.properties` file and set:

```properties
BASE_URL=https://example.com/
```

## Project Identity

- App name: `AI Boilerplate Android`
- Package/application ID: `com.fzflabs.aiboilerplate.android`
- Root Gradle project: `AI Boilerplate Android`

Update the application ID and API base URL before shipping a real app.

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
