import SwiftUI

struct SettingsView: View {

    private let appVersion = "1.0.0"

    var body: some View {
        List {
            Section(header: Text("Preferences")) {
                settingRow(
                    title: "Notifications",
                    systemImage: "bell",
                    detail: "Managed by system settings"
                )
                settingRow(
                    title: "Privacy",
                    systemImage: "hand.raised",
                    detail: "Review app policies"
                )
            }

            Section(header: Text("About")) {
                settingRow(
                    title: "Version",
                    systemImage: "info.circle",
                    detail: appVersion
                )
                settingRow(
                    title: "Build",
                    systemImage: "number",
                    detail: "Template"
                )
            }
        }
#if os(iOS)
        .listStyle(.insetGrouped)
#endif
        .navigationTitle("Settings")
    }

    @ViewBuilder
    private func settingRow(title: String, systemImage: String, detail: String) -> some View {
        HStack {
            Label(title, systemImage: systemImage)
            Spacer()
            Text(detail)
                .foregroundStyle(.secondary)
                .multilineTextAlignment(.trailing)
        }
    }
}
