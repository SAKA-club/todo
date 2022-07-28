// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'item.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Item _$ItemFromJson(Map<String, dynamic> json) => Item(
      title: json['title'] as String,
      body: json['body'] as String?,
      priority: json['priority'] as bool?,
      scheduleTime: json['schedule_time'] == null
          ? null
          : DateTime.parse(json['schedule_time'] as String),
      completeTime: json['complete_time'] == null
          ? null
          : DateTime.parse(json['complete_time'] as String),
    );

Map<String, dynamic> _$ItemToJson(Item instance) => <String, dynamic>{
      'title': instance.title,
      'body': instance.body,
      'priority': instance.priority,
      'schedule_time': instance.scheduleTime?.toIso8601String(),
      'complete_time': instance.completeTime?.toIso8601String(),
    };
