// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'classes.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

AuthBase _$AuthBaseFromJson(Map<String, dynamic> json) => AuthBase(
      accessToken: json['access_token'] as String,
    );

Map<String, dynamic> _$AuthBaseToJson(AuthBase instance) => <String, dynamic>{
      'access_token': instance.accessToken,
    };

GoogleDriveAPIBase _$GoogleDriveAPIBaseFromJson(Map<String, dynamic> json) =>
    GoogleDriveAPIBase(
      accessToken: json['access_token'] as String,
      googleDriveApiAccessToken:
          json['google_drive_api_access_token'] as String,
      googleDriveApiRefreshToken:
          json['google_drive_api_refresh_token'] as String,
    );

Map<String, dynamic> _$GoogleDriveAPIBaseToJson(GoogleDriveAPIBase instance) =>
    <String, dynamic>{
      'access_token': instance.accessToken,
      'google_drive_api_access_token': instance.googleDriveApiAccessToken,
      'google_drive_api_refresh_token': instance.googleDriveApiRefreshToken,
    };

YoutubeDataAPIBase _$YoutubeDataAPIBaseFromJson(Map<String, dynamic> json) =>
    YoutubeDataAPIBase(
      accessToken: json['access_token'] as String,
      youtubeDataApiAccessToken:
          json['youtube_data_api_access_token'] as String,
      youtubeDataApiRefreshToken:
          json['youtube_data_api_refresh_token'] as String,
    );

Map<String, dynamic> _$YoutubeDataAPIBaseToJson(YoutubeDataAPIBase instance) =>
    <String, dynamic>{
      'access_token': instance.accessToken,
      'youtube_data_api_access_token': instance.youtubeDataApiAccessToken,
      'youtube_data_api_refresh_token': instance.youtubeDataApiRefreshToken,
    };

GoogleDriveYoutubeDataAPIBase _$GoogleDriveYoutubeDataAPIBaseFromJson(
        Map<String, dynamic> json) =>
    GoogleDriveYoutubeDataAPIBase(
      accessToken: json['access_token'] as String,
      googleDriveApiAccessToken:
          json['google_drive_api_access_token'] as String,
      googleDriveApiRefreshToken:
          json['google_drive_api_refresh_token'] as String,
      youtubeDataApiAccessToken:
          json['youtube_data_api_access_token'] as String,
      youtubeDataApiRefreshToken:
          json['youtube_data_api_refresh_token'] as String,
    );

Map<String, dynamic> _$GoogleDriveYoutubeDataAPIBaseToJson(
        GoogleDriveYoutubeDataAPIBase instance) =>
    <String, dynamic>{
      'access_token': instance.accessToken,
      'google_drive_api_access_token': instance.googleDriveApiAccessToken,
      'google_drive_api_refresh_token': instance.googleDriveApiRefreshToken,
      'youtube_data_api_access_token': instance.youtubeDataApiAccessToken,
      'youtube_data_api_refresh_token': instance.youtubeDataApiRefreshToken,
    };
