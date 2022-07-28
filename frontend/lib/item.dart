import 'package:json_annotation/json_annotation.dart';
part 'item.g.dart';

@JsonSerializable()
class Item {
  final String title;
  final String? body;
  final bool? priority;
  final DateTime? scheduleTime, completeTime;

  Item({required this.title, this.body, this.priority, this.scheduleTime, this.completeTime});

  factory Item.fromJson(Map<String, dynamic> json) => _$ItemFromJson(json);

  Map<String, dynamic> toJson() => _$ItemToJson(this);
}
