import SwiftUI

struct SignOutView: View {

    var onContinue: () -> Void = {}

    var body: some View {
        VStack(spacing: 20.0) {
            Image(systemName: "person.crop.circle.badge.exclamationmark")
                .font(.system(size: 48))
                .foregroundColor(.accentColor)
            Text("Signed Out")
                .font(.title2.bold())
            Text("You are signed out. Continue with the local demo session to preview the signed-in flow.")
                .multilineTextAlignment(.center)
                .foregroundStyle(.secondary)

            Button("Continue with Demo Session", action: onContinue)
                .buttonStyle(.borderedProminent)
        }
        .padding()
    }
}
