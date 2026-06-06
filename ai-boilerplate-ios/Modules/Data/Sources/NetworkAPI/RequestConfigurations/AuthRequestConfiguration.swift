//
//  AuthRequestConfiguration.swift
//

import Alamofire
import Foundation

/// Auth endpoints used by login, logout, and token refresh requests.
enum AuthRequestConfiguration: RequestConfiguration {

    case logIn(username: String, password: String)
    case logOut(accessToken: String)
    case refreshToken(String)

    var baseURL: String { AppConfiguration.apiBaseURL() }

    var endpoint: String {
        switch self {
        case .logIn:
            "/auth/login"
        case .logOut:
            "/auth/logout"
        case .refreshToken:
            "/auth/refresh"
        }
    }

    var method: HTTPMethod { .post }

    var encoding: ParameterEncoding { JSONEncoding.default }

    var parameters: Parameters? {
        switch self {
        case let .logIn(username, password):
            [
                "username": username,
                "password": password
            ]
        case let .logOut(accessToken):
            [
                "access_token": accessToken
            ]
        case let .refreshToken(refreshToken):
            [
                "refresh_token": refreshToken
            ]
        }
    }

    var headers: HTTPHeaders? {
        switch self {
        case let .logOut(accessToken):
            [.authorization(bearerToken: accessToken)]
        default:
            nil
        }
    }
}

enum AppConfiguration {

    private static let apiBaseURLKey = "API_BASE_URL"
    private static let defaultAPIBaseURL = "http://127.0.0.1:8000/api"

    static func apiBaseURL(
        from infoDictionary: [String: Any]? = Bundle.main.infoDictionary
    ) -> String {
        let configuredValue = infoDictionary?[apiBaseURLKey] as? String
        let value = configuredValue?.trimmingCharacters(in: .whitespacesAndNewlines) ?? ""

        return value.isEmpty || value.hasPrefix("$(") ? defaultAPIBaseURL : value
    }
}
