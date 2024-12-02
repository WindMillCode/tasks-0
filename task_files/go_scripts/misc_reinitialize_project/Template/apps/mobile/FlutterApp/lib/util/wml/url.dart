class WMLEndpoint {
  final String Function() url;
  final bool automate;

  WMLEndpoint({required this.url, this.automate = false});
}

class WMLUri {
  Uri uri;

  WMLUri({
    required String scheme,
    required String host,
    int? port,
    String? path,
    String? query,
    String? fragment,
  }) : uri = Uri(
          scheme: scheme,
          host: host,
          port: port,
          path: path,
          query: query,
          fragment: fragment,
        );

  String get domain {
    return uri.port == 0 ? uri.host : '${uri.host}:${uri.port}';
  }

  String get fqdn {
    return '${uri.scheme}://${uri.host}${uri.port != 0 ? ':${uri.port}' : ''}';
  }

  @override
  String toString() {
    return uri.toString();
  }
}
