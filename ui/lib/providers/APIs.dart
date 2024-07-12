import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:shared_preferences/shared_preferences.dart';

class APIs {
  static final _baseUrl = baseUrl();
  static final searchUrl = "$_baseUrl/api/v1/tv/search";
  static final settingsUrl = "$_baseUrl/api/v1/setting/do";
  static final watchlistUrl = "$_baseUrl/api/v1/tv/watchlist";
  static final seriesDetailUrl = "$_baseUrl/api/v1/tv/series/";
  static final searchAndDownloadUrl = "$_baseUrl/api/v1/indexer/download";
  static final allIndexersUrl = "$_baseUrl/api/v1/indexer/";
  static final addIndexerUrl = "$_baseUrl/api/v1/indexer/add";
  static final delIndexerUrl = "$_baseUrl/api/v1/indexer/del/";
  static final allDownloadClientsUrl = "$_baseUrl/api/v1/downloader";
  static final addDownloadClientUrl = "$_baseUrl/api/v1/downloader/add";
  static final delDownloadClientUrl = "$_baseUrl/api/v1/downloader/del/";
  static final storageUrl = "$_baseUrl/api/v1/storage/";
  static final loginUrl = "$_baseUrl/api/login";
  static final loginSettingUrl = "$_baseUrl/api/v1/setting/auth";

  static const tmdbImgBaseUrl = "https://image.tmdb.org/t/p/w500/";

  static const tmdbApiKey = "tmdb_api_key";
  static const downloadDirKey = "download_dir";

  static final GlobalKey<NavigatorState> navigatorKey =
      GlobalKey<NavigatorState>();

  static String baseUrl() {
    if (kReleaseMode) {
      return "";
    }
    return "http://127.0.0.1:8080";
  }

  static Dio? dio1;

  static Future<Dio> getDio() async {
    if (dio1 != null) {
      return dio1!;
    }
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    var token = prefs.getString("token");
    var dio = Dio();
    dio.interceptors.add(InterceptorsWrapper(
      onRequest: (options, handler) {
        options.headers['Authorization'] = "Bearer $token";
        return handler.next(options);
      },
      onError: (error, handler) {
        if (error.response?.statusCode != null &&
            error.response?.statusCode! == 403) {
          final context = navigatorKey.currentContext;
          if (context != null) {
            context.go('/login');
          }
        }
        return handler.next(error);
      },
    ));
    dio1 = dio;
    return dio;
  }
}