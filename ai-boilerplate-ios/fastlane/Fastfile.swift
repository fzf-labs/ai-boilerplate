// This file contains the fastlane.tools configuration
// You can find the documentation at https://docs.fastlane.tools
//
// For a list of all available actions, check out
//
//     https://docs.fastlane.tools/actions
//

import Foundation

class Fastfile: LaneFile {

    // MARK: - Code signing

    func syncDevelopmentStagingCodeSigningLane() {
        desc("Sync the Development match signing for the Staging build")
        Match.syncCodeSigning(
            type: .development,
            environment: .staging
        )
    }

    func syncDevelopmentProductionCodeSigningLane() {
        desc("Sync the Development match signing for the Production build")
        Match.syncCodeSigning(
            type: .development,
            environment: .production
        )
    }
    
    func syncAdHocDevCodeSigningLane() {
        desc("Sync the Ad Hoc match signing for the Dev build")
        Match.syncCodeSigning(
            type: .adHoc,
            environment: .dev
        )
    }

    func syncAdHocStagingCodeSigningLane() {
        desc("Sync the Ad Hoc match signing for the Staging build")
        Match.syncCodeSigning(
            type: .adHoc,
            environment: .staging
        )
    }

    func syncAdHocProductionCodeSigningLane() {
        desc("Sync the Ad Hoc match signing for the Production build")
        Match.syncCodeSigning(
            type: .adHoc,
            environment: .production
        )
    }

    func syncAppStoreCodeSigningLane() {
        desc("Sync the App Store match signing for the Production build")
        Match.syncCodeSigning(
            type: .appStore,
            environment: .production
        )
    }

    func removeKeychainLane() {
        desc("Delete keychain")
        Keychain.remove()
    }

    // MARK: - Build
    
    func buildAdHocDevLane() {
        desc("Build ad-hoc dev")
        Build.adHoc(environment: .dev)
    }


    func buildAdHocStagingLane() {
        desc("Build ad-hoc staging")
        Build.adHoc(environment: .staging)
    }

    func buildAdHocProductionLane() {
        desc("Build ad-hoc production")
        Build.adHoc(environment: .production)
    }

    func buildAppStoreLane() {
        desc("Build app store")
        Build.appStore()
    }

    // MARK: - Upload builds to Firebase and AppStore

    private func firebaseReleaseNotes() -> String {
        let releaseNote = Constant.releaseNote.trimmingCharacters(in: .whitespacesAndNewlines)
        return releaseNote.isEmpty ? "Release notes not provided." : releaseNote
    }
    
    func buildDevAndUploadToFirebaseLane() {
        desc("Build Dev app and upload to Firebase")

        setAppVersion()
        bumpBuild()

        buildAdHocDevLane()

        Distribution.uploadToFirebase(environment: .dev, releaseNotes: firebaseReleaseNotes())

        Build.saveBuildContextToCI()
    }

    func buildStagingAndUploadToFirebaseLane() {
        desc("Build Staging app and upload to Firebase")

        setAppVersion()
        bumpBuild()

        buildAdHocStagingLane()

        Distribution.uploadToFirebase(environment: .staging, releaseNotes: firebaseReleaseNotes())

        Build.saveBuildContextToCI()
    }

    func buildProductionAndUploadToFirebaseLane() {
        desc("Build Production app and upload to Firebase")

        setAppVersion()
        bumpBuild()

        buildAdHocProductionLane()

        Distribution.uploadToFirebase(environment: .production, releaseNotes: firebaseReleaseNotes())

        Build.saveBuildContextToCI()
    }

    func buildAndUploadToAppStoreLane() {
        desc("Build Production app and upload to App Store")

        setAppVersion()
        AppStoreAuthentication.connectAPIKey()
        if Secret.bumpAppStoreBuildNumber {
            bumpAppstoreBuild()
        } else {
            bumpBuild()
        }

        buildAppStoreLane()

        Distribution.uploadToAppStore()

        Build.saveBuildContextToCI()
    }

    func buildAndUploadToTestFlightLane() {
        desc("Build Production app and upload to TestFlight")

        setAppVersion()
        bumpBuild()

        buildAppStoreLane()

        AppStoreAuthentication.connectAPIKey()
        Distribution.uploadToTestFlight(changeLog: firebaseReleaseNotes())

        Build.saveBuildContextToCI()
    }

    // MARK: - Test

    func buildAndTestLane() {
        desc("Build and Test project")
        guard testTargetCount() > 0 else {
            echo(message: "🚨 Nothing to test")
            return
        }
        Test.buildAndTest(
            environment: .staging,
            devices: Constant.devices
        )
    }

    func buildAndTestDevLane() {
        desc("Build and Test dev project")
        guard testTargetCount() > 0 else {
            echo(message: "🚨 Nothing to test")
            return
        }
        Test.buildAndTest(
            environment: .dev,
            devices: Constant.devices
        )
    }

    func setUpTestProjectLane() {
        desc("Disable Exempt Encryption")
        Test.disableExemptEncryption()
    }

    // MARK: - Register device

    func registerNewDeviceLane() {
        let deviceName = prompt(text: "Enter the device name:")
        let deviceUDID = prompt(text: "Enter the device UDID:")

        registerDevice(
            name: deviceName,
            udid: deviceUDID,
            apiKey: .userDefined(Constant.apiKey),
            teamId: .userDefined(Constant.appleStagingTeamId)
        )

        Match.syncCodeSigning(type: .development, environment: .staging, isForce: true)
        Match.syncCodeSigning(type: .adHoc, environment: .staging, isForce: true)
    }

    func addDevicesGenerateProfilesLane() {
        desc("Add device and regenerate profiles with match")

        guard let data = Secret.devices.data(using: .utf8),
              let devices = try? JSONSerialization.jsonObject(with: data, options: []) as? [String: Any] else {
            return
        }
        registerDevices(
            devices: .userDefined(devices),
            apiKey: .userDefined(Constant.apiKey),
            teamId: .userDefined(Constant.appleStagingTeamId),
            platform: Secret.platform
        )

        Match.syncCodeSigning(type: .development, environment: .staging, isForce: true)
        Match.syncCodeSigning(type: .adHoc, environment: .staging, isForce: true)
    }

    // MARK: - Utilities

    func cleanUpOutputLane() {
        desc("Clean up Output")
        clearDerivedData(derivedDataPath: Constant.outputPath)
    }

    // MARK: - Private Helper

    private func setAppVersion() {
        desc("Check if any specific version number in build environment")
        guard !Constant.manualVersion.isEmpty else { return }
        incrementVersionNumber(
            versionNumber: .userDefined(Constant.manualVersion)
        )
    }

    private func bumpBuild(buildNumber: Int? = nil) {
        let ciBuildNumber = Build.ciBuildNumber
        let resolvedBuildNumber = buildNumber ?? ciBuildNumber ?? numberOfCommits()
        let buildSource: String

        if buildNumber != nil {
            buildSource = "provided build number"
        } else if ciBuildNumber != nil {
            buildSource = "CI build number"
        } else {
            buildSource = "number of commits"
        }

        desc("Set build number with \(buildSource)")
        incrementBuildNumber(
            buildNumber: .userDefined(String(resolvedBuildNumber)),
            xcodeproj: .userDefined(Constant.projectPath)
        )
    }

    private func bumpAppstoreBuild() {
        desc("Set build number with App Store latest build")
        let theLatestBuildNumber = latestTestflightBuildNumber(
            appIdentifier: Constant.productionBundleId
        ) + 1
        incrementBuildNumber(
            buildNumber: .userDefined("\(theLatestBuildNumber)")
        )
    }

    private func testTargetCount(in testPlanPath: String = "\(Constant.projectName).xctestplan") -> Int {
        guard let data = FileManager.default.contents(atPath: testPlanPath),
              let json = try? JSONSerialization.jsonObject(with: data) as? [String: Any],
              let testTargets = json["testTargets"] as? [Any] else {
            return 0
        }
        return testTargets.count
    }
}
