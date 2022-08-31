import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'item.dart';

class ItemView extends StatelessWidget {
  const ItemView(this.item, {Key? key}) : super(key: key);
  final Item item;

  // Helper function in order to fill the itemView with null values
  Widget BodyText(String? description) {
    if (description == null || description.isEmpty) {
      return Container();
    }
    return Container(
      margin: const EdgeInsets.only(bottom: 10),
      child: Text(
        description,
        overflow: TextOverflow.ellipsis,
        textWidthBasis: TextWidthBasis.longestLine,
        maxLines: 2,
        style: TextStyle(
          fontSize: 14,
          color: Color(0xFF545454),
          fontFamily: 'SourceSansPro',
        ),
      ),
    );
  }

  //  final bool? priority;
  //   final DateTime? scheduleTime, completeTime;

  Widget DateField(DateTime? description) {
    if (description == null) {
      return Container();
    }
    return Container(
      child: Text(
        item.completeTime.toString(),
        style: TextStyle(
          fontSize: 14,
          fontFamily: 'SourceSansPro',
          color: Color(0xFF545454),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Container(
          margin: const EdgeInsets.only(top: 10, bottom: 5),
          child: Text(
            item.title,
            overflow: TextOverflow.ellipsis,
            style: TextStyle(
              fontSize: 20,
              color: Colors.black,
              fontFamily: 'Nunito',
              fontWeight: FontWeight.w200,
            ),
          ),
        ),
        BodyText(item.body),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Row(
              children: [
                Container(
                  margin: EdgeInsets.only(right: 5),
                  padding: EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.all(Radius.circular(20.0)),
                    color: Color(0xFFD9D9D9),
                  ),
                  child: Text(
                    'Work',
                    style: TextStyle(
                      color: Colors.black,
                      fontSize: 12,
                      fontFamily: 'Nunito',
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ),
                Container(
                  margin: EdgeInsets.only(right: 5),
                  padding: EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.all(Radius.circular(20.0)),
                    color: Color(0xFFD9D9D9),
                  ),
                  child: Text(
                    'Work',
                    style: TextStyle(
                      color: Colors.black,
                      fontSize: 12,
                      fontFamily: 'Nunito',
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ),
              ],
            ),
            DateField(item.completeTime),
          ],
        ),
        Container(
          margin: EdgeInsets.only(top: 10),
          color: Color(0xFFEAEAEA),
          height: 1,
        )
      ],
    );
  }
}
