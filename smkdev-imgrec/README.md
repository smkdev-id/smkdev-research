# MobileImgRec

Real Time Image Recognition with MobileNet in Flutter Applications.

## Rules
- minSdkVersion: 26
- Gradle Version: gradle-8.4

## Tools
- Flutter SDK
- Android Studio
- Visual Studio Code
- Terminal

## Intialize Environment
- Specify [Development Platform for Flutter](https://docs.flutter.dev/get-started/install), based on your Operating System
- Choose type of app: Android
- Install Android Studio for Android SDK
- Follow Hardware requirements, Software requirements, Operating system, and more based on Operating System
- Install the Flutter SDK
  - Install Flutter Extension for Visual Studio Code
    - press `CTRL + SHIFT + P` for open Command Pallete
    - type `flutter`, and select `Flutter: New Project`
    - click `Download SDK`, so Flutter SDK will installed
    - `Do you want to add the Flutter SDK to PATH so it's accessible in external terminals?`, type/click Yes
    - For Linux, add `export PATH="$PATH:$HOME/flutter/bin"` or based on where your SDK are downloaded.
    - So, we can use CLI Utilities.
  - Manual Install
    - Download latest stable SDK [Here](https://storage.googleapis.com/flutter_infra_release/releases/stable/windows/flutter_windows_3.19.6-stable.zip)
    - You can extract this .zip file on preferable folder that can be found easily.s
    - Add tou your PATH, so you can use CLI utilities.
- Configure Android development
  - You can download Android Studio, so you can easily setup Android Device Target, either Virtual or Physical Device(s).
  - Additionaly, you can check `SDK Manager` if you want other SDK Installations.
- Check Development Setup
  - you can use `flutter doctor`, the result will be like this (depend on your environment)
  ```bash
    Doctor summary (to see all details, run flutter doctor -v):
    [✓] Flutter (Channel stable, 3.19.6, on Ubuntu 20.04.6 LTS 5.15.0-105-generic,
        locale en_US.UTF-8)
    [✓] Android toolchain - develop for Android devices (Android SDK version 34.0.0)
    [✓] Chrome - develop for the web
    [✗] Linux toolchain - develop for Linux desktop
        ✗ clang++ is required for Linux development.
          It is likely available from your distribution (e.g.: apt install clang),
          or can be downloaded from https://releases.llvm.org/
        ✗ ninja is required for Linux development.
          It is likely available from your distribution (e.g.: apt install
          ninja-build), or can be downloaded from
          https://github.com/ninja-build/ninja/releases
        ✗ GTK 3.0 development libraries are required for Linux development.
          They are likely available from your distribution (e.g.: apt install
          libgtk-3-dev)
    [✓] Android Studio (version 2023.2)
    [✓] VS Code (version 1.88.1)
    [✓] Connected device (3 available)
    [✓] Network resources

    ! Doctor found issues in 1 category.
  ```

## Run the Project
- Clean the past build with `flutter clean`
- on Visual Studio Code, you can click `Start Debugging`
- The Source will compiled into your target device, based on Doctor connected device(s).