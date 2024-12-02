class HttpUtils {
  // Checks if the status code is a server error (5xx)
  static bool isServerError(int code) {
    return code >= 500 && code < 600;
  }

  // Checks if the status code is a client error (4xx)
  static bool isClientError(int code) {
    return code >= 400 && code < 500;
  }
}
