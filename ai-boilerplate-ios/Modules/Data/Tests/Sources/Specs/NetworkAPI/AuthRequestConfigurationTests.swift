import Testing

@testable import Data

@Suite("AuthRequestConfiguration")
struct AuthRequestConfigurationTests {

    @Test("uses the API base URL from app configuration")
    func usesTheAPIBaseURLFromAppConfiguration() {
        let baseURL = AppConfiguration.apiBaseURL(from: [
            "API_BASE_URL": "https://api.fzflabs.test/api"
        ])

        #expect(baseURL == "https://api.fzflabs.test/api")
    }

    @Test("uses the local API base URL default")
    func usesTheLocalAPIBaseURLDefault() {
        let configuration = AuthRequestConfiguration.logIn(username: "demo", password: "secret")

        #expect(configuration.baseURL == "http://127.0.0.1:8000/api")
    }
}
